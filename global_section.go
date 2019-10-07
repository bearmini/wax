package wax

import (
	"bytes"
)

/*
Global Section
https://webassembly.github.io/multi-value/core/binary/modules.html#global-section

The global section has the id 6.
It decodes into a vector of globals that represent the globals component of a module.

	globalsec ::= glob*:section_6(vec(global)) => glob*
	global    ::= gt:globaltype e:expr         => {type gt, init e}
*/
type GlobalSection struct {
	SectionBase
	Globals []*Global
}

func ParseGlobalSection(ber *BinaryEncodingReader, id SectionID) (*GlobalSection, error) {
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

	globals := make([]*Global, 0, count)
	for i := uint32(0); i < count; i++ {
		g, err := ParseGlobal(cr)
		if err != nil {
			return nil, err
		}
		globals = append(globals, g)
	}

	return &GlobalSection{
		SectionBase: *sb,
		Globals:     globals,
	}, nil
}
