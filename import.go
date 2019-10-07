package wax

import (
	"github.com/pkg/errors"
)

/*
import     ::= mod:name nm:name d:importdesc  => {module mod, name nm, desc d}
importdesc ::= 0x00 x:typeidx                 => func x
						 | 0x01 tt:tabletype              => table tt
						 | 0x02 mt:memtype                => mem mt
						 | 0x03 gt:globaltype             => global gt
*/
type Import struct {
	Mod      Name
	Nm       Name
	DescType ImportDescType
	Desc     interface{}
}

func ParseImport(ber *BinaryEncodingReader) (*Import, error) {
	mod, _, err := ParseName(ber)
	if err != nil {
		return nil, err
	}

	nm, _, err := ParseName(ber)
	if err != nil {
		return nil, err
	}

	descType, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}

	var desc interface{}
	switch ImportDescType(descType) {
	case ImportDescTypeFunc:
		x, _, err := ParseTypeIdx(ber)
		if err != nil {
			return nil, err
		}
		desc = x
	case ImportDescTypeTable:
		tt, err := ParseTableType(ber)
		if err != nil {
			return nil, err
		}
		desc = tt
	case ImportDescTypeMem:
		mt, err := ParseMemType(ber)
		if err != nil {
			return nil, err
		}
		desc = mt
	case ImportDescTypeGlobal:
		gt, err := ParseGlobalType(ber)
		if err != nil {
			return nil, err
		}
		desc = gt
	default:
		return nil, errors.New("unknown import desc type")
	}

	return &Import{
		Mod:      *mod,
		Nm:       *nm,
		DescType: ImportDescType(descType),
		Desc:     desc,
	}, nil
}
