package wax

import (
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

func (w *BinaryEncodingWriter) WriteVaruintN(n uint, v uint64) error {
	v &= uint64((1 << n) - 1)
	for {
		mask := uint64(0x7f)
		if n < 7 {
			mask = uint64((1 << n) - 1)
		}
		x := byte(v & mask)
		v >>= 7
		if v != 0 {
			x |= 0x80
		}
		written, err := w.Write([]byte{x})
		if err != nil {
			return err
		}
		if written != 1 {
			return errors.New("could not write 1 byte")
		}

		if n <= 7 || v == 0 {
			return nil
		}
		n -= 7
	}
}
