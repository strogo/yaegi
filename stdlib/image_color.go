package stdlib

// Generated by 'goexports image/color'. Do not edit!

import (
	"image/color"
	"reflect"
)

func init() {
	Value["image/color"] = map[string]reflect.Value{
		"Alpha16Model": reflect.ValueOf(color.Alpha16Model),
		"AlphaModel":   reflect.ValueOf(color.AlphaModel),
		"Black":        reflect.ValueOf(color.Black),
		"CMYKModel":    reflect.ValueOf(color.CMYKModel),
		"CMYKToRGB":    reflect.ValueOf(color.CMYKToRGB),
		"Gray16Model":  reflect.ValueOf(color.Gray16Model),
		"GrayModel":    reflect.ValueOf(color.GrayModel),
		"ModelFunc":    reflect.ValueOf(color.ModelFunc),
		"NRGBA64Model": reflect.ValueOf(color.NRGBA64Model),
		"NRGBAModel":   reflect.ValueOf(color.NRGBAModel),
		"NYCbCrAModel": reflect.ValueOf(color.NYCbCrAModel),
		"Opaque":       reflect.ValueOf(color.Opaque),
		"RGBA64Model":  reflect.ValueOf(color.RGBA64Model),
		"RGBAModel":    reflect.ValueOf(color.RGBAModel),
		"RGBToCMYK":    reflect.ValueOf(color.RGBToCMYK),
		"RGBToYCbCr":   reflect.ValueOf(color.RGBToYCbCr),
		"Transparent":  reflect.ValueOf(color.Transparent),
		"White":        reflect.ValueOf(color.White),
		"YCbCrModel":   reflect.ValueOf(color.YCbCrModel),
		"YCbCrToRGB":   reflect.ValueOf(color.YCbCrToRGB),
	}
	Type["image/color"] = map[string]reflect.Type{
		"Alpha":   reflect.TypeOf((*color.Alpha)(nil)).Elem(),
		"Alpha16": reflect.TypeOf((*color.Alpha16)(nil)).Elem(),
		"CMYK":    reflect.TypeOf((*color.CMYK)(nil)).Elem(),
		"Color":   reflect.TypeOf((*color.Color)(nil)).Elem(),
		"Gray":    reflect.TypeOf((*color.Gray)(nil)).Elem(),
		"Gray16":  reflect.TypeOf((*color.Gray16)(nil)).Elem(),
		"Model":   reflect.TypeOf((*color.Model)(nil)).Elem(),
		"NRGBA":   reflect.TypeOf((*color.NRGBA)(nil)).Elem(),
		"NRGBA64": reflect.TypeOf((*color.NRGBA64)(nil)).Elem(),
		"NYCbCrA": reflect.TypeOf((*color.NYCbCrA)(nil)).Elem(),
		"Palette": reflect.TypeOf((*color.Palette)(nil)).Elem(),
		"RGBA":    reflect.TypeOf((*color.RGBA)(nil)).Elem(),
		"RGBA64":  reflect.TypeOf((*color.RGBA64)(nil)).Elem(),
		"YCbCr":   reflect.TypeOf((*color.YCbCr)(nil)).Elem(),
	}
}
