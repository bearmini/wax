package wax

import "github.com/pkg/errors"

/*
http://webassembly.github.io/spec/core/binary/values.html#names

Names

Names are encoded as a vector of bytes containing the Unicode (Section 3.9) UTF-8 encoding of the name’s code point sequence.

name ::= b*:vec(byte) => name(if utf8(name)=b*)
*/

type Name string

func ParseName(ber *BinaryEncodingReader) (*Name, []byte, error) {
	// Read NameLen
	nameLen64, c, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, nil, err
	}
	nameLen := uint32(nameLen64)

	// Read Name
	nameBytes := make([]byte, nameLen)
	n, err := ber.Read(nameBytes)
	if err != nil {
		return nil, nil, err
	}
	if uint32(n) != nameLen {
		return nil, nil, errors.New("insufficient data")
	}
	name := Name(nameBytes)

	return &name, append(c, nameBytes...), nil
}
