package math

import (
	"reflect"
)

func Add(x, y interface{}) interface{} {
	vX := reflect.ValueOf(x)
	vY := reflect.ValueOf(y)
	vXKind := vX.Type().Kind()
	vYKind := vY.Type().Kind()
	if vXKind != vYKind {
		return -1
	}
	switch vXKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if vXKind == reflect.Int {
			return vX.Interface().(int) + vY.Interface().(int)
		} else if vXKind == reflect.Int8 {
			return vX.Interface().(int8) + vY.Interface().(int8)
		} else if vXKind == reflect.Int16 {
			return vX.Interface().(int16) + vY.Interface().(int16)
		} else if vXKind == reflect.Int32 {
			return vX.Interface().(int32) + vY.Interface().(int32)
		} else if vXKind == reflect.Int64 {
			return vX.Interface().(int64) + vY.Interface().(int64)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if vXKind == reflect.Uint {
			return vX.Interface().(uint) + vY.Interface().(uint)
		} else if vXKind == reflect.Uint8 {
			return vX.Interface().(uint8) + vY.Interface().(uint8)
		} else if vXKind == reflect.Uint16 {
			return vX.Interface().(uint16) + vY.Interface().(uint16)
		} else if vXKind == reflect.Uint32 {
			return vX.Interface().(uint32) + vY.Interface().(uint32)
		} else if vXKind == reflect.Uint64 {
			return vX.Interface().(uint64) + vY.Interface().(uint64)
		}
	case reflect.Float32, reflect.Float64:
		if vXKind == reflect.Float32 {
			return vX.Interface().(float32) + vY.Interface().(float32)
		} else {
			return vX.Interface().(float64) + vY.Interface().(float64)
		}
	}
	return -1
}
