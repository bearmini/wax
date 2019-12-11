package wax

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

/*
Custom section name field: "name"

The name section is a custom section.
It is therefore encoded with id 0 followed by the name string "name".
Like all custom sections, this section being malformed does not cause the validation of the module to fail.
It is up to the implementation how it handles a malformed or partially malformed name section.
The WebAssembly implementation is also free to choose to read and process this section lazily,
after the module has been instantiated, should debugging be required.

The name section may appear only once, and only after the Data section.
The expectation is that, when a binary WebAssembly module is viewed in a browser or other development environment,
the data in this section will be used as the names of functions and locals in the text format.

The name section contains a sequence of name subsections:

| Field               | Type        | Description
| name_type           | varuint7    | code identifying type of name contained in this subsection
| name_payload_len    | varuint32   | size of this subsection in bytes
| name_payload_data   | bytes       | content of this section, of length name_payload_len
*/
type NameSection struct {
	CustomSection
	Subsections []NameSubsection
}

func ParseNameSection(cs *CustomSection) (*NameSection, error) {
	cr := NewBinaryEncodingReader(bytes.NewReader(cs.Content))

	subsections := make([]NameSubsection, 0)

	for {
		nt64, _, err := cr.ReadVaruint()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		nt := NameType(nt64)

		npLen64, _, err := cr.ReadVaruint()
		if err != nil {
			return nil, err
		}
		npLen := uint32(npLen64)

		npData := make([]byte, npLen)
		n, err := cr.Read(npData)
		if err != nil {
			return nil, err
		}
		if n != len(npData) {
			return nil, errors.New("insufficient data")
		}

		dr := NewBinaryEncodingReader(bytes.NewReader(npData))
		var nss NameSubsection
		switch nt {
		case NameTypeModule:
			nss, err = ParseModuleName(dr)
		case NameTypeFunction:
			nss, err = ParseFunctionNames(dr)
		case NameTypeLocal:
			nss, err = ParseLocalNames(dr)
		}
		if err != nil {
			return nil, err
		}
		subsections = append(subsections, nss)
	}

	return &NameSection{
		CustomSection: *cs,
		Subsections:   subsections,
	}, nil
}

func (ns *NameSection) GetFunctionSubsection() *FunctionNames {
	for _, nss := range ns.Subsections {
		if nss.GetNameType() == NameTypeFunction {
			fn, ok := nss.(*FunctionNames)
			if ok {
				return fn
			}
		}
	}

	return nil
}
