package cast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type _testBadType func() bool
type _testStringo struct {
	s string
}

var _testBadVar = _testBadType(func() bool { return false })

func (s *_testStringo) String() string {
	return fmt.Sprintf("stringo:%s", s.s)
}

func _testCastMap() map[interface{}]interface{} {

	return map[interface{}]interface{}{
		"a":              "b",
		"c":              1,
		int32(2):         3,
		uint(4):          "5",
		float64(6.5):     float64(1.5),
		interface{}(nil): &_testBadVar,
		uint8(9):         interface{}(nil),
	}
}

func TestCastPartialMapValue(t *testing.T) {
	res, err := CastPartialMapValue(vof(_testCastMap()), createInterfaceType(), createInterfaceType())
	assert.Nil(t, err, "Could be casted")
	assert.Equal(t, _testCastMap(), res.Interface(), "Result matches input")

	resm, ok := CastMap(res.Interface())
	assert.True(t, ok, "Valid map could be casted")
	assert.Equal(t, _testCastMap(), resm, "Result matches input")

	resm, ok = CastMap("str")
	assert.False(t, ok, "Invalid source should not be casted")
	resm, ok = CastMap(999)
	assert.False(t, ok, "Invalid source should not be casted")

	res, err = CastPartialMapValue(vof(_testCastMap()), createStringType(), createInterfaceType())
	assert.Nil(t, err, "Could be casted")
	assert.Equal(t, map[string]interface{}{"c": 1, "2": 3, "4": "5", "6.5": 1.5, "a": "b", "9": interface{}(nil)}, res.Interface(), "Result matches input")

	res, err = CastPartialMapValue(vof(_testCastMap()), createIntType(), createFloatType())
	assert.Nil(t, err, "Could be casted")
	assert.Equal(t, map[int64]float64{2: 3, 4: 5, 6: 1.5}, res.Interface(), "Result matches input")

	m := map[interface{}]interface{}{
		"stringo:foo":        "valid key",
		&_testStringo{"foo"}: "same string representation as above string",
	}
	assert.Equal(t, 2, len(m), "Yes, these arw two keys")
	res, err = CastPartialMapValue(vof(m), createStringType(), createInterfaceType())
	assert.NotNil(t, err, "Could be casted")
	assert.Equal(t, "cannot use duplicate key \"stringo:foo\"", err.Error(), "stringo, the foo thief")

	res, err = CastPartialMapValue(vof("str"), createIntType(), createFloatType())
	assert.NotNil(t, err, "Could be casted")
	assert.Equal(t, "can only cast from map kind to map kind, but got string kind", err.Error(), "Result matches input")
	res, err = CastPartialMapValue(vof(123), createIntType(), createFloatType())
	assert.NotNil(t, err, "Could be casted")
	assert.Equal(t, "can only cast from map kind to map kind, but got int kind", err.Error(), "Result matches input")

}

func TestCastMapString(t *testing.T) {
	res, err := CastPartialMapValue(vof(_testCastMap()), createStringType(), createInterfaceType())
	assert.Nil(t, err, "No casting error")
	assert.Equal(t, map[string]interface{}{"4": "5", "6.5": 1.5, "a": "b", "c": 1, "2": 3, "9": interface{}(nil)}, res.Interface(), "Casted matches assumption")
	resm, ok := CastMapString(res.Interface())
	assert.True(t, ok, "Valid map could be casted")
	assert.Equal(t, map[string]interface{}{"4": "5", "6.5": 1.5, "a": "b", "c": 1, "2": 3, "9": interface{}(nil)}, resm, "Result matches input")

	res, err = CastPartialMapValue(vof(_testCastMap()), createStringType(), reflect.TypeOf((*_testBadType)(nil)).Elem())
	assert.NotNil(t, err, "Should fail to cast with unsupported value type")
	assert.Contains(t, err.Error(), "could not cast any keys or values into requested", "Failed cause nothing was casted")

	m := _testCastMap()
	v := _testBadType(func() bool { return true })
	m["bad-apple"] = &v
	resm, ok = CastMapString(m)
	assert.False(t, ok, "Not entirely valid map could not be casted")
}

func TestCastMapStringString(t *testing.T) {
	res, err := CastPartialMapValue(vof(_testCastMap()), createStringType(), createStringType())
	assert.Nil(t, err, "No casting error")
	assert.Equal(t, map[string]string{"2": "3", "4": "5", "6.5": "1.5", "a": "b", "c": "1"}, res.Interface(), "Casted matches assumption")
	resm, ok := CastMapStringString(res.Interface())
	assert.True(t, ok, "Valid map could be casted")
	assert.Equal(t, map[string]string{"2": "3", "4": "5", "6.5": "1.5", "a": "b", "c": "1"}, resm, "Result matches input")
	resm, ok = CastMapStringString(_testCastMap())
	assert.False(t, ok, "Not entirely valid map could not be casted")
}

func TestCastMapStringInt(t *testing.T) {
	res, err := CastPartialMapValue(vof(_testCastMap()), createStringType(), createIntType())
	assert.Nil(t, err, "No casting error")
	assert.Equal(t, map[string]int64{"6.5": 1, "c": 1, "2": 3, "4": 5}, res.Interface(), "Casted matches assumption")
	resm, ok := CastMapStringInt(res.Interface())
	assert.True(t, ok, "Valid map could be casted")
	assert.Equal(t, map[string]int64{"6.5": 1, "c": 1, "2": 3, "4": 5}, resm, "Result matches input")
	resm, ok = CastMapStringInt(_testCastMap())
	assert.False(t, ok, "Not entirely valid map could not be casted")
}

func TestCastMapStringFloat(t *testing.T) {
	res, err := CastPartialMapValue(vof(_testCastMap()), createStringType(), createFloatType())
	assert.Nil(t, err, "No casting error")
	assert.Equal(t, map[string]float64{"4": 5, "6.5": 1.5, "c": 1, "2": 3}, res.Interface(), "Casted matches assumption")
	resm, ok := CastMapStringFloat(res.Interface())
	assert.True(t, ok, "Valid map could be casted")
	assert.Equal(t, map[string]float64{"4": 5, "6.5": 1.5, "c": 1, "2": 3}, resm, "Result matches input")
	resm, ok = CastMapStringFloat(_testCastMap())
	assert.False(t, ok, "Not entirely valid map could not be casted")
}

func TestCastMapStringBool(t *testing.T) {
	res, err := CastPartialMapValue(vof(_testCastMap()), createStringType(), createBoolType())
	assert.Nil(t, err, "No casting error")
	assert.Equal(t, map[string]bool{"4": true, "6.5": true, "c": true, "2": true}, res.Interface(), "Casted matches assumption")
	resm, ok := CastMapStringBool(res.Interface())
	assert.True(t, ok, "Valid map could be casted")
	assert.Equal(t, map[string]bool{"4": true, "6.5": true, "c": true, "2": true}, resm, "Result matches input")
	resm, ok = CastMapStringBool(_testCastMap())
	assert.False(t, ok, "Not entirely valid map could not be casted")
}
