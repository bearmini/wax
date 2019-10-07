package wax

import (
	"bytes"
)

/*
Memory Section
https://webassembly.github.io/multi-value/core/binary/modules.html#memory-section

The memory section has the id 5.
It decodes into a vector of memories that represent the mems component of a module.

	memsec ::= mem*:section_5(vec(mem)) => mem*
	mem    ::= mt:memtype               => {type mt}
*/
type MemorySection struct {
	SectionBase
	MemTypes []*MemType
}

func ParseMemorySection(ber *BinaryEncodingReader, id SectionID) (*MemorySection, error) {
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

	mem := make([]*MemType, 0, count)
	for i := uint32(0); i < count; i++ {
		mt, err := ParseMemType(cr)
		if err != nil {
			return nil, err
		}
		mem = append(mem, mt)
	}

	return &MemorySection{
		SectionBase: *sb,
		MemTypes:    mem,
	}, nil
}
