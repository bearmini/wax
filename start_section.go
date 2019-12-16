package wax

import "bytes"

/*
Start Section
https://webassembly.github.io/multi-value/core/binary/modules.html#start-section

The start section has the id 8. It decodes into an optional start function that represents the start component of a module.

	startsec ::= st?:section_8(start) => st?
	start    ::= x:funcidx            => {func x}
*/
type StartSection struct {
	SectionBase
	Index FuncIdx
}

func ParseStartSection(ber *BinaryEncodingReader, id SectionID) (*StartSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	idx, _, err := ParseFuncIdx(cr)
	if err != nil {
		return nil, err
	}

	return &StartSection{
		SectionBase: *sb,
		Index:       *idx,
	}, nil
}
