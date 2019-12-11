package wax

import (
	"bytes"
)

/*
The data section has the id 11.
It decodes into a vector of data segments that represent the data component of a module.

datasec ::= seg*:section11(vec(data)) -> seg
data    ::= x:memidx e:expr b*:vec(byte) -> {data x, offset e, init b*}
*/
type DataSection struct {
	SectionBase
	Data []Data
}

func ParseDataSection(ber *BinaryEncodingReader, id SectionID) (*DataSection, error) {
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

	data := make([]Data, 0, count)
	for i := uint32(0); i < count; i++ {
		d, err := ParseData(cr)
		if err != nil {
			return nil, err
		}
		data = append(data, *d)
	}

	return &DataSection{
		SectionBase: *sb,
		Data:        data,
	}, nil
}
