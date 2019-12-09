package wax

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type BinaryEncodingWriter struct {
	w io.Writer
}

func NewBinaryEncodingWriter(w io.Writer) *BinaryEncodingWriter {
	return &BinaryEncodingWriter{
		w: w,
	}
}

func (w *BinaryEncodingWriter) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

func (w *BinaryEncodingWriter) WriteU8(n uint8) error {
	written, err := w.w.Write([]byte{n})
	if err != nil {
		return err
	}
	if written != 1 {
		return errors.New("could not write 1 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteU16BE(n uint16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, n)

	written, err := w.w.Write(b)
	if err != nil {
		return err
	}
	if written != 2 {
		return errors.New("could not write 2 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteU16LE(n uint16) error {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, n)

	written, err := w.w.Write(b)
	if err != nil {
		return err
	}
	if written != 2 {
		return errors.New("could not write 2 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteU32BE(n uint32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, n)

	written, err := w.w.Write(b)
	if err != nil {
		return err
	}
	if written != 4 {
		return errors.New("could not write 4 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteU32LE(n uint32) error {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, n)

	written, err := w.w.Write(b)
	if err != nil {
		return err
	}
	if written != 4 {
		return errors.New("could not write 4 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteU64BE(n uint64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)

	written, err := w.w.Write(b)
	if err != nil {
		return err
	}
	if written != 8 {
		return errors.New("could not write 8 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteU64LE(n uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)

	written, err := w.w.Write(b)
	if err != nil {
		return err
	}
	if written != 8 {
		return errors.New("could not write 8 byte")
	}
	return nil
}

func (w *BinaryEncodingWriter) WriteVaruint(v uint64) error {
	for {
		b := byte(v & 0x7f) // low order 7 bits of value;
		v >>= 7
		if v != 0 { // more bytes to come
			b |= 0x80 // set high order bit of byte;
		}
		err := w.WriteU8(b)
		if err != nil {
			return err
		}
		if v == 0 {
			return nil
		}
	}
}

func (w *BinaryEncodingWriter) WriteVarint(v int64) error {
	for {
		b := uint8(v)
		// Keep the sign bit for testing
		v >>= 6

		done := v == 0 || v == -1
		if done {
			b &= 0x7f
		} else {
			// Remove the sign bit
			v >>= 1
			// More bytes to come, so set the continuation bit.
			b |= 0x80
		}

		err := w.WriteU8(b)
		if err != nil {
			return err
		}

		if done {
			return nil
		}
	}
}
