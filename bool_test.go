package cast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCastBool(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val bool
		to_ok  bool
	}{
		{nil, false, false},
		{123, true, true},
		{-123, true, true},
		{int8(123), true, true},
		{int16(123), true, true},
		{int32(123), true, true},
		{int64(123), true, true},
		{uint(123), true, true},
		{uint8(123), true, true},
		{uint16(123), true, true},
		{uint32(123), true, true},
		{uint64(123), true, true},
		{uint64(0), false, true},
		{22.88, true, true},
		{float32(22.88), true, true},
		{float64(22.88), true, true},
		{float64(0), false, true},
		{testStringer1, false, false},
		{testStringer2, false, false},
		{testStringer3, false, false},
		{"123",  true, true},
		{"22.88", true, true},
		{"aa", false, false},
		{"true", true, true},
		{"false", false, true},
		{"T", true, true},
		{"F", false, true},
		{"True", true, true},
		{"False", false, true},
		{"TRUE", true, true},
		{"FALSE", false, true},
		{[]string{"aa"}, false, false},
		{[]int{123}, false, false},
		{true, true, true},
		{false, false, true},
	}
	for _, e := range expects {
		val, ok := CastBool(e.from)
		assert.Equal(t, e.to_ok, ok, fmt.Sprintf("Expect ok %###v -> %v", e.from, e.to_ok))
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect val %###v -> %v", e.from, e.to_ok))
	}
}

func TestCastBools(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val []bool
	}{
		{nil, nil},
		{12.5, nil},
		{[]float64{12.5}, []bool{true}},
		{[]int{12}, []bool{true}},
		{[]int8{12}, []bool{true}},
		{[]int16{12}, []bool{true}},
		{[]int32{12}, []bool{true}},
		{[]int64{12}, []bool{true}},
		{[]uint{12}, []bool{true}},
		{[]uint8{12}, []bool{true}},
		{[]uint16{12}, []bool{true}},
		{[]uint32{12}, []bool{true}},
		{[]uint64{12}, []bool{true}},
		{[]float32{12.5}, []bool{true}},
		{[]float64{12.5}, []bool{true}},
		{[]string{"12"}, []bool{true}},
		{[]string{"12.5"}, []bool{true}},
		{[]bool{true, false, true, false}, []bool{true, false, true, false}},
		{[]interface{}{12.5}, []bool{true}},
		{[]interface{}{"12.5"}, []bool{true}},
		{[]interface{}{testStringer1}, nil},
		{[]*testString{testStringer1}, nil},
		{[]interface{}{12.5, "12.5"}, []bool{true, true}},
		{[]interface{}{12.5, "12.5", "a"}, nil},
	}
	for _, e := range expects {
		val := CastBools(e.from)
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect %###v -> %###v", e.from, e.to_val))
	}
}
