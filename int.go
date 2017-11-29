package cast

import (
	"strconv"
	"reflect"
	"fmt"
)

// CastInt accepts interface and returns any int or float type as int64. Parsable string values
// are accepted as well. If casting is not possible, second return parameter is false
func CastInt(val interface{}) (int64, bool) {
	switch val.(type) {
	case bool:
		if val.(bool) {
			return 1, true
		}
		return 0, true
	case int:
		return int64(val.(int)), true
	case int8:
		return int64(val.(int8)), true
	case int16:
		return int64(val.(int16)), true
	case int32:
		return int64(val.(int32)), true
	case int64:
		return val.(int64), true
	case uint:
		return int64(val.(uint)), true
	case uint8:
		return int64(val.(uint8)), true
	case uint16:
		return int64(val.(uint16)), true
	case uint32:
		return int64(val.(uint32)), true
	case uint64:
		return int64(val.(uint64)), true
	case float32:
		return int64(val.(float32)), true
	case float64:
		return int64(val.(float64)), true
	case string:
		return strToInt(val.(string))
	case struct{}, interface{}:
		if s, ok := val.(fmt.Stringer); ok {
			return strToInt(s.String())
		}
	}
	return 0, false
}

func strToInt(str string) (int64, bool) {
	if ival, err := strconv.Atoi(str); err == nil {
		return int64(ival), true
	} else if fval, err := strconv.ParseFloat(str, 64); err == nil {
		return int64(fval), true
	} else if bval, err := strconv.ParseBool(str); err == nil {
		if bval {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

// CastInts returns int slice, if input is int slice or slice of which each member can be casted to int.
func CastInts(val interface{}) []int64 {
	// fast out, if already there
	switch val.(type) {
	case []int64:
		return val.([]int64)
	}

	// must be slice, or else
	r := reflect.ValueOf(val)
	if r.Kind() != reflect.Slice {
		return nil
	}

	// build result & cast each item individually, or else
	res := make([]int64, r.Len())
	for i := 0; i < r.Len(); i++ {
		if ival, ok := CastInt(r.Index(i).Interface()); ok {
			res[i] = ival
		} else {
			return nil
		}
	}
	return res
}