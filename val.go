package wax

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"

	"github.com/pkg/errors"
)

/*
Values
https://webassembly.github.io/multi-value/core/exec/runtime.html#syntax-val

WebAssembly computations manipulate values of the four basic value types: integers and floating-point data of 32 or 64 bit width each, respectively.

In most places of the semantics, values of different types can occur.
In order to avoid ambiguities, values are therefore represented with an abstract syntax that makes their type explicit.
It is convenient to reuse the same notation as for the 'const' instructions producing them:

val ::= i32.const i32
			| i64.const i64
			| f32.const f32
			| f64.const f64

*/
type Val []byte

func NewVal(t ValType, v interface{}) (*Val, error) {
	switch t {
	case ValTypeI32:
		x, ok := v.(uint32)
		if !ok {
			return nil, errors.New("invalid type")
		}
		return NewValI32(x), nil
	case ValTypeI64:
		x, ok := v.(uint64)
		if !ok {
			return nil, errors.New("invalid type")
		}
		return NewValI64(x), nil
	case ValTypeF32:
		x, ok := v.(float32)
		if !ok {
			return nil, errors.New("invalid type")
		}
		return NewValF32(x), nil
	case ValTypeF64:
		x, ok := v.(float64)
		if !ok {
			return nil, errors.New("invalid type")
		}
		return NewValF64(x), nil
	default:
		return nil, errors.New("invalid val type")
	}
}

func NewZeroVal(t ValType) (*Val, error) {
	switch t {
	case ValTypeI32:
		return NewValI32(0), nil
	case ValTypeI64:
		return NewValI64(0), nil
	case ValTypeF32:
		return NewValF32(0.0), nil
	case ValTypeF64:
		return NewValF64(0.0), nil
	default:
		return nil, errors.New("invalid val type")
	}
}

func NewValI32(v uint32) *Val {
	b := make([]byte, 0, 5)
	buf := bytes.NewBuffer(b)
	bew := NewBinaryEncodingWriter(buf)
	err := bew.WriteU8(byte(OpcodeI32Const))
	if err != nil {
		panic(err)
	}
	err = bew.WriteVaruint(uint64(v))
	if err != nil {
		panic(err)
	}

	result := Val(buf.Bytes())
	return &result
}

func NewValI64(v uint64) *Val {
	b := make([]byte, 0, 9)
	buf := bytes.NewBuffer(b)
	bew := NewBinaryEncodingWriter(buf)
	err := bew.WriteU8(byte(OpcodeI64Const))
	if err != nil {
		panic(err)
	}
	err = bew.WriteVaruint(v)
	if err != nil {
		panic(err)
	}

	result := Val(buf.Bytes())
	return &result
}

func NewValF32(v float32) *Val {
	b := make([]byte, 5)
	b[0] = byte(OpcodeF32Const)

	binary.LittleEndian.PutUint32(b[1:], math.Float32bits(v))

	result := Val(b)
	return &result
}

func NewValF64(v float64) *Val {
	b := make([]byte, 9)
	b[0] = byte(OpcodeF64Const)

	binary.LittleEndian.PutUint64(b[1:], math.Float64bits(v))

	result := Val(b)
	return &result
}

func (v *Val) GetI32() (uint32, error) {
	ber := NewBinaryEncodingReader(bytes.NewReader(*v))
	opcode, err := ber.ReadU8()
	if err != nil {
		return 0, err
	}

	if Opcode(opcode) != OpcodeI32Const {
		return 0, errors.New("type mismatch")
	}

	val, _, err := ber.ReadVaruintN(32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func (v *Val) GetI64() (uint64, error) {
	ber := NewBinaryEncodingReader(bytes.NewReader(*v))
	opcode, err := ber.ReadU8()
	if err != nil {
		return 0, err
	}

	if Opcode(opcode) != OpcodeI64Const {
		return 0, errors.New("type mismatch")
	}

	val, _, err := ber.ReadVaruintN(64)
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

func (v *Val) GetF32() (float32, error) {
	ber := NewBinaryEncodingReader(bytes.NewReader(*v))
	opcode, err := ber.ReadU8()
	if err != nil {
		return 0, err
	}

	if Opcode(opcode) != OpcodeF32Const {
		return 0, errors.New("type mismatch")
	}

	val, err := ber.ReadU32LE()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(val), nil
}

func (v *Val) GetF64() (float64, error) {
	ber := NewBinaryEncodingReader(bytes.NewReader(*v))
	opcode, err := ber.ReadU8()
	if err != nil {
		return 0, err
	}

	if Opcode(opcode) != OpcodeF64Const {
		return 0, errors.New("type mismatch")
	}

	val, err := ber.ReadU64LE()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(val), nil
}

func (v *Val) MustGetI32() uint32 {
	u, err := v.GetI32()
	if err != nil {
		panic(err)
	}
	return u
}

func (v *Val) MustGetI64() uint64 {
	u, err := v.GetI64()
	if err != nil {
		panic(err)
	}
	return u
}

func (v *Val) MustGetF32() float32 {
	u, err := v.GetF32()
	if err != nil {
		panic(err)
	}
	return u
}

func (v *Val) MustGetF64() float64 {
	u, err := v.GetF64()
	if err != nil {
		panic(err)
	}
	return u
}

func (v *Val) GetType() (*ValType, error) {
	if len(*v) < 1 {
		return nil, errors.New("invalid val type")
	}

	switch Opcode((*v)[0]) {
	case OpcodeI32Const:
		return &[]ValType{ValTypeI32}[0], nil
	case OpcodeI64Const:
		return &[]ValType{ValTypeI64}[0], nil
	case OpcodeF32Const:
		return &[]ValType{ValTypeF32}[0], nil
	case OpcodeF64Const:
		return &[]ValType{ValTypeF64}[0], nil
	default:
		return nil, errors.New("unknown val type")
	}
}

func (v *Val) EqualsTo(x *Val) bool {
	vt1, err := v.GetType()
	if err != nil {
		return false
	}

	vt2, err := x.GetType()
	if err != nil {
		return false
	}

	if *vt1 != *vt2 {
		return false
	}

	switch *vt1 {
	case ValTypeI32:
		v1 := v.MustGetI32()
		v2 := x.MustGetI32()
		return v1 == v2
	case ValTypeI64:
		v1 := v.MustGetI64()
		v2 := x.MustGetI64()
		return v1 == v2
	case ValTypeF32:
		v1 := v.MustGetF32()
		v2 := x.MustGetF32()
		if math.IsNaN(float64(v1)) && math.IsNaN(float64(v2)) {
			return true
		}
		return v1 == v2
	case ValTypeF64:
		v1 := v.MustGetF64()
		v2 := x.MustGetF64()
		if math.IsNaN(v1) && math.IsNaN(v2) {
			return true
		}
		return v1 == v2
	default:
		return false
	}
}

func (v *Val) String() string {
	switch Opcode((*v)[0]) {
	case OpcodeI32Const:
		x := v.MustGetI32()
		return fmt.Sprintf("i32:%#08x %d %d", x, uint32(x), int32(x))
	case OpcodeI64Const:
		x := v.MustGetI64()
		return fmt.Sprintf("i64:%#016x %d %d", x, uint64(x), int64(x))
	case OpcodeF32Const:
		x := v.MustGetF32()
		return fmt.Sprintf("f32:%#08x %f %e", math.Float32bits(x), x, x)
	case OpcodeF64Const:
		x := v.MustGetF64()
		return fmt.Sprintf("f64:%#016x %f %e", math.Float64bits(x), x, x)
	default:
		return "-"
	}
}

func DumpVals(vals []*Val) string {
	s := []string{}
	for i, v := range vals {
		s = append(s, fmt.Sprintf("%d:%s", i, v.String()))
	}
	return strings.Join(s, ", ")
}
