package stdlib

// Generated by 'goexports text/template/parse'. Do not edit!

import (
	"reflect"
	"text/template/parse"
)

func init() {
	Value["text/template/parse"] = map[string]reflect.Value{
		"IsEmptyTree":    reflect.ValueOf(parse.IsEmptyTree),
		"New":            reflect.ValueOf(parse.New),
		"NewIdentifier":  reflect.ValueOf(parse.NewIdentifier),
		"NodeAction":     reflect.ValueOf(parse.NodeAction),
		"NodeBool":       reflect.ValueOf(parse.NodeBool),
		"NodeChain":      reflect.ValueOf(parse.NodeChain),
		"NodeCommand":    reflect.ValueOf(parse.NodeCommand),
		"NodeDot":        reflect.ValueOf(parse.NodeDot),
		"NodeField":      reflect.ValueOf(parse.NodeField),
		"NodeIdentifier": reflect.ValueOf(parse.NodeIdentifier),
		"NodeIf":         reflect.ValueOf(parse.NodeIf),
		"NodeList":       reflect.ValueOf(parse.NodeList),
		"NodeNil":        reflect.ValueOf(parse.NodeNil),
		"NodeNumber":     reflect.ValueOf(parse.NodeNumber),
		"NodePipe":       reflect.ValueOf(parse.NodePipe),
		"NodeRange":      reflect.ValueOf(parse.NodeRange),
		"NodeString":     reflect.ValueOf(parse.NodeString),
		"NodeTemplate":   reflect.ValueOf(parse.NodeTemplate),
		"NodeText":       reflect.ValueOf(parse.NodeText),
		"NodeVariable":   reflect.ValueOf(parse.NodeVariable),
		"NodeWith":       reflect.ValueOf(parse.NodeWith),
		"Parse":          reflect.ValueOf(parse.Parse),
	}
	Type["text/template/parse"] = map[string]reflect.Type{
		"ActionNode":     reflect.TypeOf((*parse.ActionNode)(nil)).Elem(),
		"BoolNode":       reflect.TypeOf((*parse.BoolNode)(nil)).Elem(),
		"BranchNode":     reflect.TypeOf((*parse.BranchNode)(nil)).Elem(),
		"ChainNode":      reflect.TypeOf((*parse.ChainNode)(nil)).Elem(),
		"CommandNode":    reflect.TypeOf((*parse.CommandNode)(nil)).Elem(),
		"DotNode":        reflect.TypeOf((*parse.DotNode)(nil)).Elem(),
		"FieldNode":      reflect.TypeOf((*parse.FieldNode)(nil)).Elem(),
		"IdentifierNode": reflect.TypeOf((*parse.IdentifierNode)(nil)).Elem(),
		"IfNode":         reflect.TypeOf((*parse.IfNode)(nil)).Elem(),
		"ListNode":       reflect.TypeOf((*parse.ListNode)(nil)).Elem(),
		"NilNode":        reflect.TypeOf((*parse.NilNode)(nil)).Elem(),
		"Node":           reflect.TypeOf((*parse.Node)(nil)).Elem(),
		"NodeType":       reflect.TypeOf((*parse.NodeType)(nil)).Elem(),
		"NumberNode":     reflect.TypeOf((*parse.NumberNode)(nil)).Elem(),
		"PipeNode":       reflect.TypeOf((*parse.PipeNode)(nil)).Elem(),
		"Pos":            reflect.TypeOf((*parse.Pos)(nil)).Elem(),
		"RangeNode":      reflect.TypeOf((*parse.RangeNode)(nil)).Elem(),
		"StringNode":     reflect.TypeOf((*parse.StringNode)(nil)).Elem(),
		"TemplateNode":   reflect.TypeOf((*parse.TemplateNode)(nil)).Elem(),
		"TextNode":       reflect.TypeOf((*parse.TextNode)(nil)).Elem(),
		"Tree":           reflect.TypeOf((*parse.Tree)(nil)).Elem(),
		"VariableNode":   reflect.TypeOf((*parse.VariableNode)(nil)).Elem(),
		"WithNode":       reflect.TypeOf((*parse.WithNode)(nil)).Elem(),
	}
}
