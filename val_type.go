package wax

/*
Value Types
https://webassembly.github.io/multi-value/core/syntax/types.html#syntax-valtype

Value types classify the individual values that WebAssembly code can compute with and the values that a variable accepts.

valtype ::= i32 | i64 | f32 | f64

The types i32 and i64 classify 32 and 64 bit integers, respectively.
Integers are not inherently signed or unsigned, their interpretation is determined by individual operations.

The types f32 and f64 classify 32 and 64 bit floating-point data, respectively.
They correspond to the respective binary floating-point representations, also known as single and double precision, as defined by the IEEE 754-2008 standard (Section 3.3).

Conventions

- The meta variable t ranges over value types where clear from context.
- The notation |t| denotes the bit width of a value type. That is, |i32| = |f32| = 32 and |i64| = |f64| = 64.


Value Types
http://webassembly.github.io/spec/core/binary/types.html#value-types

Value types are encoded by a single byte.
valtype ::= 0x7F => i32
					| 0x7E => i64
					| 0x7D => f32
					| 0x7C => f64

Note:
In future versions of WebAssembly, value types may include types denoted by type indices.
Thus, the binary format for types corresponds to the signed LEB128 encoding of small negative sN
values, so that they can coexist with (positive) type indices in the future.
*/
type ValType uint8

const (
	ValTypeI32 ValType = 0x7f
	ValTypeI64 ValType = 0x7e
	ValTypeF32 ValType = 0x7d
	ValTypeF64 ValType = 0x7c
)

func (vt ValType) String() string {
	switch vt {
	case ValTypeI32:
		return "i32"
	case ValTypeI64:
		return "i64"
	case ValTypeF32:
		return "f32"
	case ValTypeF64:
		return "f64"
	default:
		return "unknown"
	}
}

func (vt ValType) BitCount() int {
	switch vt {
	case ValTypeI32, ValTypeF32:
		return 32
	case ValTypeI64, ValTypeF64:
		return 64
	default:
		return -1
	}
}

func ParseValType(ber *BinaryEncodingReader) (*ValType, []byte, error) {
	v, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, nil, err
	}

	vt := ValType(v)
	return &vt, []byte{byte(v)}, nil
}
