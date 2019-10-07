package wax

import (
	"bytes"
)

/*
Function section

The function section has the id 3.
It decodes into a vector of type indices that represent the 𝗍𝗒𝗉𝖾 fields of the functions in the 𝖿𝗎𝗇𝖼𝗌 component of a module.
The locals and body fields of the respective functions are encoded separately in the code section.

  funcsec ::= x*:section_3(vec(typeidx)) => x*
*/
type FunctionSection struct {
	SectionBase
	Types []TypeIdx
}

func ParseFunctionSection(ber *BinaryEncodingReader, id SectionID) (*FunctionSection, error) {
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

	types := make([]TypeIdx, 0, count)
	for i := uint32(0); i < count; i++ {
		t, _, err := cr.ReadVaruintN(32)
		if err != nil {
			return nil, err
		}
		types = append(types, TypeIdx(t))
	}

	return &FunctionSection{
		SectionBase: *sb,
		Types:       types,
	}, nil
}
