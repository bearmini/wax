package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bearmini/wax"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

var opts struct {
	OutputFilename string `short:"o" long:"output" description:"output file name"`
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
		err = strip(f)
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

	if opts.OutputFilename != "" && len(args[1:]) != 1 {
		return nil, errors.New("output file name can be specified only when stripping one file")
	}

	return args[1:], nil
}

func strip(ifname string) error {
	b, err := ioutil.ReadFile(ifname)
	if err != nil {
		return err
	}

	m, err := wax.ParseBinaryModule(bytes.NewReader(b))
	if err != nil {
		return err
	}

	reduced := make([]wax.Section, 0, len(m.Sections))

	for _, section := range m.Sections {
		if section.GetID() == wax.CustomSectionID {
			cs := section.(*wax.CustomSection)
			if strings.HasPrefix(string(cs.Name), ".debug_") {
				continue
			}
		}

		reduced = append(reduced, section)
	}

	m.Sections = reduced

	ofname := ifname
	if opts.OutputFilename != "" {
		ofname = opts.OutputFilename
	}

	bb := bytes.NewBuffer([]byte{})
	err = m.EncodeToBinaryModule(bb)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ofname, bb.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
