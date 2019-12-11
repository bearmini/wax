package wax

import (
	"github.com/pkg/errors"
)

type NameMap struct {
	Count   uint32
	Namings []*Naming
}

func (nm *NameMap) FindByName(name string) *Naming {
	for _, n := range nm.Namings {
		if n.NameStr == name {
			return n
		}
	}

	return nil
}

func ParseNameMap(ber *BinaryEncodingReader) (*NameMap, error) {
	count64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}

	namings := make([]*Naming, 0, count64)
	for i := uint64(0); i < count64; i++ {
		idx64, _, err := ber.ReadVaruint()
		if err != nil {
			return nil, err
		}
		nameLen64, _, err := ber.ReadVaruint()
		if err != nil {
			return nil, err
		}

		nameBytes := make([]byte, nameLen64)
		n, err := ber.Read(nameBytes)
		if err != nil {
			return nil, err
		}
		if uint64(n) != nameLen64 {
			return nil, errors.New("insufficient data")
		}

		naming := &Naming{
			Index:   uint32(idx64),
			NameLen: uint32(nameLen64),
			NameStr: string(nameBytes),
		}
		namings = append(namings, naming)
	}

	return &NameMap{
		Count:   uint32(count64),
		Namings: namings,
	}, nil
}
