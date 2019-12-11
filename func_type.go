package wax

import "github.com/pkg/errors"

/*
func_type

Function types are encoded by the byte 0x60 followed by the respective vectors of parameter and result types.
functype ::= 0x60 t1*:vec(valtype) t2*:vec(valtype) => [t1*] → [t2*]
*/
type FuncType struct {
	ParamTypes  []ValType
	ReturnTypes []ValType
}

func ParseFuncType(ber *BinaryEncodingReader) (*FuncType, error) {
	fixed, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	if fixed != 0x60 {
		return nil, errors.Errorf("unexpected fixed value: %#02x", fixed)
	}

	pc64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	pc := uint32(pc64)

	pt := make([]ValType, 0)
	if pc > 0 {
		for i := uint32(0); i < pc; i++ {
			vt, _, err := ParseValType(ber)
			if err != nil {
				return nil, err
			}
			pt = append(pt, *vt)
		}
	}

	rc64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	rc := uint32(rc64)

	rt := make([]ValType, 0)
	if rc > 0 {
		for i := uint32(0); i < rc; i++ {
			vt, _, err := ParseValType(ber)
			if err != nil {
				return nil, err
			}
			rt = append(rt, *vt)
		}
	}

	return &FuncType{
		ParamTypes:  pt,
		ReturnTypes: rt,
	}, nil
}
