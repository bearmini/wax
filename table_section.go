package wax

import (
	"bytes"
)

/*
Table Section
https://webassembly.github.io/multi-value/core/binary/modules.html#table-section

The table section has the id 4.
It decodes into a vector of tables that represent the tables component of a module.

	tablesec ::= tab*:section_4(vec(table)) => tab*
	table    ::= tt:tabletype               => {type tt}
*/
type TableSection struct {
	SectionBase
	TableTypes []*TableType
}

func ParseTableSection(ber *BinaryEncodingReader, id SectionID) (*TableSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	// Read Count
	count64, _, err := cr.ReadVaruint()
	if err != nil {
		return nil, err
	}
	count := uint32(count64)

	tableTypes := make([]*TableType, 0, count)
	for i := uint32(0); i < count; i++ {
		tt, err := ParseTableType(cr)
		if err != nil {
			return nil, err
		}
		tableTypes = append(tableTypes, tt)
	}

	return &TableSection{
		SectionBase: *sb,
		TableTypes:  tableTypes,
	}, nil
}
