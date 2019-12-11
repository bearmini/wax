package wax

import (
	"github.com/pkg/errors"
)

/*
init_expr

The encoding of an initializer expression is the normal encoding of the expression followed by the end opcode as a delimiter.

Note that get_global in an initializer expression can only refer to immutable imported globals and all uses of init_expr can only appear after the Imports section.
*/
type InitExpr []byte

func ParseInitExpr(ber *BinaryEncodingReader) ([]byte, error) {
	supportedOps := map[byte]string{
		0x23: "varuint32", // global.get
		0x41: "varint32",  // i32.const
		0x42: "varint64",  // i64.const
		0x43: "uint32",    // f32.const
		0x44: "uint64",    // f64.const
	}
	op := make([]byte, 1)
	ie := make([]byte, 0)
	for {
		_, err := ber.Read(op)
		if err != nil {
			return nil, err
		}
		opc := op[0]
		ie = append(ie, opc)

		if Opcode(opc) == OpcodeEnd {
			return ie, nil
		}

		t, ok := supportedOps[op[0]]
		if !ok {
			return nil, errors.Errorf("unsupported opcode in init_expr: %#02x", op[0])
		}

		switch t {
		case "varuint32":
			_, c, err := ber.ReadVaruint()
			if err != nil {
				return nil, err
			}
			ie = append(ie, c...)

		case "varint32":
			_, c, err := ber.ReadVarint()
			if err != nil {
				return nil, err
			}
			ie = append(ie, c...)

		case "varint64":
			_, c, err := ber.ReadVarint()
			if err != nil {
				return nil, err
			}
			ie = append(ie, c...)

		case "uint32":
			b := make([]byte, 4)
			n, err := ber.Read(b)
			if err != nil {
				return nil, err
			}
			if n != 4 {
				return nil, errors.New("insufficient data")
			}
			ie = append(ie, b...)

		case "uint64":
			b := make([]byte, 8)
			n, err := ber.Read(b)
			if err != nil {
				return nil, err
			}
			if n != 8 {
				return nil, errors.New("insufficient data")
			}
			ie = append(ie, b...)
		}
	}
}
