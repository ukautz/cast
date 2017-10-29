package cast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCastString(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val string
		to_ok  bool
	}{
		{123, "123", true},
		{-123, "-123", true},
		{int8(123), "123", true},
		{int16(123), "123", true},
		{int32(123), "123", true},
		{int64(123), "123", true},
		{uint(123), "123", true},
		{uint8(123), "123", true},
		{uint16(123), "123", true},
		{uint32(123), "123", true},
		{uint64(123), "123", true},
		{22.88, "22.88", true},
		{float32(22.88), "22.88", true},
		{float64(22.88), "22.88", true},
		{"123", "123", true},
		{"22.88", "22.88", true},
		{"aa", "aa", true},
		{testStringer1, "123", true},
		{testStringer2, "", false},
		{testStringer3, "", false},
		{[]string{"aa"}, "", false},
		{[]int{123}, "", false},
	}
	for _, e := range expects {
		val, ok := CastString(e.from)
		typ := reflect.TypeOf(e.from)
		assert.Equal(t, e.to_ok, ok, fmt.Sprintf("Expect (%s) %###v -> %v", typ, e.from, e.to_ok))
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect (%s) %###v -> %s", typ, e.from, e.to_val))
	}
}

func TestCastStrings(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val []string
	}{
		{nil, nil},
		{12.5, nil},
		{[]float64{12.5}, []string{"12.5"}},
		{[]int{12}, []string{"12"}},
		{[]int8{12}, []string{"12"}},
		{[]int16{12}, []string{"12"}},
		{[]int32{12}, []string{"12"}},
		{[]string{"12"}, []string{"12"}},
		{[]uint{12}, []string{"12"}},
		{[]uint8{12}, []string{"12"}},
		{[]uint16{12}, []string{"12"}},
		{[]uint32{12}, []string{"12"}},
		{[]uint64{12}, []string{"12"}},
		{[]float32{12.5}, []string{"12.5"}},
		{[]float64{12.5}, []string{"12.5"}},
		{[]string{"12"}, []string{"12"}},
		{[]string{"12.5"}, []string{"12.5"}},
		{[]interface{}{12.5}, []string{"12.5"}},
		{[]interface{}{"12.5"}, []string{"12.5"}},
		{[]interface{}{testStringer1}, []string{"123"}},
		{[]*testString{testStringer1}, []string{"123"}},
		{[]interface{}{testStringer2}, nil},
		{[]interface{}{12.5, "12.5"}, []string{"12.5", "12.5"}},
		{[]interface{}{12.5, "12.5", "a"}, []string{"12.5", "12.5", "a"}},
	}
	for _, e := range expects {
		val := CastStrings(e.from)
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect %###v -> %###v", e.from, e.to_val))
	}
}
