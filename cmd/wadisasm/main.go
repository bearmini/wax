package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bearmini/wax"
	"github.com/jessevdk/go-flags"
)

var opts struct {
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

	cs := m.GetCodeSection()
	for i, c := range cs.Code {
		d, err := wax.Disassemble(c)
		if err != nil {
			return err
		}
		fmt.Printf("\nfunc:%d\n%s\n\n", i, d)
	}

	return nil
}
