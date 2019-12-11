package wax

/*
elem_type

A varint7 indicating the types of elements in a table. In the MVP, only one type is available:

  -  anyfunc

Note: In the future :unicorn:, other element types may be allowed.
*/
type ElemType uint8

func ParseElemType(ber *BinaryEncodingReader) (*ElemType, error) {
	v, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}

	et := ElemType(v)
	return &et, nil
}
