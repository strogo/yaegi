package stdlib

// Generated by 'goexports encoding/ascii85'. Do not edit!

import (
	"encoding/ascii85"
	"reflect"
)

func init() {
	Value["encoding/ascii85"] = map[string]reflect.Value{
		"Decode":        reflect.ValueOf(ascii85.Decode),
		"Encode":        reflect.ValueOf(ascii85.Encode),
		"MaxEncodedLen": reflect.ValueOf(ascii85.MaxEncodedLen),
		"NewDecoder":    reflect.ValueOf(ascii85.NewDecoder),
		"NewEncoder":    reflect.ValueOf(ascii85.NewEncoder),
	}
	Type["encoding/ascii85"] = map[string]reflect.Type{
		"CorruptInputError": reflect.TypeOf((*ascii85.CorruptInputError)(nil)).Elem(),
	}
}
