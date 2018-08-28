package stdlib

// Generated by 'goexports os/user'. Do not edit!

import (
	"os/user"
	"reflect"
)

func init() {
	Value["os/user"] = map[string]reflect.Value{
		"Current":       reflect.ValueOf(user.Current),
		"Lookup":        reflect.ValueOf(user.Lookup),
		"LookupGroup":   reflect.ValueOf(user.LookupGroup),
		"LookupGroupId": reflect.ValueOf(user.LookupGroupId),
		"LookupId":      reflect.ValueOf(user.LookupId),
	}
	Type["os/user"] = map[string]reflect.Type{
		"Group":               reflect.TypeOf((*user.Group)(nil)).Elem(),
		"UnknownGroupError":   reflect.TypeOf((*user.UnknownGroupError)(nil)).Elem(),
		"UnknownGroupIdError": reflect.TypeOf((*user.UnknownGroupIdError)(nil)).Elem(),
		"UnknownUserError":    reflect.TypeOf((*user.UnknownUserError)(nil)).Elem(),
		"UnknownUserIdError":  reflect.TypeOf((*user.UnknownUserIdError)(nil)).Elem(),
		"User":                reflect.TypeOf((*user.User)(nil)).Elem(),
	}
}
