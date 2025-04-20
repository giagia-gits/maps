package math

import (
	"reflect"
)

func Add(x, y interface{}) interface{} {
	vX := reflect.ValueOf(x)
	vY := reflect.ValueOf(y)
	tKind := InValid(vX.Type().Kind(), vY.Type().Kind())
	switch tKind {
	case 1:
		return vX.Interface().(int) + vY.Interface().(int)
	case 2:
		return vX.Interface().(uint) + vY.Interface().(uint)
	case 3:
		return vX.Interface().(float64) + vY.Interface().(float64)
	}
	return -1
}

func InValid(x, y reflect.Kind) int {
	switch x {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if x == y {
			return 1
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if x == y {
			return 2
		}
	case reflect.Float32, reflect.Float64:
		if x == y {
			return 3
		}
	}
	return -1
}
