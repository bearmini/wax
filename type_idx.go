package wax

// https://webassembly.github.io/multi-value/core/syntax/modules.html#syntax-typeidx
type TypeIdx uint32

func ParseTypeIdx(ber *BinaryEncodingReader) (*TypeIdx, []byte, error) {
	x, c, err := ber.ReadVaruint()
	if err != nil {
		return nil, nil, err
	}
	t := TypeIdx(x)
	return &t, c, nil
}
