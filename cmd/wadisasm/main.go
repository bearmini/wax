package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bearmini/wax"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	FuncAddr *uint32 `short:"a" long:"funcaddr" description:"Specify func addr to disasm"`
	FuncName *string `short:"n" long:"funcname" description:"Specify func name to disasm"`
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

	for _, f := range args {
		err = disasm(f)
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

func disasm(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	m, err := wax.ParseBinaryModule(f)
	if err != nil {
		return err
	}

	if opts.FuncAddr != nil {
		fa := *opts.FuncAddr
		return disasmFuncAddr(m, fa)
	}

	if opts.FuncName != nil {
		fn := *opts.FuncName
		return disasmFuncName(m, fn)
	}

	return disasmAll(m)
}

func disasmFuncAddr(m *wax.Module, fa uint32) error {
	cs := m.GetCodeSection()
	if fa >= uint32(len(cs.Code)) {
		return errors.New("func addr is out of range")
	}
	c := cs.Code[fa]
	d, err := wax.Disassemble(c)
	if err != nil {
		return err
	}
	fmt.Printf("\nfunc:%d\n%s\n\n", fa, d)
	return nil
}

func disasmFuncName(m *wax.Module, fn string) error {
	rt, err := wax.NewRuntime(m, wax.NewRuntimeConfig())
	if err != nil {
		return err
	}

	fa, err := rt.FindFuncAddr(fn)
	if err != nil {
		return err
	}

	return disasmFuncAddr(m, uint32(*fa))
}

func disasmAll(m *wax.Module) error {
	cs := m.GetCodeSection()
	for fa := range cs.Code {
		err := disasmFuncAddr(m, uint32(fa))
		if err != nil {
			return err
		}
	}

	return nil
}
