package wax

import (
	"math"
	"testing"
)

func TestParseNaN32WithPayload(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  string
		Expected uint32
	}{
		{
			Name:     "pattern 1 - +sNaN",
			Pattern:  "nan",
			Expected: 0x7fc00000,
		},
		{
			Name:     "pattern 2 - -sNaN",
			Pattern:  "-nan",
			Expected: 0xffc00000,
		},
		{
			Name:     "pattern 3 - qNaN",
			Pattern:  "nan:0x1",
			Expected: 0x7f800001,
		},
		{
			Name:     "pattern 4 - -qNaN",
			Pattern:  "-nan:0x1",
			Expected: 0xff800001,
		},
		{
			Name:     "pattern 5",
			Pattern:  "nan:0x200000",
			Expected: 0x7fa00000,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			v, err := ParseNaN32WithPayload(data.Pattern)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if math.Float32bits(v) != data.Expected {
				t.Fatalf("\nExpected: %#16x\nActual:   %#16x", data.Expected, math.Float32bits(v))
			}
		})
	}
}

func TestParseNaN64WithPayload(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  string
		Expected uint64
	}{
		{
			Name:     "pattern 1",
			Pattern:  "nan",
			Expected: 0x7ff8000000000000,
		},
		{
			Name:     "pattern 2",
			Pattern:  "-nan",
			Expected: 0xfff8000000000000,
		},
		{
			Name:     "pattern 3",
			Pattern:  "nan:0x1",
			Expected: 0x7ff0000000000001,
		},
		{
			Name:     "pattern 4",
			Pattern:  "-nan:0x1",
			Expected: 0xfff0000000000001,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			v, err := ParseNaN64WithPayload(data.Pattern)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if math.Float64bits(v) != data.Expected {
				t.Fatalf("\nExpected: %#16x\nActual:   %#16x", data.Expected, math.Float64bits(v))
			}
		})
	}
}
