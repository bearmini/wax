package wax

import (
	"github.com/pkg/errors"
)

/*
	export     ::= nm:name d:exportdesc       => {name nm, desc d}
	exportdesc ::= 0x00 x:funcidx             => func x
							 | 0x01 x:tableidx            => table x
							 | 0x02 x:memidx              => mem x
							 | 0x03 x:globalidx           => global x
*/
type Export struct {
	Nm       Name
	DescType ExportDescType
	Desc     interface{}
}

func ParseExport(ber *BinaryEncodingReader) (*Export, error) {
	nm, _, err := ParseName(ber)
	if err != nil {
		return nil, err
	}

	descType, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}

	var desc interface{}
	switch ExportDescType(descType) {
	case ExportDescTypeFunc:
		x, _, err := ParseFuncIdx(ber)
		if err != nil {
			return nil, err
		}
		desc = x
	case ExportDescTypeTable:
		x, _, err := ParseTableIdx(ber)
		if err != nil {
			return nil, err
		}
		desc = x
	case ExportDescTypeMem:
		x, _, err := ParseMemIdx(ber)
		if err != nil {
			return nil, err
		}
		desc = x
	case ExportDescTypeGlobal:
		x, _, err := ParseGlobalIdx(ber)
		if err != nil {
			return nil, err
		}
		desc = x
	default:
		return nil, errors.New("unknown preceding flag for exportdesc")
	}

	return &Export{
		Nm:       *nm,
		DescType: ExportDescType(descType),
		Desc:     desc,
	}, nil
}
