package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/nsf/sexp"
)

func TestNextSexp(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  string
		Expected *sexp.Node
	}{
		{
			Name:     "pattern 1",
			Pattern:  ";; only a comment line",
			Expected: nil,
		},
		{
			Name:    "pattern 2 - only one assertion",
			Pattern: `(assert_return (invoke "add" (i32.const 1) (i32.const 1)) (i32.const 2))`,
			Expected: &sexp.Node{
				Children: &sexp.Node{
					Location: 0,
					Children: &sexp.Node{
						Location: 1,
						Value:    "assert_return",
						Next: &sexp.Node{
							Location: 15,
							Children: &sexp.Node{
								Location: 16,
								Value:    "invoke",
								Next: &sexp.Node{
									Location: 23,
									Value:    "add",
									Next: &sexp.Node{
										Location: 29,
										Children: &sexp.Node{
											Location: 30,
											Value:    "i32.const",
											Next: &sexp.Node{
												Location: 40,
												Value:    "1",
											},
										},
										Next: &sexp.Node{
											Location: 43,
											Children: &sexp.Node{
												Location: 44,
												Value:    "i32.const",
												Next: &sexp.Node{
													Location: 54,
													Value:    "1",
												},
											},
										},
									},
								},
							},
							Next: &sexp.Node{
								Location: 58,
								Children: &sexp.Node{
									Location: 59,
									Value:    "i32.const",
									Next: &sexp.Node{
										Location: 69,
										Value:    "2",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			r := NewWastReader(strings.NewReader(data.Pattern))
			a, err := r.NextSexp()
			if err != nil {
				t.Fatalf("unexpected error: %+v\n", err)
			}
			if !reflect.DeepEqual(data.Expected, a) {
				t.Fatalf("%s", pretty.Compare(data.Expected, a))
			}
		})
	}
}
