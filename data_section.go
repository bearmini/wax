package wax

type DataSection struct {
	SectionBase
}

func ParseDataSection(ber *BinaryEncodingReader, id SectionID) (*DataSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	return &DataSection{
		SectionBase: *sb,
	}, nil
}
