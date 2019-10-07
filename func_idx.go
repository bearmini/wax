package wax

type FuncIdx uint32

func ParseFuncIdx(ber *BinaryEncodingReader) (*FuncIdx, []byte, error) {
	x, c, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, nil, err
	}
	f := FuncIdx(x)
	return &f, c, nil
}
