package wax

import (
	"github.com/pkg/errors"
)

/*
Result Types

The only result types occurring in the binary format are the types of blocks.
These are encoded in special compressed form, by either the byte 0x40 indicating the empty type or as a single value type.

blocktype ::= 0x40      => []
            | t:valtype => [t]

Note

In future versions of WebAssembly, this scheme may be extended to support multiple results or more general block types.

*/
type BlockType uint8

func ParseBlockType(ber *BinaryEncodingReader) (*BlockType, error) {
	b := make([]byte, 1)
	n, err := ber.Read(b)
	if err != nil {
		return nil, err
	}
	if n != 1 {
		return nil, errors.New("insufficient data")
	}

	bt := BlockType(b[0])
	return &bt, nil
}
