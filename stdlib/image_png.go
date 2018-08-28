package stdlib

// Generated by 'goexports image/png'. Do not edit!

import (
	"image/png"
	"reflect"
)

func init() {
	Value["image/png"] = map[string]reflect.Value{
		"BestCompression":    reflect.ValueOf(png.BestCompression),
		"BestSpeed":          reflect.ValueOf(png.BestSpeed),
		"Decode":             reflect.ValueOf(png.Decode),
		"DecodeConfig":       reflect.ValueOf(png.DecodeConfig),
		"DefaultCompression": reflect.ValueOf(png.DefaultCompression),
		"Encode":             reflect.ValueOf(png.Encode),
		"NoCompression":      reflect.ValueOf(png.NoCompression),
	}
	Type["image/png"] = map[string]reflect.Type{
		"CompressionLevel":  reflect.TypeOf((*png.CompressionLevel)(nil)).Elem(),
		"Encoder":           reflect.TypeOf((*png.Encoder)(nil)).Elem(),
		"EncoderBuffer":     reflect.TypeOf((*png.EncoderBuffer)(nil)).Elem(),
		"EncoderBufferPool": reflect.TypeOf((*png.EncoderBufferPool)(nil)).Elem(),
		"FormatError":       reflect.TypeOf((*png.FormatError)(nil)).Elem(),
		"UnsupportedError":  reflect.TypeOf((*png.UnsupportedError)(nil)).Elem(),
	}
}
