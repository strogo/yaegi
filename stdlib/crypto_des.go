package stdlib

// Generated by 'goexports crypto/des'. Do not edit!

import (
	"crypto/des"
	"reflect"
)

func init() {
	Value["crypto/des"] = map[string]reflect.Value{
		"BlockSize":          reflect.ValueOf(des.BlockSize),
		"NewCipher":          reflect.ValueOf(des.NewCipher),
		"NewTripleDESCipher": reflect.ValueOf(des.NewTripleDESCipher),
	}
	Type["crypto/des"] = map[string]reflect.Type{
		"KeySizeError": reflect.TypeOf((*des.KeySizeError)(nil)).Elem(),
	}
}
