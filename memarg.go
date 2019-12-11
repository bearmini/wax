package wax

type MemArg struct {
	Align  uint32
	Offset uint32
}

func ParseMemArg(ber *BinaryEncodingReader) (*MemArg, []byte, error) {
	a64, ca, err := ber.ReadVaruint()
	if err != nil {
		return nil, nil, err
	}
	o64, co, err := ber.ReadVaruint()
	if err != nil {
		return nil, nil, err
	}

	return &MemArg{
		Align:  uint32(a64),
		Offset: uint32(o64),
	}, append(ca, co...), nil
}
