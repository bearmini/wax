package wax

type Elem struct {
	Table  TableIdx
	Offset Expr
	Init   []FuncIdx
}

func ParseElem(ber *BinaryEncodingReader) (*Elem, error) {
	// Read TableIdx
	tableIdx, _, err := ParseTableIdx(ber)
	if err != nil {
		return nil, err
	}

	// Read Expr
	expr, err := ParseExpr(ber)
	if err != nil {
		return nil, err
	}

	// Read the size of Init
	size64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	size := uint32(size64)

	fis := make([]FuncIdx, size)
	for i := uint32(0); i < size; i++ {
		fi, _, err := ParseFuncIdx(ber)
		if err != nil {
			return nil, err
		}
		fis[i] = *fi
	}

	return &Elem{
		Table:  *tableIdx,
		Offset: *expr,
		Init:   fis,
	}, nil
}
