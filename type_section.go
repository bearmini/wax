package wax

import (
	"bytes"

	"github.com/pkg/errors"
)

/*
Type Section
https://webassembly.github.io/multi-value/core/binary/modules.html#type-section

The type section has the id 1.
It decodes into a vector of function types that represent the types component of a module.

  typesec ::= ft*:section_1(vec(functype)) => ft*
*/
type TypeSection struct {
	SectionBase
	FuncTypes []FuncType
}

func ParseTypeSection(ber *BinaryEncodingReader, id SectionID) (*TypeSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	// Read Count
	count64, _, err := cr.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	count := uint32(count64)

	entries := make([]FuncType, 0, count)
	for i := uint32(0); i < count; i++ {
		ft, err := ParseFuncType(cr)
		if err != nil {
			return nil, err
		}
		entries = append(entries, *ft)
	}

	return &TypeSection{
		SectionBase: *sb,
		FuncTypes:   entries,
	}, nil
}

func (ts *TypeSection) GetFuncTypeAt(idx uint32) (*FuncType, error) {
	if idx > uint32(len(ts.FuncTypes)) {
		return nil, errors.New("out of range")
	}

	ft := ts.FuncTypes[idx]
	return &ft, nil
}
