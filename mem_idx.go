package wax

type MemIdx uint32

func ParseMemIdx(ber *BinaryEncodingReader) (*MemIdx, []byte, error) {
	x, c, err := ber.ReadVaruint()
	if err != nil {
		return nil, nil, err
	}
	m := MemIdx(x)
	return &m, c, nil
}
