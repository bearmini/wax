package wax

import (
	"math"
	"testing"
)

func TestParseHexfloat32(t *testing.T) {
	testData := []struct {
		Name     string
		Hexfloat string
		Expected float32
	}{
		{
			Name:     "pattern 1",
			Hexfloat: "0x1",
			Expected: 1,
		},
		{
			Name:     "pattern 2",
			Hexfloat: "0x1.1",
			Expected: 1.0625,
		},
		{
			Name:     "pattern 3",
			Hexfloat: "0xabc.def",
			Expected: 2748.870849609375,
		},
		{
			Name:     "pattern 4",
			Hexfloat: "0x1p1",
			Expected: 2,
		},
		{
			Name:     "pattern 5",
			Hexfloat: "0x1p+1",
			Expected: 2,
		},
		{
			Name:     "pattern 6",
			Hexfloat: "0x1p-1",
			Expected: 0.5,
		},
		{
			Name:     "pattern 7",
			Hexfloat: "0x1.1p-1",
			Expected: 0.53125,
		},
		{
			Name:     "pattern 8",
			Hexfloat: "-0x1p-149",
			Expected: math.Float32frombits(0x80000001),
		},
		{
			Name:     "pattern 9",
			Hexfloat: "-0x0p+0",
			Expected: math.Float32frombits(0x80000000),
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			v, err := ParseHexfloat32(data.Hexfloat)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if v != data.Expected {
				t.Fatalf("\nExpected: %#16x (%f, %e)\nActual:   %#16x (%f, %e)", math.Float32bits(data.Expected), data.Expected, data.Expected, math.Float32bits(v), v, v)
			}
		})
	}
}
func TestParseHexfloat64(t *testing.T) {
	testData := []struct {
		Name     string
		Hexfloat string
		Expected float64
	}{
		{
			Name:     "pattern 1",
			Hexfloat: "0x1",
			Expected: 1,
		},
		{
			Name:     "pattern 2",
			Hexfloat: "0x1.1",
			Expected: 1.0625,
		},
		{
			Name:     "pattern 3",
			Hexfloat: "0xabc.def",
			Expected: 2748.870849609375,
		},
		{
			Name:     "pattern 4",
			Hexfloat: "0x1p1",
			Expected: 2,
		},
		{
			Name:     "pattern 5",
			Hexfloat: "0x1p+1",
			Expected: 2,
		},
		{
			Name:     "pattern 6",
			Hexfloat: "0x1p-1",
			Expected: 0.5,
		},
		{
			Name:     "pattern 7",
			Hexfloat: "0x1.1p-1",
			Expected: 0.53125,
		},
		{
			Name:     "pattern 8",
			Hexfloat: "-0x1p-149",
			Expected: math.Float64frombits(0xb6a0000000000000),
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			v, err := ParseHexfloat64(data.Hexfloat)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if v != data.Expected {
				t.Fatalf("\nExpected: %#16x (%f)\nActual:   %#16x (%f)", math.Float64bits(data.Expected), data.Expected, math.Float64bits(v), v)
			}
		})
	}
}
