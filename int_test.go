package cast

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestCastInt(t *testing.T) {
	expects := []struct{
		from interface{}
		to_val int64
		to_ok bool
	}{
		{123, 123, true},
		{-123, -123, true},
		{int8(123), 123, true},
		{int16(123), 123, true},
		{int32(123), 123, true},
		{int64(123), 123, true},
		{uint(123), 123, true},
		{uint8(123), 123, true},
		{uint16(123), 123, true},
		{uint32(123), 123, true},
		{uint64(123), 123, true},
		{22.88, 22, true},
		{float32(22.88), 22, true},
		{float64(22.88), 22, true},
		{testStringer1, 123, true},
		{testStringer2, 0, false},
		{testStringer3, 0, false},
		{"123", 123, true},
		{"22.88", 22, true},
		{"aa", 0, false},
		{[]string{"aa"}, 0, false},
		{[]int{123}, 0, false},
	}
	for _, e := range expects {
		val, ok := CastInt(e.from)
		assert.Equal(t, e.to_ok, ok, fmt.Sprintf("Expect %###v -> %v", e.from, e.to_ok))
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect %###v -> %d", e.from, e.to_val))
	}
}

func TestCastInts(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val []int64
	}{
		{nil, nil},
		{12.5, nil},
		{[]float64{12.5}, []int64{12}},
		{[]int{12}, []int64{12}},
		{[]int8{12}, []int64{12}},
		{[]int16{12}, []int64{12}},
		{[]int32{12}, []int64{12}},
		{[]int64{12}, []int64{12}},
		{[]uint{12}, []int64{12}},
		{[]uint8{12}, []int64{12}},
		{[]uint16{12}, []int64{12}},
		{[]uint32{12}, []int64{12}},
		{[]uint64{12}, []int64{12}},
		{[]float32{12.5}, []int64{12}},
		{[]float64{12.5}, []int64{12}},
		{[]string{"12"}, []int64{12}},
		{[]string{"12.5"}, []int64{12}},
		{[]interface{}{12.5}, []int64{12}},
		{[]interface{}{"12.5"}, []int64{12}},
		{[]interface{}{testStringer1}, []int64{123}},
		{[]*testString{testStringer1}, []int64{123}},
		{[]interface{}{12.5, "12.5"}, []int64{12, 12}},
		{[]interface{}{12.5, "12.5", "a"}, nil},
	}
	for _, e := range expects {
		val := CastInts(e.from)
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect %###v -> %###v", e.from, e.to_val))
	}
}