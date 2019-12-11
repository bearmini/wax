package wax

/*
2.3.6 Table Types
Table types classify tables over elements of element types within a size range.

tabletype ::= limits elemtype
elemtype ::= funcref

5.3.6 Table Types
Table types are encoded with their limits and a constant byte indicating their element type.
tabletype ::= et:elemtype lim:limits ⇒ lim et
elemtype ::= 0x70 ⇒ funcref
*/
type TableType struct {
	Limits      Limits
	ElementType ElemType
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
		Limits:      *l,
		ElementType: *et,
	}, nil
}
