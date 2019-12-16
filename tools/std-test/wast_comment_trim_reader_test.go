package main

import (
	"testing"
)

func TestTrimComment(t *testing.T) {
	testData := []struct {
		Name         string
		Pattern      string
		InitialNest  int
		Expected     string
		ExpectedNest int
	}{
		{
			Name:         "pattern 1",
			Pattern:      ";; should be trimmed all the line",
			InitialNest:  0,
			Expected:     "",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 2",
			Pattern:      "abc ;; comment",
			InitialNest:  0,
			Expected:     "abc ",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 3 - block comment in a single line",
			Pattern:      "abc (; comment ;) def",
			InitialNest:  0,
			Expected:     "abc  def",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 4 - mixture of line comment and block comment 1",
			Pattern:      "abc (; ;;comment ;) def",
			InitialNest:  0,
			Expected:     "abc  def",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 5 - mixture of line comment and block comment 1",
			Pattern:      "abc ;; (; comment ;) def",
			InitialNest:  0,
			Expected:     "abc ",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 6 - open block comment",
			Pattern:      "abc (; comment",
			InitialNest:  0,
			Expected:     "abc ",
			ExpectedNest: 1,
		},
		{
			Name:         "pattern 7 - middle of block comment",
			Pattern:      "abc def",
			InitialNest:  1,
			Expected:     "",
			ExpectedNest: 1,
		},
		{
			Name:         "pattern 8 - end of block comment",
			Pattern:      "abc ;) def ;; ghi",
			InitialNest:  1,
			Expected:     " def ",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 9 - end of block comment and open again",
			Pattern:      "abc ;) def (; ghi",
			InitialNest:  1,
			Expected:     " def ",
			ExpectedNest: 1,
		},
		{
			Name:         "pattern 10 - nest",
			Pattern:      "abc (; def (; ghi ;) jkl ;) mno",
			InitialNest:  0,
			Expected:     "abc  mno",
			ExpectedNest: 0,
		},
		{
			Name:         "pattern 11 - nest and line comment",
			Pattern:      "abc (; def (; //ghi ;) jkl ;) mno",
			InitialNest:  0,
			Expected:     "abc  mno",
			ExpectedNest: 0,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a, nest := trimComments(data.Pattern, data.InitialNest)
			if data.Expected != a {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, a)
			}
			if data.ExpectedNest != nest {
				t.Fatalf("\nExpected: %d\nActual:   %d", data.ExpectedNest, nest)
			}
		})
	}
}
