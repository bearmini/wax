package wax

import (
	"bytes"
)

type ElementSection struct {
	SectionBase
	Elem []Elem
}

func ParseElementSection(ber *BinaryEncodingReader, id SectionID) (*ElementSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	// Read count of vector
	count64, _, err := cr.ReadVaruint()
	if err != nil {
		return nil, err
	}
	count := uint32(count64)

	elem := make([]Elem, 0, count)
	for i := uint32(0); i < count; i++ {
		e, err := ParseElem(cr)
		if err != nil {
			return nil, err
		}
		elem = append(elem, *e)
	}

	return &ElementSection{
		SectionBase: *sb,
		Elem:        elem,
	}, nil
}
