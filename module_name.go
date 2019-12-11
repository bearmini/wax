package wax

import (
	"github.com/pkg/errors"
)

type ModuleName struct {
	NameLen uint32
	NameStr string
}

func (mn *ModuleName) GetNameType() NameType {
	return NameTypeModule
}

func ParseModuleName(ber *BinaryEncodingReader) (*ModuleName, error) {
	l, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}

	nameBytes := make([]byte, l)
	n, err := ber.Read(nameBytes)
	if err != nil {
		return nil, err
	}
	if uint64(n) != l {
		return nil, errors.New("insufficient data")
	}

	return &ModuleName{
		NameLen: uint32(l),
		NameStr: string(nameBytes),
	}, nil
}
