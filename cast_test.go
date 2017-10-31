package cast

import (
	"reflect"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testString struct{}

func (s *testString) String() string {
	return "123"
}

var (
	testStringer1 = &testString{}
	testStringer2 = &testStringer1
	testStringer3 = &testStringer2
)

func TestCastTo(t *testing.T) {
	expects := []struct{
		from   interface{}
		kind   reflect.Kind
		expect interface{}
		ok     bool
	}{
		{int8(1), reflect.Int, int(1), true},
		{int(1), reflect.Int8, int8(1), true},
		{int(1), reflect.Int16, int16(1), true},
		{int(1), reflect.Int32, int32(1), true},
		{int(1), reflect.Int64, int64(1), true},
		{uint8(1), reflect.Uint, uint(1), true},
		{uint(1), reflect.Uint8, uint8(1), true},
		{uint(1), reflect.Uint16, uint16(1), true},
		{uint(1), reflect.Uint32, uint32(1), true},
		{uint(1), reflect.Uint64, uint64(1), true},
		{int8(1), reflect.Uint, uint(1), true},
		{int(1), reflect.Uint8, uint8(1), true},
		{int(1), reflect.Uint16, uint16(1), true},
		{int(1), reflect.Uint32, uint32(1), true},
		{int(1), reflect.Uint64, uint64(1), true},
		{float32(1), reflect.Uint64, uint64(1), true},
		{float64(1), reflect.Uint64, uint64(1), true},
		{float32(1), reflect.Float64, float64(1), true},
		{float64(1), reflect.Float32, float32(1), true},
		{"1.5", reflect.Float32, float32(1.5), true},
		{"1.5", reflect.Float64, float64(1.5), true},
		{"1.5", reflect.Int, int(1), true},
		{"1.5", reflect.Uint8, uint8(1), true},
		{"1.5", reflect.String, "1.5", true},
		{float32(1.5), reflect.String, "1.5", true},
		{int8(1), reflect.String, "1", true},
		{map[string]interface{}{}, reflect.Uint8, nil, false},
		{[]uint8{1}, reflect.Uint8, nil, false},
	}
	for _, e := range expects {
		is, ok := CastTo(e.from, e.kind)
		out := fmt.Sprintf("%###v (%s) with %s -> %###v (%s)", e.from, vof(e.from).Kind(), e.kind, is, vof(is).Kind())
		assert.Equal(t, e.ok, ok, "Expected %s to work", out)
		assert.Equal(t, e.expect, is, "Expected %s, but got %###v", out, is)
	}
}
