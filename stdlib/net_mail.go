package stdlib

// Generated by 'goexports net/mail'. Do not edit!

import (
	"net/mail"
	"reflect"
)

func init() {
	Value["net/mail"] = map[string]reflect.Value{
		"ErrHeaderNotPresent": reflect.ValueOf(mail.ErrHeaderNotPresent),
		"ParseAddress":        reflect.ValueOf(mail.ParseAddress),
		"ParseAddressList":    reflect.ValueOf(mail.ParseAddressList),
		"ParseDate":           reflect.ValueOf(mail.ParseDate),
		"ReadMessage":         reflect.ValueOf(mail.ReadMessage),
	}
	Type["net/mail"] = map[string]reflect.Type{
		"Address":       reflect.TypeOf((*mail.Address)(nil)).Elem(),
		"AddressParser": reflect.TypeOf((*mail.AddressParser)(nil)).Elem(),
		"Header":        reflect.TypeOf((*mail.Header)(nil)).Elem(),
		"Message":       reflect.TypeOf((*mail.Message)(nil)).Elem(),
	}
}
