package wax

type ElementSection struct {
	SectionBase
}

func ParseElementSection(ber *BinaryEncodingReader, id SectionID) (*ElementSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	return &ElementSection{
		SectionBase: *sb,
	}, nil
}
