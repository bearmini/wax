package wax

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

/*
The module starts with a preamble of two fields:
| Field          | Type     | Description
| magic number   | uint32   | Magic number 0x6d736100 (i.e., ‘\0asm’)
| version        | uint32   | Version number, 0x1
*/
type Preamble struct {
	MagicNumber uint32
	Version     uint32
}

func (p *Preamble) Encode(w io.Writer) error {
	b := make([]byte, 8)

	binary.LittleEndian.PutUint32(b[0:], 0x6d736100)
	binary.LittleEndian.PutUint32(b[4:], 0x00000001)

	n, err := w.Write(b)
	if err != nil {
		return err
	}
	if n != 8 {
		return errors.New("unable to write all data")
	}

	return nil
}

func ParsePreamble(ber *BinaryEncodingReader) (*Preamble, error) {
	mn, err := ber.ReadU32LE()
	if err != nil {
		return nil, err
	}
	if mn != 0x6d736100 {
		return nil, errors.New("invalid magic number")
	}

	v, err := ber.ReadU32LE()
	if err != nil {
		return nil, err
	}
	if v != 0x00000001 {
		return nil, errors.Errorf("unsupported version: %#08x (%d)", v, v)
	}

	return &Preamble{
		MagicNumber: mn,
		Version:     v,
	}, nil
}
