package wax

import (
	"github.com/pkg/errors"
)

/*
external_kind

A single-byte unsigned integer indicating the kind of definition being imported or defined:

- 0 indicating a Function import or definition
- 1 indicating a Table import or definition
- 2 indicating a Memory import or definition
- 3 indicating a Global import or definition
*/
type ExternalKind uint8

const (
	ExternalKindFunction ExternalKind = iota
	ExternalKindTable
	ExternalKindMemory
	ExternalKindGlobal
)

func ParseExternalKind(ber *BinaryEncodingReader) (*ExternalKind, error) {
	b := make([]byte, 1)
	n, err := ber.Read(b)
	if err != nil {
		return nil, err
	}
	if n != 1 {
		return nil, errors.New("insufficient data")
	}

	ek := ExternalKind(b[0])
	return &ek, nil
}
