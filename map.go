package cast

import (
	"fmt"
	"reflect"
)

// CastMap accepts any map and casts its into the most generic map[interface{}]interface{} - second
// return parameter is false, if any key or value cannot be casted back to interface (which should not happen)
func CastMap(from interface{}) (map[interface{}]interface{}, bool) {
	t := createInterfaceType()
	if val, err := CastMapValue(vof(from), t, t); err == nil {
		return val.Interface().(map[interface{}]interface{}), true
	}
	return nil, false
}

// CastMapString accepts any map and casts its into map[string]interface{} - second return parameter
// is false, if any key cannot be casted to string
func CastMapString(from interface{}) (map[string]interface{}, bool) {
	kt := createStringType()
	vt := createInterfaceType()
	if val, err := CastMapValue(vof(from), kt, vt); err == nil {
		return val.Interface().(map[string]interface{}), true
	}
	return nil, false
}

// CastMapStringString accepts any map and casts its into map[string]string - second return parameter
// is false, if any key or value cannot be casted to string
func CastMapStringString(from interface{}) (map[string]string, bool) {
	t := createStringType()
	if val, err := CastMapValue(vof(from), t, t); err == nil {
		return val.Interface().(map[string]string), true
	}
	return nil, false
}

// CastMapStringInt accepts any map and casts its into map[string]int64 - second return parameter
// is false, if any key cannot be casted to string or any value cannot be casted into int64
func CastMapStringInt(from interface{}) (map[string]int64, bool) {
	kt := createStringType()
	vt := createIntType()
	if val, err := CastMapValue(vof(from), kt, vt); err == nil {
		return val.Interface().(map[string]int64), true
	}
	return nil, false
}

// CastMapStringInt accepts any map and casts its into map[string]float64 - second return parameter
// is false, if any key cannot be casted to string or any value cannot be casted into float64
func CastMapStringFloat(from interface{}) (map[string]float64, bool) {
	kt := createStringType()
	vt := createFloatType()
	if val, err := CastMapValue(vof(from), kt, vt); err == nil {
		return val.Interface().(map[string]float64), true
	}
	return nil, false
}

// CastMapStringBool accepts any map and casts its into map[string]bool - second return parameter
// is false, if any key cannot be casted to string or any value cannot be casted into bool
func CastMapStringBool(from interface{}) (map[string]bool, bool) {
	kt := createStringType()
	vt := createBoolType()
	if val, err := CastMapValue(vof(from), kt, vt); err == nil {
		return val.Interface().(map[string]bool), true
	}
	return nil, false
}

// CastMapValue accepts any map and casts its into map of provided key and value types - second
// return parameter contains error, if casting to given types is not possible
func CastMapValue(src reflect.Value, keyType, valType reflect.Type) (*reflect.Value, error) {
	return CastPartialMapValue(src, keyType, valType, true)
}

// CastPartialMapValue accepts any map and casts its into map of provided key and value types. The 4th
// (optional) input parameter controls whether a partial map of all the key/values which could be casted
// as requested (castFail = false or leave it out) is returned or whether only a map of provided types
// is returned if ALL keys and values can be casted (castFail = true)
func CastPartialMapValue(src reflect.Value, keyType, valType reflect.Type, castFail ...bool) (*reflect.Value, error) {
	failOnFail := len(castFail) > 0 && castFail[0]
	src = reflect.Indirect(src)
	if src.Kind() != reflect.Map {
		return nil, fmt.Errorf("can only cast from map kind to map kind, but got %s kind", src.Kind())
	}
	res := reflect.MakeMap(reflect.MapOf(keyType, valType))
	keyKind := keyType.Kind()
	valKind := valType.Kind()
	for _, srcKey := range src.MapKeys() {
		var resKey, resValue reflect.Value
		if keyKind == reflect.Interface {
			resKey = srcKey
		} else if resKeyP := CastToValue(srcKey, keyKind); resKeyP != nil {
			resKey = *resKeyP
		} else if failOnFail {
			return nil, fmt.Errorf("cannot cast key of kind %s to required kind %s", srcKey.Kind(), keyKind)
		} else {
			continue
		}
		if res.MapIndex(resKey).IsValid() {
			return nil, fmt.Errorf("cannot use duplicate key %###v", resKey)
		}
		srcVal := src.MapIndex(srcKey)
		if valKind == reflect.Interface {
			resValue = srcVal
		} else if resValP := CastToValue(srcVal, valKind); resValP != nil {
			resValue = *resValP
		} else if failOnFail {
			return nil, fmt.Errorf("cannot cast value of kind %s to required kind %s", srcVal.Kind(), valKind)
		} else {
			continue
		}
		res.SetMapIndex(resKey, resValue)
	}
	if !failOnFail && res.Len() == 0 && src.Len() > 0 {
		return nil, fmt.Errorf("could not cast any keys or values into requested %s key and % value type", keyType, valType)
	}
	return &res, nil
}

func createStringType() reflect.Type {
	return reflect.TypeOf((*string)(nil)).Elem()
}
func createIntType() reflect.Type {
	return reflect.TypeOf((*int64)(nil)).Elem()
}
func createFloatType() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}
func createBoolType() reflect.Type {
	return reflect.TypeOf((*bool)(nil)).Elem()
}
func createInterfaceType() reflect.Type {
	return reflect.TypeOf((*interface{})(nil)).Elem()
}
