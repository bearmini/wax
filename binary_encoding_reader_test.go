package wax_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/bearmini/wax"
)

func TestBinaryEncodingReaderVarintN(t *testing.T) {
	testData := []struct {
		Name     string
		Bytes    []byte
		Consumed []byte
		Expected int64
	}{
		{
			Name:     "pattern 1",
			Bytes:    []byte{0x01},
			Consumed: []byte{0x01},
			Expected: 1,
		},
		{
			Name:     "pattern 2",
			Bytes:    []byte{0x7F},
			Consumed: []byte{0x7F},
			Expected: -1,
		},
		{
			Name:     "pattern 3",
			Bytes:    []byte{0xFE, 0x7F},
			Consumed: []byte{0xFE, 0x7F},
			Expected: -2,
		},
		{
			Name:     "pattern 4",
			Bytes:    []byte{0xFE, 0xFF, 0x7F},
			Consumed: []byte{0xFE, 0xFF, 0x7F},
			Expected: -2,
		},
		{
			Name:     "pattern 5",
			Bytes:    []byte{0x9B, 0xF1, 0x59},
			Consumed: []byte{0x9B, 0xF1, 0x59},
			Expected: -624485,
		},
		{
			Name:     "pattern 6",
			Bytes:    []byte{0x80, 0x88, 0x80, 0x80, 0x00},
			Consumed: []byte{0x80, 0x88, 0x80, 0x80, 0x00},
			Expected: 0x400,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			ber := wax.NewBinaryEncodingReader(bytes.NewReader(data.Bytes))
			v, consumed, err := ber.ReadVarint()
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if v != data.Expected {
				t.Fatalf("\nExpected: %#16x (%d)\nActual:   %#16x (%d)", data.Expected, data.Expected, v, v)
			}
			if !reflect.DeepEqual(data.Consumed, consumed) {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Consumed, consumed)
			}
		})
	}
}
