package cast

import (
	"strconv"
	"reflect"
	"fmt"
)

// CastFloat accepts interface and returns any int or float type as float64. Parsable string values
// are accepted as well. If casting is not possible, second return parameter is false
func CastFloat(val interface{}) (float64, bool) {
	switch val.(type) {
	case int:
		return float64(val.(int)), true
	case int8:
		return float64(val.(int8)), true
	case int16:
		return float64(val.(int16)), true
	case int32:
		return float64(val.(int32)), true
	case int64:
		return float64(val.(int64)), true
	case uint:
		return float64(val.(uint)), true
	case uint8:
		return float64(val.(uint8)), true
	case uint16:
		return float64(val.(uint16)), true
	case uint32:
		return float64(val.(uint32)), true
	case uint64:
		return float64(val.(uint64)), true
	case float32:
		return float64(val.(float32)), true
	case float64:
		return val.(float64), true
	case string:
		return strToFloat(val.(string))
	case struct{}, interface{}:
		if s, ok := val.(fmt.Stringer); ok {
			return strToFloat(s.String())
		}
	}
	return 0, false
}

func strToFloat(str string) (float64, bool) {
	if fval, err := strconv.ParseFloat(str, 64); err == nil {
		return fval, true
	}
	return 0, false
}

// CastFloats returns float slice, if input is float slice or slice of which each member can be casted to float.
func CastFloats(val interface{}) []float64 {
	// fast out, if already there
	switch val.(type) {
	case []float64:
		return val.([]float64)
	}

	// must be slice, or else
	r := reflect.ValueOf(val)
	if r.Kind() != reflect.Slice {
		return nil
	}

	// build result & cast each item individually, or else
	res := make([]float64, r.Len())
	for i := 0; i < r.Len(); i++ {
		if fval, ok := CastFloat(r.Index(i).Interface()); ok {
			res[i] = fval
		} else {
			return nil
		}
	}
	return res
}