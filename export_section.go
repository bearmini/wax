package wax

import (
	"bytes"
)

/*
Export Section
https://webassembly.github.io/multi-value/core/binary/modules.html#export-section

The export section has the id 7. It decodes into a vector of exports that represent the 𝖾𝗑𝗉𝗈𝗋𝗍𝗌
component of a module.

	exportsec  ::= ex*:section_7(vec(export)) => ex*
	export     ::= nm:name d:exportdesc       => {name nm, desc d}
	exportdesc ::= 0x00 x:funcidx             => func x
							 | 0x01 x:tableidx            => table x
							 | 0x02 x:memidx              => mem x
							 | 0x03 x:globalidx           => global x
*/
type ExportSection struct {
	SectionBase
	Exports []*Export
}

func ParseExportSection(ber *BinaryEncodingReader, id SectionID) (*ExportSection, error) {
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

	exports := make([]*Export, 0, count)
	for i := uint32(0); i < count; i++ {
		e, err := ParseExport(cr)
		if err != nil {
			return nil, err
		}
		exports = append(exports, e)
	}

	return &ExportSection{
		SectionBase: *sb,
		Exports:     exports,
	}, nil
}
