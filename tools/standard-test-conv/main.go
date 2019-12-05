package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type assertion struct {
	Type      string
	FuncName  string
	Arguments []string
	Expected  string
}

var opts struct {
	InputFileName  string `short:"i" long:"input" required:"true" description:"Path to an input file"`
	OutputFileName string `short:"o" long:"output" description:"Path to an output file"`
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("error: %+v\n", err)
	}
}

func run() error {
	_, err := flags.Parse(&opts)
	if err != nil {
		return err
	}

	i, err := os.Open(opts.InputFileName)
	if err != nil {
		return err
	}
	defer i.Close()

	var o io.WriteCloser
	if opts.OutputFileName == "" {
		o = os.Stdout
	} else {
		o, err = os.Open(opts.OutputFileName)
		if err != nil {
			return err
		}
		defer o.Close()
	}

	return convert(i, o)
}

const header = `#!/bin/bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
wax="$d/cmd/wax/wax"
spec_core_test="$d/vendor/WebAssembly/spec/test/core"
source "$d/scripts/test_common.sh"
set -x

`

func convert(r io.Reader, w io.Writer) error {
	as, err := readAllAssertions(r)
	if err != nil {
		return err
	}

	ifn := replaceExt(filepath.Base(opts.InputFileName), ".wasm")

	fmt.Println(header)
	for _, a := range as {
		s, err := convertAssertion(a, ifn)
		if err != nil {
			return err
		}
		if s == "" {
			continue
		}
		w.Write([]byte(s + "\n"))
	}
	return nil
}

func readAllAssertions(r io.Reader) ([]*assertion, error) {
	result := make([]*assertion, 0)
	for {
		a, err := readOneAssertion(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if a != nil {
			result = append(result, a)
		}
	}
	return result, nil
}

func readOneAssertion(r io.Reader) (*assertion, error) {
	buf := bytes.NewBuffer([]byte{})
	b := make([]byte, 1)
	c := 0

	for {
		n, err := r.Read(b)
		if err != nil {
			return nil, err
		}
		if n != 1 {
			return nil, errors.New("unable to read a byte")
		}

		err = buf.WriteByte(b[0])
		if err != nil {
			return nil, err
		}

		if string(b) == "(" {
			c++
		} else if string(b) == ")" {
			c--
			if c == 0 {
				break
			}
		}
	}

	return parseAssertion(string(buf.Bytes()))
}

type state int

const (
	stateInit state = iota
	stateAssertionType
	stateInvoke
	stateArgs
	stateArgType
	stateArgValue
	stateResult
	stateResultType
	stateResultValue
	stateFinish
)

func parseAssertion(str string) (*assertion, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(str))
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '.' || ch == '_' || ch == '-' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}

	var t, f, at, r string
	args := []string{}

	state := stateInit

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tt := s.TokenText()
		switch state {
		case stateInit:
			if tt == "(" {
				state = stateAssertionType
			}
		case stateAssertionType:
			if tt == "module" || tt == "assert_malformed" || tt == "assert_invalid" {
				return nil, nil
			}
			if tt != "(" {
				t = tt
			} else {
				state = stateInvoke
			}
		case stateInvoke:
			if tt == "invoke" {
				// do nothing
			} else {
				f = strings.Trim(tt, `"`)
				state = stateArgs
			}
		case stateArgs:
			if tt == "(" {
				state = stateArgType
			} else if tt == ")" {
				state = stateResult
			}
		case stateArgType:
			at = tt
			state = stateArgValue
		case stateArgValue:
			if tt == ")" {
				state = stateArgs
			} else {
				args = append(args, at+" "+tt)
			}
		case stateResult:
			if t == "assert_trap" {
				r = strings.Trim(tt, `"`)
				state = stateFinish
			}
			if t == "assert_return" {
				if tt == "(" {
					state = stateResultType
				}
			}
		case stateResultType:
			r = tt
			state = stateResultValue
		case stateResultValue:
			if tt == ")" {
				state = stateFinish
			} else {
				r += " " + tt
			}
		}
	}
	return &assertion{
		Type:      t,
		FuncName:  f,
		Arguments: args,
		Expected:  r,
	}, nil
}

func convertAssertion(a *assertion, fname string) (string, error) {
	switch a.Type {
	case "assert_return":
		rt := getType(a.Expected)
		rvs := getValue(a.Expected)
		switch rt {
		case "i32":
			rv, err := parseI32(rvs)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf(`a="$( $wax -f "%s" %s "$spec_core_test/%s" )"; assert_equal "$a" "0:%s:%#08x %d %d"`, a.FuncName, argsString(a), fname, rt, rv, rv, int32(rv)), nil
		case "i64":
			rv, err := parseI64(rvs)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf(`a="$( $wax -f "%s" %s "$spec_core_test/%s" )"; assert_equal "$a" "0:%s:%#016x %d %d"`, a.FuncName, argsString(a), fname, rt, rv, rv, int64(rv)), nil
		case "f32":
			rv, err := parseF32(rvs)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf(`a="$( $wax -f "%s" %s "$spec_core_test/%s" )"; assert_equal "$a" "0:%s:%#08x %f"`, a.FuncName, argsString(a), fname, rt, math.Float32bits(rv), rv), nil
		case "f64":
			rv, err := parseF64(rvs)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf(`a="$( $wax -f "%s" %s "$spec_core_test/%s" )"; assert_equal "$a" "0:%s:%#08x %f"`, a.FuncName, argsString(a), fname, rt, math.Float64bits(rv), rv), nil
		default:
			return "", errors.New("not implemented yet")
		}
	case "assert_trap":
		return fmt.Sprintf(`set +e
a="$( $wax -f "%s" %s "$spec_core_test/%s" 2>&1 )"; assert_contains "$a" "%s"
set -e`, a.FuncName, argsString(a), fname, a.Expected), nil

	default:
		fmt.Fprintf(os.Stderr, "unknown assertion type %s\n", a.Type)
		return "", nil
	}
}

func replaceExt(fn, newExt string) string {
	ext := filepath.Ext(fn)
	if ext == "" {
		return fn + newExt
	}

	fn = fn[:len(fn)-len(ext)]
	return fn + newExt
}

func argsString(a *assertion) string {
	s := []string{}
	for _, arg := range a.Arguments {
		argType := getType(arg)
		argOpt := fmt.Sprintf(`-a "%s:%s"`, argType, getValue(arg))
		pad := 0
		switch argType {
		case "i32":
			pad = 19
		case "i64":
			pad = 27
		}
		argOpt = padSpace(argOpt, pad)
		s = append(s, argOpt)
	}
	return strings.Join(s, " ")
}

func padSpace(a string, n int) string {
	for len(a) < n {
		a += " "
	}
	return a
}

// "i32.const 1" => "i32"
func getType(s string) string {
	switch {
	case strings.HasPrefix(s, "i32."):
		return "i32"
	case strings.HasPrefix(s, "i64."):
		return "i64"
	case strings.HasPrefix(s, "f32."):
		return "f32"
	case strings.HasPrefix(s, "f64."):
		return "f64"
	default:
		return "unsupported"
	}
}

// "i32.const 1" => "1"
func getValue(s string) string {
	ss := strings.Split(s, " ")
	if len(ss) < 2 {
		return "--"
	}
	return ss[1]
}

func parseI32(s string) (uint32, error) {
	u, err := strconv.ParseUint(s, 0, 32)
	if err == nil {
		return uint32(u), nil
	}

	ne, ok := err.(*strconv.NumError)
	if !ok {
		return 0, err
	}
	if ne.Err != strconv.ErrSyntax { // ParseUint() fails for negative numbers such as "-123" with ErrSyntax. So we will try ParseInt() if the error is ErrSyntax
		return 0, err
	}

	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}

	return uint32(i), nil
}

func parseI64(s string) (uint64, error) {
	u, err := strconv.ParseUint(s, 0, 64)
	if err == nil {
		return uint64(u), nil
	}

	ne, ok := err.(*strconv.NumError)
	if !ok {
		return 0, err
	}
	if ne.Err != strconv.ErrSyntax { // ParseUint() fails for negative numbers such as "-123" with ErrSyntax. So we will try ParseInt() if the error is ErrSyntax
		return 0, err
	}

	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}

	return uint64(i), nil
}

func parseF32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

func parseF64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
