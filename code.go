package wax

import (
	"bytes"

	"github.com/pkg/errors"
)

/*
	code    ::= size:u32 code:func          => code               (if size = ||func||)
	func    ::= (t*)*:vec(locals) e:expr    => concat((t*)*),e*   (if |concat((t*)*)| < 2^32)
	locals  ::= n:u32 t:valtype             => tn

Here, code ranges over pairs (valtype*,expr).
The meta function concat((t*)*) concatenates all sequences ti* in (t*)*.
*/
type Code struct {
	Size uint32
	Code Func
}

func ParseCode(ber *BinaryEncodingReader) (*Code, error) {
	// Read size
	size64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	size := uint32(size64)

	buf := make([]byte, size)
	n, err := ber.Read(buf)
	if err != nil {
		return nil, err
	}
	if uint32(n) != size {
		return nil, errors.New("insufficient data")
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(buf))
	fc, err := parseFuncInCodeSection(cr)
	if err != nil {
		return nil, err
	}

	f := Func{
		Type:   0, // This must be retrieved from 'function section' later
		Locals: concatLocals(fc.t),
		Body:   fc.e,
	}

	return &Code{
		Size: size,
		Code: f,
	}, nil
}

type funcInCodeSection struct {
	t []locals
	e Expr
}

func parseFuncInCodeSection(ber *BinaryEncodingReader) (*funcInCodeSection, error) {
	t, err := parseLocals(ber)
	if err != nil {
		return nil, err
	}
	e, err := ParseExpr(ber)
	if err != nil {
		return nil, err
	}

	return &funcInCodeSection{
		t: t,
		e: *e,
	}, nil
}

type locals struct {
	n uint32
	t ValType
}

func parseLocals(ber *BinaryEncodingReader) ([]locals, error) {
	// Read size of vector
	size64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	size := uint32(size64)

	result := make([]locals, 0, size)
	for i := uint32(0); i < size; i++ {
		// Read n
		n64, _, err := ber.ReadVaruint()
		if err != nil {
			return nil, err
		}
		n := uint32(n64)

		t, _, err := ParseValType(ber)
		if err != nil {
			return nil, err
		}
		result = append(result, locals{
			n: n,
			t: *t,
		})
	}
	return result, nil
}

func concatLocals(ls []locals) []ValType {
	result := make([]ValType, 0)

	for _, l := range ls {
		for i := uint32(0); i < l.n; i++ {
			result = append(result, l.t)
		}
	}

	return result
}
