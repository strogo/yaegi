package stdlib

// Generated by 'goexports archive/tar'. Do not edit!

import (
	"archive/tar"
	"reflect"
)

func init() {
	Value["archive/tar"] = map[string]reflect.Value{
		"ErrFieldTooLong":    reflect.ValueOf(tar.ErrFieldTooLong),
		"ErrHeader":          reflect.ValueOf(tar.ErrHeader),
		"ErrWriteAfterClose": reflect.ValueOf(tar.ErrWriteAfterClose),
		"ErrWriteTooLong":    reflect.ValueOf(tar.ErrWriteTooLong),
		"FileInfoHeader":     reflect.ValueOf(tar.FileInfoHeader),
		"FormatGNU":          reflect.ValueOf(tar.FormatGNU),
		"FormatPAX":          reflect.ValueOf(tar.FormatPAX),
		"FormatUSTAR":        reflect.ValueOf(tar.FormatUSTAR),
		"FormatUnknown":      reflect.ValueOf(tar.FormatUnknown),
		"NewReader":          reflect.ValueOf(tar.NewReader),
		"NewWriter":          reflect.ValueOf(tar.NewWriter),
		"TypeBlock":          reflect.ValueOf(tar.TypeBlock),
		"TypeChar":           reflect.ValueOf(tar.TypeChar),
		"TypeCont":           reflect.ValueOf(tar.TypeCont),
		"TypeDir":            reflect.ValueOf(tar.TypeDir),
		"TypeFifo":           reflect.ValueOf(tar.TypeFifo),
		"TypeGNULongLink":    reflect.ValueOf(tar.TypeGNULongLink),
		"TypeGNULongName":    reflect.ValueOf(tar.TypeGNULongName),
		"TypeGNUSparse":      reflect.ValueOf(tar.TypeGNUSparse),
		"TypeLink":           reflect.ValueOf(tar.TypeLink),
		"TypeReg":            reflect.ValueOf(tar.TypeReg),
		"TypeRegA":           reflect.ValueOf(tar.TypeRegA),
		"TypeSymlink":        reflect.ValueOf(tar.TypeSymlink),
		"TypeXGlobalHeader":  reflect.ValueOf(tar.TypeXGlobalHeader),
		"TypeXHeader":        reflect.ValueOf(tar.TypeXHeader),
	}
	Type["archive/tar"] = map[string]reflect.Type{
		"Format": reflect.TypeOf((*tar.Format)(nil)).Elem(),
		"Header": reflect.TypeOf((*tar.Header)(nil)).Elem(),
		"Reader": reflect.TypeOf((*tar.Reader)(nil)).Elem(),
		"Writer": reflect.TypeOf((*tar.Writer)(nil)).Elem(),
	}
}
