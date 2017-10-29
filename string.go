package cast

import (
	"fmt"
	"reflect"
)

// CastString accepts interface and returns any string and renders any int or float type as string.
// If casting is not possible, second return parameter is false
func CastString(val interface{}) (string, bool) {
	switch val.(type) {
	case string:
		return val.(string), true
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val), true
	case float32, float64:
		return fmt.Sprintf("%g", val), true
	case struct{}, interface{}:
		if s, ok := val.(fmt.Stringer); ok {
			return s.String(), true
		}
	}
	return "", false
}

// CastStrings returns string slice, if input is string slice or slice of which each member can be casted to string.
func CastStrings(val interface{}) []string {
	// fast out, if already there
	switch val.(type) {
	case []string:
		return val.([]string)
	}

	// must be slice, or else
	r := reflect.ValueOf(val)
	if r.Kind() != reflect.Slice {
		return nil
	}

	// build result & cast each item individually, or else
	res := make([]string, r.Len())
	for i := 0; i < r.Len(); i++ {
		if sval, ok := CastString(r.Index(i).Interface()); ok {
			res[i] = sval
		} else {
			return nil
		}
	}
	return res
}
