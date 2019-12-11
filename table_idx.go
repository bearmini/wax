package wax

type TableIdx uint32

func ParseTableIdx(ber *BinaryEncodingReader) (*TableIdx, []byte, error) {
	x, c, err := ber.ReadVaruint()
	if err != nil {
		return nil, nil, err
	}
	t := TableIdx(x)
	return &t, c, nil
}
