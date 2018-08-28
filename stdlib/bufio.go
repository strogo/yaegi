package stdlib

// Generated by 'goexports bufio'. Do not edit!

import (
	"bufio"
	"reflect"
)

func init() {
	Value["bufio"] = map[string]reflect.Value{
		"ErrAdvanceTooFar":     reflect.ValueOf(bufio.ErrAdvanceTooFar),
		"ErrBufferFull":        reflect.ValueOf(bufio.ErrBufferFull),
		"ErrFinalToken":        reflect.ValueOf(bufio.ErrFinalToken),
		"ErrInvalidUnreadByte": reflect.ValueOf(bufio.ErrInvalidUnreadByte),
		"ErrInvalidUnreadRune": reflect.ValueOf(bufio.ErrInvalidUnreadRune),
		"ErrNegativeAdvance":   reflect.ValueOf(bufio.ErrNegativeAdvance),
		"ErrNegativeCount":     reflect.ValueOf(bufio.ErrNegativeCount),
		"ErrTooLong":           reflect.ValueOf(bufio.ErrTooLong),
		"MaxScanTokenSize":     reflect.ValueOf(bufio.MaxScanTokenSize),
		"NewReadWriter":        reflect.ValueOf(bufio.NewReadWriter),
		"NewReader":            reflect.ValueOf(bufio.NewReader),
		"NewReaderSize":        reflect.ValueOf(bufio.NewReaderSize),
		"NewScanner":           reflect.ValueOf(bufio.NewScanner),
		"NewWriter":            reflect.ValueOf(bufio.NewWriter),
		"NewWriterSize":        reflect.ValueOf(bufio.NewWriterSize),
		"ScanBytes":            reflect.ValueOf(bufio.ScanBytes),
		"ScanLines":            reflect.ValueOf(bufio.ScanLines),
		"ScanRunes":            reflect.ValueOf(bufio.ScanRunes),
		"ScanWords":            reflect.ValueOf(bufio.ScanWords),
	}
	Type["bufio"] = map[string]reflect.Type{
		"ReadWriter": reflect.TypeOf((*bufio.ReadWriter)(nil)).Elem(),
		"Reader":     reflect.TypeOf((*bufio.Reader)(nil)).Elem(),
		"Scanner":    reflect.TypeOf((*bufio.Scanner)(nil)).Elem(),
		"SplitFunc":  reflect.TypeOf((*bufio.SplitFunc)(nil)).Elem(),
		"Writer":     reflect.TypeOf((*bufio.Writer)(nil)).Elem(),
	}
}
