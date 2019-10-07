package wax

type GlobalIdx uint32

func ParseGlobalIdx(ber *BinaryEncodingReader) (*GlobalIdx, []byte, error) {
	x, c, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, nil, err
	}
	g := GlobalIdx(x)
	return &g, c, nil
}
