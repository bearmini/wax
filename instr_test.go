package wax

import (
	"reflect"
	"testing"
)

func TestExtend(t *testing.T) {
	testData := []struct {
		Name       string
		Signedness string
		M          int
		ValType    ValType
		Val        *Val
		Expected   *Val
	}{
		{
			Name:       "Extend u - 8, 32",
			Signedness: "u",
			M:          8,
			ValType:    ValTypeI32,
			Val:        NewValI32(uint32(uint8(0xff))),
			Expected:   NewValI32(uint32(0x000000ff)),
		},
		{
			Name:       "Extend u - 16, 32",
			Signedness: "u",
			M:          16,
			ValType:    ValTypeI32,
			Val:        NewValI32(uint32(uint16(0xffff))),
			Expected:   NewValI32(uint32(0x0000ffff)),
		},
		{
			Name:       "Extend u - 8, 64",
			Signedness: "u",
			M:          8,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(uint8(0xff))),
			Expected:   NewValI64(uint64(0x000000ff)),
		},
		{
			Name:       "Extend u - 16, 64",
			Signedness: "u",
			M:          16,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(uint16(0xffff))),
			Expected:   NewValI64(uint64(0x000000000000ffff)),
		},
		{
			Name:       "Extend u - 32, 64",
			Signedness: "u",
			M:          32,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(0xffffffff)),
			Expected:   NewValI64(uint64(0x00000000ffffffff)),
		},
		{
			Name:       "Extend s - 8, 32 - no sign extension",
			Signedness: "s",
			M:          8,
			ValType:    ValTypeI32,
			Val:        NewValI32(uint32(uint8(0x7f))),
			Expected:   NewValI32(uint32(0x0000007f)),
		},
		{
			Name:       "Extend s - 8, 32 - sign extension",
			Signedness: "s",
			M:          8,
			ValType:    ValTypeI32,
			Val:        NewValI32(uint32(uint8(0xff))),
			Expected:   NewValI32(uint32(0xffffffff)),
		},
		{
			Name:       "Extend s - 16, 32 - no sign extension",
			Signedness: "s",
			M:          16,
			ValType:    ValTypeI32,
			Val:        NewValI32(uint32(uint16(0x7fff))),
			Expected:   NewValI32(uint32(0x0000007fff)),
		},
		{
			Name:       "Extend s - 16, 32 - sign extension",
			Signedness: "s",
			M:          16,
			ValType:    ValTypeI32,
			Val:        NewValI32(uint32(uint16(0xffff))),
			Expected:   NewValI32(uint32(0xffffffff)),
		},
		{
			Name:       "Extend s - 8, 64 - no sign extension",
			Signedness: "s",
			M:          8,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(uint8(0x7f))),
			Expected:   NewValI64(uint64(0x000000000000007f)),
		},
		{
			Name:       "Extend s - 8, 64 - sign extension",
			Signedness: "s",
			M:          8,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(uint8(0xff))),
			Expected:   NewValI64(uint64(0xffffffffffffffff)),
		},
		{
			Name:       "Extend s - 16, 64 - no sign extension",
			Signedness: "s",
			M:          16,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(uint16(0x7fff))),
			Expected:   NewValI64(uint64(0x0000000000007fff)),
		},
		{
			Name:       "Extend s - 16, 64 - sign extension",
			Signedness: "s",
			M:          16,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(uint16(0xffff))),
			Expected:   NewValI64(uint64(0xffffffffffffffff)),
		},
		{
			Name:       "Extend s - 32, 64 - no sign extension",
			Signedness: "s",
			M:          32,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(0x7fffffff)),
			Expected:   NewValI64(uint64(0x000000007fffffff)),
		},
		{
			Name:       "Extend s - 32, 64 - sign extension",
			Signedness: "s",
			M:          32,
			ValType:    ValTypeI64,
			Val:        NewValI32(uint32(0xffffffff)),
			Expected:   NewValI64(uint64(0xffffffffffffffff)),
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			actual, err := extend(data.Signedness, data.M, data.ValType, data.Val)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !reflect.DeepEqual(data.Expected, actual) {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, actual)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	testData := []struct {
		Name     string
		ValType  ValType
		N        int
		Val      *Val
		Expected *Val
	}{
		{
			Name:     "Wrap i32, i8",
			ValType:  ValTypeI32,
			N:        8,
			Val:      NewValI32(uint32(0xffffffff)),
			Expected: NewValI32(uint32(0x000000ff)),
		},
		{
			Name:     "Wrap i32, i16",
			ValType:  ValTypeI32,
			N:        16,
			Val:      NewValI32(uint32(0xffffffff)),
			Expected: NewValI32(uint32(0x0000ffff)),
		},
		{
			Name:     "Wrap i64, i8",
			ValType:  ValTypeI64,
			N:        8,
			Val:      NewValI64(uint64(0xffffffffffffffff)),
			Expected: NewValI32(uint32(0x000000ff)),
		},
		{
			Name:     "Wrap i64, i16",
			ValType:  ValTypeI64,
			N:        16,
			Val:      NewValI64(uint64(0xffffffffffffffff)),
			Expected: NewValI32(uint32(0x0000ffff)),
		},
		{
			Name:     "Wrap i64, i32",
			ValType:  ValTypeI64,
			N:        32,
			Val:      NewValI64(uint64(0xffffffffffffffff)),
			Expected: NewValI32(uint32(0xffffffff)),
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			actual, err := wrap(data.ValType, data.N, data.Val)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !reflect.DeepEqual(data.Expected, actual) {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, actual)
			}
		})
	}
}
