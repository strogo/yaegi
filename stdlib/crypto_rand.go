package stdlib

// Generated by 'goexports crypto/rand'. Do not edit!

import (
	"crypto/rand"
	"reflect"
)

func init() {
	Value["crypto/rand"] = map[string]reflect.Value{
		"Int":    reflect.ValueOf(rand.Int),
		"Prime":  reflect.ValueOf(rand.Prime),
		"Read":   reflect.ValueOf(rand.Read),
		"Reader": reflect.ValueOf(rand.Reader),
	}
	Type["crypto/rand"] = map[string]reflect.Type{}
}