package wax

type LocalNames struct {
	Count uint32
	Funcs []*LocalNamesEntry
}

func (ln *LocalNames) GetNameType() NameType {
	return NameTypeLocal
}

func ParseLocalNames(ber *BinaryEncodingReader) (*LocalNames, error) {
	count64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}

	funcs := make([]*LocalNamesEntry, 0, count64)
	for i := uint64(0); i < count64; i++ {
		lne, err := ParseLocalNamesEntry(ber)
		if err != nil {
			return nil, err
		}
		funcs = append(funcs, lne)
	}

	return &LocalNames{
		Count: uint32(count64),
		Funcs: funcs,
	}, nil
}
