package main

import (
	"testing"
)

func TestTrimComment(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  string
		Expected string
	}{
		{
			Name:     "pattern 1",
			Pattern:  ";; should be trimmed all the line",
			Expected: "",
		},
		{
			Name:     "pattern 2",
			Pattern:  "abc ;; comment",
			Expected: "abc",
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a := trimComment(data.Pattern)
			if data.Expected != a {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, a)
			}
		})
	}
}
