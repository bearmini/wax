package wax

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type BinaryEncodingReader struct {
	r io.Reader
}

func NewBinaryEncodingReader(r io.Reader) *BinaryEncodingReader {
	return &BinaryEncodingReader{
		r: r,
	}
}

func (ber *BinaryEncodingReader) Read(b []byte) (int, error) {
	return ber.r.Read(b)
}

func (ber *BinaryEncodingReader) ReadU8() (uint8, error) {
	b := make([]byte, 1)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.New("unable to read 1 byte")
	}
	return b[0], nil
}

func (ber *BinaryEncodingReader) ReadU16LE() (uint16, error) {
	b := make([]byte, 2)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, errors.New("unable to read 2 bytes")
	}

	return binary.LittleEndian.Uint16(b), nil
}

func (ber *BinaryEncodingReader) ReadU16BE() (uint16, error) {
	b := make([]byte, 2)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, errors.New("unable to read 2 bytes")
	}

	return binary.BigEndian.Uint16(b), nil
}

func (ber *BinaryEncodingReader) ReadU32LE() (uint32, error) {
	b := make([]byte, 4)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 4 {
		return 0, errors.New("unable to read 4 bytes")
	}

	return binary.LittleEndian.Uint32(b), nil
}

func (ber *BinaryEncodingReader) ReadU32BE() (uint32, error) {
	b := make([]byte, 4)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 4 {
		return 0, errors.New("unable to read 4 bytes")
	}

	return binary.BigEndian.Uint32(b), nil
}

func (ber *BinaryEncodingReader) ReadU64LE() (uint64, error) {
	b := make([]byte, 8)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, errors.New("unable to read 8 bytes")
	}

	return binary.LittleEndian.Uint64(b), nil
}

func (ber *BinaryEncodingReader) ReadU64BE() (uint64, error) {
	b := make([]byte, 8)
	n, err := ber.r.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, errors.New("unable to read 8 bytes")
	}

	return binary.BigEndian.Uint64(b), nil
}

func (ber *BinaryEncodingReader) ReadVaruintN(n uint) (uint64, []byte, error) {
	if n > 64 {
		return 0, nil, errors.Errorf("%d is not supported", n)
	}

	var result uint64
	var shift uint
	consumedBytes := make([]byte, 0)

	for {
		b, err := ber.ReadU8()
		if err != nil {
			return 0, consumedBytes, err
		}

		consumedBytes = append(consumedBytes, b)
		result |= (uint64(b&0x7f) << shift)
		if b&0x80 == 0 {
			return result, consumedBytes, nil
		}

		shift += 7
		if n <= shift {
			return result, consumedBytes, nil
		}
	}
}

/*
Signed integers are encoded in signed LEB128 format, which uses a two’s complement representation.
As an additional constraint, the total number of bytes encoding a value of type sN must not exceed ceil(N/7) bytes.
*/
func (ber *BinaryEncodingReader) ReadVarintN(n uint) (int64, []byte, error) {
	if n > 64 {
		return 0, nil, errors.Errorf("%d is not supported", n)
	}

	var result int64
	var shift uint
	consumedBytes := make([]byte, 0)

	for {
		b, err := ber.ReadU8()
		if err != nil {
			return 0, consumedBytes, err
		}

		consumedBytes = append(consumedBytes, b)

		largeN := n - shift
		x := 1 << (largeN - 1) // 2^(N-1)

		switch {
		case b < 64 && int(b) < x:
			return (int64(b) << shift) | result, consumedBytes, nil
		case 64 <= b && b < 128 && int(b) >= (128-x):
			return ((int64(b) - 128) << shift) | result, consumedBytes, nil
		case 128 <= b && largeN > 7:
			result |= ((int64(b) - 128) << shift)
		default:
			return 0, nil, errors.New("something wrong")
		}

		shift += 7
		if shift >= n {
			return result, consumedBytes, nil
		}
	}
}
