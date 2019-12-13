package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bearmini/wax"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

var opts struct {
	FuncName    string   `short:"f" long:"function"    description:"function name to be executed"`
	FuncArgs    []string `short:"a" long:"arg"         description:"arguments to be passed to the specified function. \neg. -a i32:123 -a f64:1.23"`
	MaxExecTime int64    `short:"t" long:"exec-time"   description:"maximum execution time in seconds. 0 (default) for infinite."`
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	args, err := parseOptions()
	if err != nil {
		return err
	}

	for _, f := range args {
		var err error
		err = execute(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseOptions() ([]string, error) {
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		return nil, err
	}

	return args[1:], nil
}

func execute(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	m, err := wax.ParseBinaryModule(f)
	if err != nil {
		return err
	}

	rt, err := wax.NewRuntime(m, wax.NewRuntimeConfig())
	if err != nil {
		return err
	}

	fa, err := rt.FindFuncAddr(opts.FuncName)
	if err != nil {
		return err
	}

	val, err := parseFuncArgs(opts.FuncArgs)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if opts.MaxExecTime > 0 {
		ctxWithDeadline, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(opts.MaxExecTime)*time.Second))
		defer cancel()
		ctx = ctxWithDeadline
	}

	ret, err := rt.InvokeFunc(ctx, *fa, val)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Runtime status: %s\n", rt.Dump())
		return err
	}

	fmt.Printf("%s", wax.DumpVals(ret))
	return nil
}

func parseFuncArgs(args []string) ([]*wax.Val, error) {
	result := make([]*wax.Val, 0)

	for _, arg := range args {
		s := strings.Split(arg, ":")
		if len(s) < 2 {
			return nil, errors.Errorf("invalid arg format: `%s`", arg)
		}

		t := strings.ToLower(s[0])
		v := s[1]
		switch t {
		case "i32":
			vi, err := parseI32(v)
			if err != nil {
				return nil, errors.Errorf("invalid arg: `%s`", arg)
			}
			result = append(result, wax.NewValI32(uint32(vi)))
		case "i64":
			vi, err := parseI64(v)
			if err != nil {
				return nil, errors.Errorf("invalid arg: `%s`", arg)
			}
			result = append(result, wax.NewValI64(uint64(vi)))

		default:
			return nil, errors.Errorf("invalid arg: `%s`", arg)
		}
	}

	return result, nil
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
