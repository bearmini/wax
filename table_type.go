package wax

/*
Table Types
https://webassembly.github.io/multi-value/core/binary/types.html#table-types

Table types are encoded with their limits and a constant byte indicating their element type.

	tabletype ::= et:elemtype lim:limits => lim et
	elemtype  ::= 0x70                   => anyfunc
*/
type TableType struct {
	ElementType ElemType
	Limits      Limits
}

func ParseTableType(ber *BinaryEncodingReader) (*TableType, error) {
	et, err := ParseElemType(ber)
	if err != nil {
		return nil, err
	}

	l, err := ParseLimits(ber)
	if err != nil {
		return nil, err
	}

	return &TableType{
		ElementType: *et,
		Limits:      *l,
	}, nil
}
