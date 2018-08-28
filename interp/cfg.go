package interp

import (
	"log"
	"path"
	"reflect"
	"unicode"
)

type SymKind uint

const (
	Const SymKind = iota
	Typ
	Var
	Func
	Bin
	Bltn
)

// A Symbol represents an interpreter object such as type, constant, var, func, builtin or binary object
type Symbol struct {
	kind    SymKind
	typ     *Type       // Type of value
	node    *Node       // Node value if index is negative
	index   int         // index of value in frame or -1
	val     interface{} // default value (used for constants)
	path    string      // package path if typ.cat is SrcPkgT or BinPkgT
	builtin Builtin     // Builtin function or nil
	global  bool        // true if symbol is defined in global space
}

// A SymMap stores symbols indexed by name
type SymMap map[string]*Symbol

// Scope type stores the list of visible symbols at current scope level
type Scope struct {
	anc    *Scope // Ancestor upper scope
	level  int    // Frame level: number of frame indirections to access var during execution
	size   int    // Frame size: number of entries to allocate during execution (package scope only)
	sym    SymMap // Map of symbols defined in this current scope
	global bool   // true if scope refers to global space (single frame for universe and package level scopes)
}

// Create a new scope and chain it to the current one
func (s *Scope) push(indirect int) *Scope {
	var size int
	if indirect == 0 {
		size = s.size // propagate size as scopes at same level share the same frame
	}
	global := s.global && indirect == 0
	return &Scope{anc: s, global: global, level: s.level + indirect, size: size, sym: map[string]*Symbol{}}
}

func (s *Scope) pop() *Scope {
	if s.level == s.anc.level {
		s.anc.size = s.size // propagate size as scopes at same level share the same frame
	}
	return s.anc
}

// Lookup for a symbol in the current scope, and upper ones if not found
func (s *Scope) lookup(ident string) (*Symbol, int, bool) {
	level := s.level
	for s != nil {
		if sym, ok := s.sym[ident]; ok {
			return sym, level - s.level, true
		}
		s = s.anc
	}
	return nil, 0, false
}

func (s *Scope) getType(ident string) *Type {
	var t *Type
	if sym, _, found := s.lookup(ident); found {
		if sym.kind == Typ {
			t = sym.typ
		}
	}
	return t
}

// Cfg generates a control flow graph (CFG) from AST (wiring successors in AST)
// and pre-compute frame sizes and indexes for all un-named (temporary) and named
// variables. A list of nodes of init functions is returned.
// Following this pass, the CFG is ready to run
func (interp *Interpreter) Cfg(root *Node) []*Node {
	scope := interp.universe
	var loop, loopRestart *Node
	var funcDef bool // True if a function is defined in the current frame context
	var initNodes []*Node
	var exports *BinMap
	var expval *ValueMap
	var iotaValue int
	var pkgName string

	root.Walk(func(n *Node) bool {
		// Pre-order processing
		switch n.kind {
		case Define, AssignStmt:
			if l := len(n.child); n.anc.kind == ConstDecl && l == 1 {
				// Implicit iota assignment. TODO: replicate previous explicit assignment
				n.child = append(n.child, &Node{anc: n, interp: interp, kind: Ident, ident: "iota"})
			} else if l%2 == 1 {
				// Odd number of children: remove the type node, useless for assign
				i := l / 2
				n.child = append(n.child[:i], n.child[i+1:]...)
			}

		case BlockStmt:
			scope = scope.push(0)

		case File:
			pkgName = n.child[0].ident
			if _, ok := interp.scope[pkgName]; !ok {
				interp.scope[pkgName] = scope.push(0)
			}
			scope = interp.scope[pkgName]
			scope.size = interp.fsize
			if pkg, ok := interp.Exports[pkgName]; ok {
				exports = pkg
				expval = interp.Expval[pkgName]
			} else {
				x := make(BinMap)
				exports = &x
				interp.Exports[pkgName] = exports
				y := make(ValueMap)
				expval = &y
				interp.Expval[pkgName] = expval
			}

		case For0, ForRangeStmt:
			loop, loopRestart = n, n.child[0]
			scope = scope.push(0)

		case For1, For2, For3, For3a, For4:
			loop, loopRestart = n, n.child[len(n.child)-1]
			scope = scope.push(0)

		case FuncDecl, FuncLit:
			// Add a frame indirection level as we enter in a func
			scope = scope.push(1)
			if n.child[1].ident == "init" {
				initNodes = append(initNodes, n)
			}
			if len(n.child[0].child) > 0 {
				// function is a method, add it to the related type
				var t *Type
				var tname string
				n.ident = n.child[1].ident
				recv := n.child[0].child[0]
				if len(recv.child) < 2 {
					// Receiver var name is skipped in method declaration (fix that in AST ?)
					tname = recv.child[0].ident
				} else {
					tname = recv.child[1].ident
				}
				if tname == "" {
					tname = recv.child[1].child[0].ident
					elemtype := scope.getType(tname)
					t = &Type{cat: PtrT, val: elemtype}
					elemtype.method = append(elemtype.method, n)
				} else {
					t = scope.getType(tname)
				}
				t.method = append(t.method, n)
			}
			if len(n.child[2].child) == 2 {
				// allocate entries for return values at start of frame
				scope.size += len(n.child[2].child[1].child)
			}
			funcDef = false

		case If0, If1, If2, If3:
			scope = scope.push(0)

		case Switch0:
			// Make sure default clause is in last position
			c := n.child[1].child
			if i, l := getDefault(n), len(c)-1; i >= 0 && i != l {
				c[i], c[l] = c[l], c[i]
			}
			scope = scope.push(0)

		case TypeSpec:
			// Type analysis is performed recursively and no post-order processing
			// needs to be done for types, so do not dive in subtree
			if n.child[1].kind == Ident {
				// Create a type alias of an existing one
				n.typ = &Type{cat: AliasT, val: nodeType(interp, scope, n.child[1])}
			} else {
				// Define a new type
				n.typ = nodeType(interp, scope, n.child[1])
			}
			scope.sym[n.child[0].ident] = &Symbol{kind: Typ, typ: n.typ}
			// TODO: export type for use by runtime
			return false

		case ArrayType, BasicLit, ChanType, MapType, StructType:
			n.typ = nodeType(interp, scope, n)
			return false
		}
		return true
	}, func(n *Node) {
		// Post-order processing
		switch n.kind {
		case Address:
			wireChild(n)
			n.typ = &Type{cat: PtrT, val: n.child[0].typ}

		case ArrayType:
			// TODO: move to pre-processing ? See when handling complex array type def
			n.typ = &Type{cat: ArrayT, val: nodeType(interp, scope, n.child[1])}

		case Define, AssignStmt:
			wireChild(n)
			if n.kind == Define || n.anc.kind == VarDecl || n.anc.kind == ConstDecl {
				// Force definition of assigned ident in current scope
				name := n.child[0].ident
				if scope.global {
					if sym, _, ok := scope.lookup(name); ok {
						n.child[0].findex = sym.index
					} else {
						interp.fsize++
						scope.size = interp.fsize
						scope.sym[name] = &Symbol{index: scope.size, global: true}
						n.child[0].findex = scope.size
					}
				} else {
					scope.size++
					scope.sym[name] = &Symbol{index: scope.size}
					n.child[0].findex = scope.size
				}
				n.child[0].typ = n.child[1].typ
				if n.child[1].action == GetFunc {
					scope.sym[name].index = -1
					scope.sym[name].node = n.child[1]
				}
			}
			n.findex = n.child[0].findex
			n.val = n.child[0].val
			// Propagate type
			// TODO: Check that existing destination type matches source type
			n.typ = n.child[0].typ
			if sym, level, ok := scope.lookup(n.child[0].ident); ok {
				sym.typ = n.typ
				n.level = level
			}
			// If LHS is an indirection, get reference instead of value, to allow setting
			if n.child[0].action == GetIndex {
				if n.child[0].child[0].typ.cat == MapT {
					n.child[0].run = getMap
					n.run = assignMap
				} else if n.child[0].child[0].typ.cat == PtrT {
					// Handle the case where the receiver is a pointer to an object
					n.child[0].run = getPtrIndexAddr
					n.run = assignPtrField
				} else {
					n.child[0].run = getIndexAddr
					n.run = assignField
				}
			} else if n.child[0].action == Star {
				n.findex = n.child[0].child[0].findex
				n.run = indirectAssign
			}
			if n.anc.kind == ConstDecl {
				iotaValue++
			}

		case IncDecStmt:
			wireChild(n)
			n.findex = n.child[0].findex
			n.child[0].typ = scope.getType("int")
			n.typ = n.child[0].typ
			if sym, level, ok := scope.lookup(n.child[0].ident); ok {
				sym.typ = n.typ
				n.level = level
			}
			if n.child[0].action == Star {
				n.findex = n.child[0].child[0].findex
				n.run = indirectInc
			}

		case DefineX, AssignXStmt:
			wireChild(n)
			l := len(n.child) - 1
			if n.kind == DefineX {
				// retrieve assigned value types from call signature
				var types []*Type
				switch n.child[l].kind {
				case CallExpr:
					if funtype := n.child[l].child[0].typ; funtype.cat == ValueT {
						// Handle functions imported from runtime
						for i := 0; i < funtype.rtype.NumOut(); i++ {
							types = append(types, &Type{cat: ValueT, rtype: funtype.rtype.Out(i)})
						}
					} else {
						types = funtype.ret
					}
				case TypeAssertExpr, IndexExpr:
					types = append(types, n.child[l].child[1].typ, scope.getType("error"))
				default:
					log.Fatalln(n.index, "Assign expression unsupported:", n.child[l].kind)
				}
				// Force definition of assigned idents in current scope
				for i, c := range n.child[:l] {
					if scope.global {
						interp.fsize++
						scope.size = interp.fsize
					} else {
						scope.size++
					}
					scope.sym[c.ident] = &Symbol{index: scope.size, typ: types[i], global: scope.global}
					c.findex = scope.size
				}
			}

		case BinaryExpr:
			wireChild(n)
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			n.typ = n.child[0].typ

		case IndexExpr:
			wireChild(n)
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			n.typ = n.child[0].typ.val
			n.recv = n
			if n.child[0].typ.cat == MapT {
				scope.size++ // Reserve an entry for getIndexMap 2nd return value
				n.run = getIndexMap
			}

		case BlockStmt:
			wireChild(n)
			n.findex = n.child[len(n.child)-1].findex
			scope = scope.pop()

		case ConstDecl:
			wireChild(n)
			iotaValue = 0

		case DeclStmt, ExprStmt, VarDecl, ParenExpr, SendStmt:
			wireChild(n)
			n.findex = n.child[len(n.child)-1].findex

		case Break:
			n.tnext = loop

		case CallExpr:
			wireChild(n)
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			if n.child[0].sym != nil && n.child[0].sym.kind == Bltn {
				n.run = n.child[0].sym.builtin
				n.child[0].typ = &Type{cat: BuiltinT}
				if n.child[0].ident == "make" {
					if n.typ = scope.getType(n.child[1].ident); n.typ == nil {
						n.typ = nodeType(interp, scope, n.child[1])
					}
					n.child[1].val = n.typ
					n.child[1].kind = BasicLit
				}
			} else if n.child[0].isType(scope) {
				n.typ = n.child[0].typ
				if n.typ.cat == AliasT {
					n.run = convert
				} else {
					if n.child[1].typ.cat == FuncT {
						// Convert type of an interpreter function to another binary type:
						// generate a different callback wrapper
						n.run = convertFuncBin
					} else {
						n.run = convertBin
					}
				}
			} else if n.child[0].kind == SelectorImport {
				// TODO: Should process according to child type, not kind.
				n.fsize = n.child[0].fsize
				rtype := n.child[0].val.(reflect.Value).Type()
				if rtype.NumOut() > 0 {
					n.typ = &Type{cat: ValueT, rtype: rtype.Out(0)}
				}
				n.child[0].kind = BasicLit
				for i, c := range n.child[1:] {
					// Wrap function defintion so it can be called from runtime
					if c.kind == FuncLit {
						n.child[1+i].rval = reflect.MakeFunc(rtype.In(i), c.wrapNode)
						n.child[1+i].kind = Rvalue
					} else if c.ident == "nil" {
						n.child[1+i].rval = reflect.New(rtype.In(i)).Elem()
						n.child[1+i].kind = Rvalue
					} else if c.typ != nil && c.typ.cat == FuncT {
						n.child[1+i].rval = reflect.MakeFunc(rtype.In(i), c.wrapNode)
						n.child[1+i].kind = Rvalue
					}
				}
				// TODO: handle multiple return value
				if len(n.child) == 2 && n.child[1].fsize > 1 {
					n.run = callBinX
				} else {
					n.run = callBin
				}
			} else if n.child[0].kind == SelectorExpr {
				if n.child[0].typ.cat == ValueT {
					//log.Println(n.index, "selector callBinMethod", n.child[0].typ.rtype)
					n.run = callBinMethod
					// TODO: handle multiple return value (_test/time2.go)
					n.child[0].kind = BasicLit // Temporary hack for force value() to return n.val at run
					n.typ = &Type{cat: ValueT, rtype: n.child[0].typ.rtype}
					n.fsize = n.child[0].fsize
				} else if n.child[0].typ.cat == PtrT && n.child[0].typ.val.cat == ValueT {
					n.run = callBinMethod
					n.child[0].kind = BasicLit // Temporary hack for force value() to return n.val at run
					n.fsize = n.child[0].fsize
					// TODO: handle type ?
				} else if n.child[0].typ.cat == SrcPkgT {
					n.val = n.child[0].val
					if def := n.val.(*Node); def != nil {
						// Reserve as many frame entries as nb of ret values for called function
						// node frame index should point to the first entry
						j := len(def.child[2].child) - 1
						l := len(def.child[2].child[j].child) // Number of return values for def
						if l == 1 {
							// If def returns exactly one value, propagate its type in call node.
							// Multiple return values will be handled differently through AssignX.
							n.typ = scope.getType(def.child[2].child[j].child[0].child[0].ident)
						}
						n.fsize = l
					}
				} else {
					// Resolve method and receiver path, store them in node static value for run
					if methodDecl := n.child[0].val.(*Node); len(methodDecl.child[2].child) > 1 {
						// Allocate frame for method return values (if any)
						n.fsize = len(methodDecl.child[2].child[1].child)
					} else {
						n.fsize = 0
					}
					n.child[0].findex = -1 // To force reading value from node instead of frame (methods)
				}
			} else if sym, _, _ := scope.lookup(n.child[0].ident); sym != nil {
				if sym.typ != nil && sym.typ.cat == BinT {
					//log.Println(n.index, "BinT", n.child[0].ident)
					n.run = callBin
					n.typ = &Type{cat: ValueT}
					r := sym.val.(reflect.Value)
					n.child[0].fsize = r.Type().NumOut()
					n.child[0].val = r
					n.child[0].kind = BasicLit
				} else if sym.typ != nil && sym.typ.cat == ValueT {
					n.run = callDirectBin
					n.typ = &Type{cat: ValueT}
				} else {
					n.val = sym.node
					if def := n.val.(*Node); def != nil {
						// Reserve as many frame entries as nb of ret values for called function
						// node frame index should point to the first entry
						j := len(def.child[2].child) - 1
						l := len(def.child[2].child[j].child) // Number of return values for def
						if l == 1 {
							// If def returns exactly one value, propagate its type in call node.
							// Multiple return values will be handled differently through AssignX.
							n.typ = scope.getType(def.child[2].child[j].child[0].child[0].ident)
						}
						n.fsize = l
					} else {
						log.Println(n.index, "call to unknown def", n.child[0].ident)
					}
				}
			} else if n.child[0].kind == SelectorSrc {
				// Forward type of first returned value
				// Multiple return values will be handled differently through AssignX.
				if len(n.child[0].typ.ret) > 0 {
					n.typ = n.child[0].typ.ret[0]
					n.fsize = len(n.child[0].typ.ret)
				}
			}
			// Reserve entries in frame to store results of call
			scope.size += n.fsize
			if scope.global {
				interp.fsize += n.fsize
				scope.size = interp.fsize
			} else {
				scope.size += n.fsize
			}
			if funcDef {
				// Trigger frame indirection to handle nested functions
				n.action = CallF
			}

		case CaseClause:
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			n.tnext = n.child[len(n.child)-1].start

		case CompositeLitExpr:
			wireChild(n)
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			if n.child[0].typ == nil {
				n.child[0].typ = nodeType(interp, scope, n.child[0])
			}
			// TODO: Check that composite litteral expr matches corresponding type
			n.typ = n.child[0].typ
			switch n.typ.cat {
			case ArrayT:
				n.run = arrayLit
			case MapT:
				n.run = mapLit
			case StructT:
				n.action, n.run = CompositeLit, compositeLit
				// Handle object assign from sparse key / values
				if len(n.child) > 1 && n.child[1].kind == KeyValueExpr {
					n.run = compositeSparse
					n.typ = nodeType(interp, scope, n.child[0])
					for _, c := range n.child[1:] {
						c.findex = n.typ.fieldIndex(c.child[0].ident)
					}
				}
			}

		case Continue:
			n.tnext = loopRestart

		case Field:
			// A single child node (no ident, just type) means that the field refers
			// to a return value, and space on frame should be accordingly allocated.
			// Otherwise, just point to corresponding location in frame, resolved in
			// ident child.
			l := len(n.child) - 1
			n.typ = nodeType(interp, scope, n.child[l])
			if l == 0 {
				if scope.global {
					interp.fsize++
					scope.size = interp.fsize
				} else {
					scope.size++
				}
				n.findex = scope.size
			} else {
				for _, f := range n.child[:l] {
					scope.sym[f.ident].typ = n.typ
				}
			}

		case File:
			wireChild(n)
			scope = scope.pop()
			n.fsize = scope.size + 1

		case For0: // for {}
			body := n.child[0]
			n.start = body.start
			body.tnext = n.start
			loop, loopRestart = nil, nil
			scope = scope.pop()

		case For1: // for cond {}
			cond, body := n.child[0], n.child[1]
			n.start = cond.start
			cond.tnext = body.start
			cond.fnext = n
			body.tnext = cond.start
			loop, loopRestart = nil, nil
			scope = scope.pop()

		case For2: // for init; cond; {}
			init, cond, body := n.child[0], n.child[1], n.child[2]
			n.start = init.start
			init.tnext = cond.start
			cond.tnext = body.start
			cond.fnext = n
			body.tnext = cond.start
			loop, loopRestart = nil, nil
			scope = scope.pop()

		case For3: // for ; cond; post {}
			cond, post, body := n.child[0], n.child[1], n.child[2]
			n.start = cond.start
			cond.tnext = body.start
			cond.fnext = n
			body.tnext = post.start
			post.tnext = cond.start
			loop, loopRestart = nil, nil
			scope = scope.pop()

		case For3a: // for int; ; post {}
			init, post, body := n.child[0], n.child[1], n.child[2]
			n.start = init.start
			init.tnext = body.start
			body.tnext = post.start
			post.tnext = body.start
			loop, loopRestart = nil, nil
			scope = scope.pop()

		case For4: // for init; cond; post {}
			init, cond, post, body := n.child[0], n.child[1], n.child[2], n.child[3]
			n.start = init.start
			init.tnext = cond.start
			cond.tnext = body.start
			cond.fnext = n
			body.tnext = post.start
			post.tnext = cond.start
			loop, loopRestart = nil, nil
			scope = scope.pop()

		case ForRangeStmt:
			loop, loopRestart = nil, nil
			n.start = n.child[0].start
			n.findex = n.child[0].findex
			n.child[0].fnext = n
			scope = scope.pop()

		case FuncDecl:
			n.flen = scope.size + 1
			if len(n.child[0].child) > 0 {
				// Store receiver frame location (used at run)
				n.child[0].findex = n.child[0].child[0].child[0].findex
			}
			scope = scope.pop()
			funcName := n.child[1].ident
			if canExport(funcName) {
				(*exports)[funcName] = reflect.MakeFunc(n.child[2].typ.TypeOf(), n.wrapNode).Interface()
				(*expval)[funcName] = reflect.MakeFunc(n.child[2].typ.TypeOf(), n.wrapNode)
			}
			n.typ = n.child[2].typ
			n.val = n
			interp.scope[pkgName].sym[funcName].index = -1
			interp.scope[pkgName].sym[funcName].typ = n.typ
			interp.scope[pkgName].sym[funcName].kind = Func
			interp.scope[pkgName].sym[funcName].node = n

		case FuncLit:
			n.typ = n.child[2].typ
			n.val = n
			n.flen = scope.size + 1
			scope = scope.pop()
			funcDef = true

		case FuncType:
			n.typ = nodeType(interp, scope, n)
			// Store list of parameter frame indices in params val
			var list []int
			for _, c := range n.child[0].child {
				for _, f := range c.child[:len(c.child)-1] {
					list = append(list, f.findex)
				}
			}
			n.child[0].val = list
			// TODO: do the same for return values

		case GoStmt:
			wireChild(n)
			// TODO: should error if call expression refers to a builtin
			n.child[0].run = callGoRoutine

		case Ident:
			if n.anc.kind == File || (n.anc.kind == SelectorExpr && n.anc.child[0] != n) {
				// skip symbol creation/lookup for idents used as key
			} else if sym, level, ok := scope.lookup(n.ident); ok {
				//log.Println(n.index, "found", n.ident)
				n.typ, n.findex, n.level = sym.typ, sym.index, level
				if n.findex < 0 {
					n.val = sym.node
					n.kind = sym.node.kind
				} else {
					n.sym = sym
					if sym.val != nil {
						n.val = sym.val
						n.kind = BasicLit
					} else if n.ident == "iota" {
						n.val = iotaValue
						n.kind = BasicLit
					}
				}
				n.recv = n
			} else {
				//log.Println(n.index, "not found", n.ident)
				if n.ident == "_" || n.anc.kind == Define || n.anc.kind == DefineX || n.anc.kind == Field || n.anc.kind == RangeStmt || n.anc.kind == ValueSpec {
					// Create a new local symbol for func argument or local var definition
					if scope.global {
						interp.fsize++
						scope.size = interp.fsize
					} else {
						scope.size++
					}
					scope.sym[n.ident] = &Symbol{index: scope.size, global: scope.global}
					n.findex = scope.size
				} else {
					// symbol may be defined globally elsewhere later, add an entry at pkg level
					interp.fsize++
					interp.scope[pkgName].size = interp.fsize
					interp.scope[pkgName].sym[n.ident] = &Symbol{index: interp.fsize, global: true}
					n.sym = interp.scope[pkgName].sym[n.ident]
					n.level = scope.level
					n.findex = interp.fsize
					//log.Println(n.index, "define pkg sym", pkgName, n.ident, n.level, n.findex)
				}
			}

		case If0: // if cond {}
			cond, tbody := n.child[0], n.child[1]
			n.start = cond.start
			cond.tnext = tbody.start
			cond.fnext = n
			tbody.tnext = n
			scope = scope.pop()

		case If1: // if cond {} else {}
			cond, tbody, fbody := n.child[0], n.child[1], n.child[2]
			n.start = cond.start
			cond.tnext = tbody.start
			cond.fnext = fbody.start
			tbody.tnext = n
			fbody.tnext = n
			scope = scope.pop()

		case If2: // if init; cond {}
			init, cond, tbody := n.child[0], n.child[1], n.child[2]
			n.start = init.start
			tbody.tnext = n
			init.tnext = cond.start
			cond.tnext = tbody.start
			cond.fnext = n
			scope = scope.pop()

		case If3: // if init; cond {} else {}
			init, cond, tbody, fbody := n.child[0], n.child[1], n.child[2], n.child[3]
			n.start = init.start
			init.tnext = cond.start
			cond.tnext = tbody.start
			cond.fnext = fbody.start
			tbody.tnext = n
			fbody.tnext = n
			scope = scope.pop()

		case ImportSpec:
			var name, ipath string
			if len(n.child) == 2 {
				ipath = n.child[1].val.(string)
				name = n.child[0].ident
			} else {
				ipath = n.child[0].val.(string)
				name = path.Base(ipath)
			}
			if pkg, ok := interp.binValue[ipath]; ok {
				if name == "." {
					for n, s := range pkg {
						scope.sym[n] = &Symbol{typ: &Type{cat: BinT}, val: s}
					}
				} else {
					scope.sym[name] = &Symbol{typ: &Type{cat: BinPkgT}, path: ipath}
				}
			} else {
				// TODO: make sure we do not import a src package more than once
				interp.importSrcFile(ipath)
				scope.sym[name] = &Symbol{typ: &Type{cat: SrcPkgT}, path: ipath}
			}

		case KeyValueExpr:
			wireChild(n)

		case LandExpr:
			n.start = n.child[0].start
			n.child[0].tnext = n.child[1].start
			n.child[0].fnext = n
			n.child[1].tnext = n
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			n.typ = n.child[0].typ

		case LorExpr:
			n.start = n.child[0].start
			n.child[0].tnext = n
			n.child[0].fnext = n.child[1].start
			n.child[1].tnext = n
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			n.typ = n.child[0].typ

		case RangeStmt:
			n.start = n
			n.child[3].tnext = n
			n.tnext = n.child[3].start
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size

		case ReturnStmt:
			wireChild(n)
			n.tnext = nil

		case SelectorExpr:
			wireChild(n)
			if scope.global {
				interp.fsize++
				scope.size = interp.fsize
			} else {
				scope.size++
			}
			n.findex = scope.size
			n.typ = n.child[0].typ
			n.recv = n.child[0].recv
			if n.typ == nil {
				log.Fatal("typ should not be nil:", n.index, n.child[0])
			}
			//log.Println(n.index, n.child[0].ident+"."+n.child[1].ident, n.typ.cat)
			if n.typ.cat == ValueT {
				// Handle object defined in runtime
				if method, ok := n.typ.rtype.MethodByName(n.child[1].ident); ok {
					if method.Func.IsValid() {
						n.rval = method.Func
						n.run = nop
						//log.Println(n.index, "select method", n.rval, method.Index)
					} else {
						n.val = method.Index
						//n.run = getIndexBinMethod
						n.run = nop
					}
					n.fsize = method.Type.NumOut()
				} else {
					// Method can be only resolved from value at execution
					n.run = getIndexBinMethod
				}
			} else if n.typ.cat == PtrT && n.typ.val.cat == ValueT {
				// Handle pointer on object defined in runtime
				if field, ok := n.typ.val.rtype.Elem().FieldByName(n.child[1].ident); ok {
					n.typ = &Type{cat: ValueT, rtype: field.Type}
					n.val = field.Index
					n.run = getPtrIndexBin
				} else if method, ok := n.typ.val.rtype.MethodByName(n.child[1].ident); ok {
					n.val = method.Func
					n.fsize = method.Type.NumOut()
					n.run = nop
				} else {
					log.Println(n.index, "selector unresolved")
				}
			} else if n.typ.cat == BinPkgT {
				// Resolve binary package symbol
				name := n.child[1].ident
				pkgSym := interp.binValue[n.child[0].sym.path]
				if s, ok := pkgSym[name]; ok {
					n.kind = SelectorImport
					n.val = s
					if typ := s.Type(); typ.Kind() == reflect.Func {
						n.typ = &Type{cat: ValueT, rtype: typ}
						n.fsize = typ.NumOut()
						//} else if typ.Kind() == reflect.Ptr {
						// a symbol of kind pointer must be dereferenced to access type
						//	n.typ = &Type{cat: ValueT, rtype: typ.Elem(), rzero: n.val.(reflect.Value).Elem()}
					} else {
						n.typ = &Type{cat: ValueT, rtype: typ}
						n.rval = n.val.(reflect.Value)
						n.kind = Rvalue
					}
					n.run = nop
				}
			} else if n.typ.cat == ArrayT {
				n.typ = n.typ.val
				n.run = nop
			} else if n.typ.cat == SrcPkgT {
				// Resolve source package symbol
				if sym, ok := interp.scope[n.child[0].ident].sym[n.child[1].ident]; ok {
					n.val = sym.node
					n.run = nop
					n.kind = SelectorSrc
					n.typ = sym.typ
				} else {
					log.Println(n.index, "selector unresolved:", n.child[0].ident+"."+n.child[1].ident)
				}
			} else if fi := n.typ.fieldIndex(n.child[1].ident); fi >= 0 {
				// Resolve struct field index
				if n.typ.cat == PtrT {
					n.run = getPtrIndex
				}
				n.typ = n.typ.fieldType(fi)
				n.child[1].kind = BasicLit
				n.child[1].val = fi
			} else if m, lind := n.typ.lookupMethod(n.child[1].ident); m != nil {
				// Handle method
				n.run = nop
				n.val = m
				n.child[1].val = lind
				n.typ = m.typ
			} else {
				// Handle promoted field in embedded struct
				if ti := n.typ.lookupField(n.child[1].ident); len(ti) > 0 {
					n.child[1].kind = BasicLit
					n.child[1].val = ti
					n.run = getIndexSeq
				} else {
					log.Println(n.index, "Selector not found:", n.child[1].ident)
					n.run = nop
					//panic("Field not found in selector")
				}
			}

		case Switch0:
			n.start = n.child[1].start
			// Chain case clauses
			clauses := n.child[1].child
			l := len(clauses)
			for i, c := range clauses[:l-1] {
				// chain to next clause
				c.tnext = c.child[1].start
				c.child[1].tnext = n
				c.fnext = clauses[i+1]
			}
			// Handle last clause
			if c := clauses[l-1]; len(c.child) > 1 {
				// No default clause
				c.tnext = c.child[1].start
				c.fnext = n
				c.child[1].tnext = n
			} else {
				// Default
				c.tnext = c.child[0].start
				c.child[0].tnext = n
			}
			scope = scope.pop()

		case TypeAssertExpr:
			if n.child[1].typ == nil {
				n.child[1].typ = scope.getType(n.child[1].ident)
			}
			n.typ = n.child[1].typ

		case ValueSpec:
			l := len(n.child) - 1
			if n.typ = n.child[l].typ; n.typ == nil {
				n.typ = scope.getType(n.child[l].ident)
			}
			for _, c := range n.child[:l] {
				c.typ = n.typ
				scope.sym[c.ident].typ = n.typ
			}
		}
	})
	return initNodes
}

// Find default case clause index of a switch statement, if any
func getDefault(n *Node) int {
	for i, c := range n.child[1].child {
		if len(c.child) == 1 {
			return i
		}
	}
	return -1
}

// isType returns true if node refers to a type definition, false otherwise
func (n *Node) isType(scope *Scope) bool {
	switch n.kind {
	case ArrayType, ChanType, FuncType, MapType, StructType:
		return true
	case Ident:
		return scope.getType(n.ident) != nil
	case Rvalue:
		return n.typ.rtype.Kind() != reflect.Func
	}
	return false
}

// Wire AST nodes for CFG in subtree
func wireChild(n *Node) {
	// Set start node, in subtree (propagated to ancestors by post-order processing)
	for _, child := range n.child {
		switch child.kind {
		case ArrayType, ChanType, FuncDecl, MapType, BasicLit, Ident:
			continue
		default:
			n.start = child.start
		}
		break
	}

	// Chain sequential operations inside a block (next is right sibling)
	for i := 1; i < len(n.child); i++ {
		n.child[i-1].tnext = n.child[i].start
	}

	// Chain subtree next to self
	for i := len(n.child) - 1; i >= 0; i-- {
		switch n.child[i].kind {
		case ArrayType, ChanType, MapType, FuncDecl, BasicLit, Ident:
			continue
		case Break, Continue, ReturnStmt:
			// tnext is already computed, no change
		default:
			n.child[i].tnext = n
		}
		break
	}
}

func canExport(name string) bool {
	if r := []rune(name); len(r) > 0 && unicode.IsUpper(r[0]) {
		return true
	}
	return false
}
