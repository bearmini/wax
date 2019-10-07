package wax

/*
 */
type LocalEntry struct {
	Count uint32
	Type  ValType
}

func ParseLocalEntry(ber *BinaryEncodingReader) (*LocalEntry, []byte, error) {
	consumedBytes := []byte{}
	count64, c, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, nil, err
	}
	count := uint32(count64)
	consumedBytes = append(consumedBytes, c...)

	t, c, err := ParseValType(ber)
	if err != nil {
		return nil, nil, err
	}
	consumedBytes = append(consumedBytes, c...)

	return &LocalEntry{
		Count: count,
		Type:  *t,
	}, consumedBytes, nil
}
