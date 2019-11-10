package main

import (
	"reflect"
	"testing"
)

func TestParseAssertion(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  string
		Expected assertion
	}{
		{
			Name:    "pattern 1",
			Pattern: `(assert_return (invoke "add" (i32.const 1) (i32.const 1)) (i32.const 2))`,
			Expected: assertion{
				Type:      "assert_return",
				FuncName:  "add",
				Arguments: []string{"i32.const 1", "i32.const 1"},
				Expected:  "i32.const 2",
			},
		},
		{
			Name:    "pattern 2",
			Pattern: `(assert_trap (invoke "div_s" (i32.const 1) (i32.const 0)) "integer divide by zero")`,
			Expected: assertion{
				Type:      "assert_trap",
				FuncName:  "div_s",
				Arguments: []string{"i32.const 1", "i32.const 0"},
				Expected:  "integer divide by zero",
			},
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a, err := parseAssertion(data.Pattern)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if !reflect.DeepEqual(data.Expected, *a) {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, *a)
			}
		})
	}
}

func TestReplaceExt(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  string
		NewExt   string
		Expected string
	}{
		{
			Name:     "pattern 1",
			Pattern:  "abc.def",
			NewExt:   ".ghi",
			Expected: "abc.ghi",
		},
		{
			Name:     "pattern 2",
			Pattern:  "abc.def.gz",
			NewExt:   ".txt",
			Expected: "abc.def.txt",
		},
		{
			Name:     "pattern 3",
			Pattern:  "abc.def",
			NewExt:   ".x",
			Expected: "abc.x",
		},
		{
			Name:     "pattern 4",
			Pattern:  "abc.def",
			NewExt:   "",
			Expected: "abc",
		},
		{
			Name:     "pattern 5",
			Pattern:  "abc.def",
			NewExt:   "ghi",
			Expected: "abcghi",
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a := replaceExt(data.Pattern, data.NewExt)
			if data.Expected != a {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, a)
			}
		})
	}
}

func TestArgsString(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  []string
		Expected string
	}{
		{
			Name:     "pattern 1",
			Pattern:  []string{"i32.const 1", "i32.const -1"},
			Expected: `-a "i32:1"          -a "i32:-1"        `,
		},
		{
			Name:     "pattern 2",
			Pattern:  []string{"i32.const 0x80000000", "i32.const 0xffffffff"},
			Expected: `-a "i32:0x80000000" -a "i32:0xffffffff"`,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a := argsString(&assertion{Arguments: data.Pattern})
			if data.Expected != a {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, a)
			}
		})
	}
}

func TestConvertAssertion(t *testing.T) {
	testData := []struct {
		Name     string
		Pattern  *assertion
		FileName string
		Expected string
	}{
		{
			Name: "pattern 1",
			Pattern: &assertion{
				Type:      "assert_return",
				FuncName:  "add",
				Arguments: []string{"i32.const -1", "i32.const -1"},
				Expected:  "i32.const -2",
			},
			FileName: "i32.wasm",
			Expected: `a="$( $wax -f "add" -a "i32:-1"         -a "i32:-1"         "$spec_core_test/i32.wasm" )"; assert_equal "$a" "0:i32:0xfffffffe 4294967294 -2"`,
		},
		{
			Name: "pattern 2",
			Pattern: &assertion{
				Type:      "assert_trap",
				FuncName:  "div_s",
				Arguments: []string{"i32.const 1", "i32.const 0"},
				Expected:  "integer divide by zero",
			},
			FileName: "i32.wasm",
			Expected: `set +e
a="$( $wax -f "div_s" -a "i32:1"          -a "i32:0"          "$spec_core_test/i32.wasm" 2>&1 )"; assert_contains "$a" "integer divide by zero"
set -e`,
		},
	}

	for _, data := range testData {
		data := data // capture
		t.Run(data.Name, func(t *testing.T) {
			//t.Parallel()

			a, err := convertAssertion(data.Pattern, data.FileName)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}
			if data.Expected != a {
				t.Fatalf("\nExpected: %+v\nActual:   %+v", data.Expected, a)
			}
		})
	}
}
