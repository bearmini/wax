package wax_test

import (
	"testing"
)

func TestDisassemble(t *testing.T) {
	/*
		testData := []struct {
			Name     string
			Code     []byte
			Expected string
		}{
			{
				Name:     "pattern 1",
				Code:     []byte{0x20, 0x01, 0x20, 0x00, 0x6a},
				Expected: "20 01 local.get 00000001\\n20 00 local.get 00000000\\n6a    i32.add",
			},
		}

		for _, data := range testData {
			data := data // capture
			t.Run(data.Name, func(t *testing.T) {
				//t.Parallel()

				v, err := wax.Disassemble(data.Code)
				if err != nil {
					t.Fatalf("unexpected error: %+v", err)
				}
				if v != data.Expected {
					t.Fatalf("\nExpected: %s\nActual:   %s", data.Expected, v)
				}
			})
		}
	*/
}
