package main

import (
	"reflect"
	"testing"

	"github.com/bearmini/sexp"
	"github.com/bearmini/wax"
)

func TestEvalConst(t *testing.T) {
	testData := []struct {
		Name     string
		Sexp     *sexp.Sexp
		Expected wax.Val
	}{
		{
			Name:     "pattern 1",
			Sexp:     sexp.MustParse("(i32.const 0)"),
			Expected: wax.Val([]byte{0x41, 0x00}),
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a, err := evalConst(data.Sexp)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !reflect.DeepEqual(data.Expected, *a) {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, *a)
			}
		})
	}
}
