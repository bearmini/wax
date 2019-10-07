package wax_test

import (
	"bytes"
	"testing"

	"github.com/bearmini/wax"
)

func TestParseBlockType(t *testing.T) {
	testData := []struct {
		Name     string
		Bytes    []byte
		Expected wax.BlockType
	}{
		{
			Name:     "pattern 1",
			Bytes:    []byte{0x40},
			Expected: wax.BlockType(0x40),
		},
		{
			Name:     "pattern 2",
			Bytes:    []byte{0x7F},
			Expected: wax.BlockType(0x7F),
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			ber := wax.NewBinaryEncodingReader(bytes.NewReader(data.Bytes))
			v, err := wax.ParseBlockType(ber)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if *v != data.Expected {
				t.Fatalf("\nExpected: %#16x (%d)\nActual:   %#16x (%d)", data.Expected, data.Expected, v, v)
			}
		})
	}
}
