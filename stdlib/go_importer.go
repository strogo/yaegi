package stdlib

// Generated by 'goexports go/importer'. Do not edit!

import (
	"go/importer"
	"reflect"
)

func init() {
	Value["go/importer"] = map[string]reflect.Value{
		"Default": reflect.ValueOf(importer.Default),
		"For":     reflect.ValueOf(importer.For),
	}
	Type["go/importer"] = map[string]reflect.Type{
		"Lookup": reflect.TypeOf((*importer.Lookup)(nil)).Elem(),
	}
}
