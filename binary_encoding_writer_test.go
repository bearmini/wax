package wax_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/bearmini/wax"
)

func TestBinaryEncodingWriterVarintN(t *testing.T) {
	testData := []struct {
		Name     string
		Value    int64
		Expected []byte
	}{
		{
			Name:     "pattern 1",
			Value:    1,
			Expected: []byte{0x01},
		},
		{
			Name:     "pattern 2",
			Value:    -1,
			Expected: []byte{0x7F},
		},
		{
			Name:     "pattern 3",
			Value:    -2,
			Expected: []byte{0x7E},
		},
		{
			Name:     "pattern 5",
			Value:    -624485,
			Expected: []byte{0x9B, 0xF1, 0x59},
		},
		/*
			{
				Name:     "pattern 4",
				Bytes:    []byte{0xFE, 0xFF, 0x7F},
				Consumed: []byte{0xFE, 0xFF, 0x7F},
				N:        16,
				Expected: -2,
			},
			{
				Name:     "pattern 6",
				Bytes:    []byte{0x80, 0x88, 0x80, 0x80, 0x00},
				Consumed: []byte{0x80, 0x88, 0x80, 0x80, 0x00},
				N:        32,
				Expected: 0x400,
			},
		*/
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			buf := bytes.NewBuffer([]byte{})
			bew := wax.NewBinaryEncodingWriter(buf)
			err := bew.WriteVarint(data.Value)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !reflect.DeepEqual(data.Expected, buf.Bytes()) {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, buf.Bytes())
			}
		})
	}
}
