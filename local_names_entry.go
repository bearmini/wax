package wax

type LocalNamesEntry struct {
	Index    uint32
	LocalMap NameMap
}

func ParseLocalNamesEntry(ber *BinaryEncodingReader) (*LocalNamesEntry, error) {
	idx64, _, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}

	lm, err := ParseNameMap(ber)
	if err != nil {
		return nil, err
	}

	return &LocalNamesEntry{
		Index:    uint32(idx64),
		LocalMap: *lm,
	}, nil
}
