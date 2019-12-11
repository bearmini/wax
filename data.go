package wax

import (
	"github.com/pkg/errors"
)

/*
data    ::= x:memidx e:expr b*:vec(byte) -> {data x, offset e, init b*}
*/
type Data struct {
	Data   MemIdx
	Offset Expr
	Init   []byte
}

func ParseData(ber *BinaryEncodingReader) (*Data, error) {
	// Read MemIdx
	memIdx, _, err := ParseMemIdx(ber)
	if err != nil {
		return nil, err
	}

	// Read Expr
	expr, err := ParseExpr(ber)
	if err != nil {
		return nil, err
	}

	// Read the size of Value
	size64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	size := uint32(size64)

	buf := make([]byte, size)
	n, err := ber.Read(buf)
	if err != nil {
		return nil, err
	}
	if uint32(n) != size {
		return nil, errors.New("insufficient data")
	}

	return &Data{
		Data:   *memIdx,
		Offset: *expr,
		Init:   buf,
	}, nil
}
