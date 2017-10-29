package cast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCastFloat(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val float64
		to_ok  bool
	}{
		{nil, 0, false},
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
		{22.88, 22.88, true},
		{float32(22.88), 22.88, true},
		{float64(22.88), 22.88, true},
		{testStringer1, 123, true},
		{testStringer2, 0, false},
		{testStringer3, 0, false},
		{"123", 123, true},
		{"22.88", 22.88, true},
		{"aa", 0, false},
		{[]string{"aa"}, 0, false},
		{[]int{123}, 0, false},
	}
	for _, e := range expects {
		val, ok := CastFloat(e.from)
		assert.Equal(t, e.to_ok, ok, fmt.Sprintf("Expect %###v -> %v", e.from, e.to_ok))
		assert.True(t, e.to_val > val-0.0001 && e.to_val < val+0.0001, fmt.Sprintf("Expect %###v -> %.5f", e.from, e.to_val))
	}
}

func TestCastFloats(t *testing.T) {
	expects := []struct {
		from   interface{}
		to_val []float64
	}{
		{nil, nil},
		{12.5, nil},
		{[]float64{12.5}, []float64{12.5}},
		{[]int{12}, []float64{12}},
		{[]int8{12}, []float64{12}},
		{[]int16{12}, []float64{12}},
		{[]int32{12}, []float64{12}},
		{[]int64{12}, []float64{12}},
		{[]uint{12}, []float64{12}},
		{[]uint8{12}, []float64{12}},
		{[]uint16{12}, []float64{12}},
		{[]uint32{12}, []float64{12}},
		{[]uint64{12}, []float64{12}},
		{[]float32{12.5}, []float64{12.5}},
		{[]float64{12.5}, []float64{12.5}},
		{[]string{"12"}, []float64{12}},
		{[]string{"12.5"}, []float64{12.5}},
		{[]interface{}{12.5}, []float64{12.5}},
		{[]interface{}{"12.5"}, []float64{12.5}},
		{[]interface{}{testStringer1}, []float64{123}},
		{[]*testString{testStringer1}, []float64{123}},
		{[]interface{}{12.5, "12.5"}, []float64{12.5, 12.5}},
		{[]interface{}{12.5, "12.5", "a"}, nil},
	}
	for _, e := range expects {
		val := CastFloats(e.from)
		assert.Equal(t, e.to_val, val, fmt.Sprintf("Expect %###v -> %###v", e.from, e.to_val))
	}
}
