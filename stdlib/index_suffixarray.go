package stdlib

// Generated by 'goexports index/suffixarray'. Do not edit!

import (
	"index/suffixarray"
	"reflect"
)

func init() {
	Value["index/suffixarray"] = map[string]reflect.Value{
		"New": reflect.ValueOf(suffixarray.New),
	}
	Type["index/suffixarray"] = map[string]reflect.Type{
		"Index": reflect.TypeOf((*suffixarray.Index)(nil)).Elem(),
	}
}
