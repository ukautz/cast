// cast is a library providing a sugar layer API for converting "unclean" (user) input into specific
// low level Go types.
//
// My primary use case: dealing with deep & complex YAML/JSON configuration files, which can
// contain string("123") numbers or annoying float64(0.0) numbers where int is expected and so on.
// See also github.com/ukautz/gpath to that end.
//
// This library is not optimized for performance but for simplifying interaction with user input - and keeping
// too often written, ugly, repetitive conditional type conversion code from my other projects.
package cast

import (
	"reflect"
)

var vof = reflect.ValueOf

// CastTo tries to cast given value into provided reflect.Kind. Supported are all reflect.Int*,
// all reflect.Uint*, all reflect.Float* and reflect.String. Second return parameter is false
// if casting is not possible.
//
//		s := "123"
//		if v, ok := CastTo(s, reflect.Int); ok {
//			fmt.Printf("Value is %d\n", v.(int))
//		}
func CastTo(value interface{}, kind reflect.Kind) (interface{}, bool) {
	if v := CastToValue(vof(value), kind); v != nil {
		return (*v).Interface(), true
	}
	return nil, false
}

// CastToValue works like CastTo but expects reflect.Value input and returns *reflect.Value or nil in case
// casting is not possible.
func CastToValue(value reflect.Value, kind reflect.Kind) *reflect.Value {
	switch kind {
	case reflect.Float32, reflect.Float64:
		if fval, ok := CastFloat(value.Interface()); ok && fval >= 0 {
			switch kind {
			case reflect.Float32:
				return vofp(float32(fval))
			case reflect.Float64:
				return vofp(fval)
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if ival, ok := CastInt(value.Interface()); ok {
			switch kind {
			case reflect.Int:
				return vofp(int(ival))
			case reflect.Int8:
				return vofp(int8(ival))
			case reflect.Int16:
				return vofp(int16(ival))
			case reflect.Int32:
				return vofp(int32(ival))
			case reflect.Int64:
				return vofp(ival)
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if ival, ok := CastInt(value.Interface()); ok && ival >= 0 {
			switch kind {
			case reflect.Uint:
				return vofp(uint(ival))
			case reflect.Uint8:
				return vofp(uint8(ival))
			case reflect.Uint16:
				return vofp(uint16(ival))
			case reflect.Uint32:
				return vofp(uint32(ival))
			case reflect.Uint64:
				return vofp(uint64(ival))
			}
		}
	case reflect.String:
		if sval, ok := CastString(value.Interface()); ok {
			return vofp(sval)
		}
	}
	return nil
}

func vofp(v interface{}) *reflect.Value {
	r := vof(v)
	return &r
}
