package wax

import "github.com/pkg/errors"

/*
Limits

Limits are encoded with a preceding flag indicating whether a maximum is present.

	limits ::= 0x00 n:u32       => { min n, max ε }
	         | 0x01 n:u32 m:u32 => { min n, max m }
*/
type Limits struct {
	Min uint32
	Max *uint32
}

func ParseLimits(ber *BinaryEncodingReader) (*Limits, error) {
	flag, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}

	n64, _, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	n := uint32(n64)

	switch flag {
	case 0x00:
		return &Limits{
			Min: n,
		}, nil

	case 0x01:
		m64, _, err := ber.ReadVaruintN(32)
		if err != nil {
			return nil, err
		}
		m := uint32(m64)

		return &Limits{
			Min: n,
			Max: &m,
		}, nil

	default:
		return nil, errors.New("unknown preceding flag for limits")
	}
}
