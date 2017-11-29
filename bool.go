package cast

import (
	"strconv"
	"reflect"
)

// CastBool accepts interface and returns bool type as bool. Non-zero int or float will be
// returned as true, zero int or float as false. Parsable string values (see: strconv.ParseBool)
// are accepted as well. If casting is not possible, second return parameter is false
func CastBool(val interface{}) (bool, bool) {
	switch val.(type) {
	case bool:
		return val.(bool), true
	case int:
		return val.(int) != 0, true
	case int8:
		return val.(int8) != 0, true
	case int16:
		return val.(int16) != 0, true
	case int32:
		return val.(int32) != 0, true
	case int64:
		return val.(int64) != 0, true
	case uint:
		return val.(uint) != 0, true
	case uint8:
		return val.(uint8) != 0, true
	case uint16:
		return val.(uint16) != 0, true
	case uint32:
		return val.(uint32) != 0, true
	case uint64:
		return val.(uint64) != 0, true
	case float32:
		return val.(float32) != 0, true
	case float64:
		return val.(float64) != 0, true
	case string:
		if bval, err := strconv.ParseBool(val.(string)); err != nil {
			if fval, ok := CastFloat(val.(string)); ok {
				return fval != 0, true
			}
			return false, false
		} else {
			return bval, true
		}
	}
	return false, false
}

// CastBools returns bool slice, if input is bool slice or slice of which each member can be casted to bool.
func CastBools(val interface{}) []bool {
	// fast out, if already there
	switch val.(type) {
	case []bool:
		return val.([]bool)
	}

	// must be slice, or else
	r := reflect.ValueOf(val)
	if r.Kind() != reflect.Slice {
		return nil
	}

	// build result & cast each item individually, or else
	res := make([]bool, r.Len())
	for i := 0; i < r.Len(); i++ {
		if fval, ok := CastBool(r.Index(i).Interface()); ok {
			res[i] = fval
		} else {
			return nil
		}
	}
	return res
}