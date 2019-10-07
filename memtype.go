package wax

/*
Memory Types
https://webassembly.github.io/multi-value/core/binary/types.html#memory-types

Memory types are encoded with their limits.

	memtype ::= lim:limits => lim
*/
type MemType struct {
	Limits Limits
}

func ParseMemType(ber *BinaryEncodingReader) (*MemType, error) {
	l, err := ParseLimits(ber)
	if err != nil {
		return nil, err
	}

	return &MemType{
		Limits: *l,
	}, nil
}
