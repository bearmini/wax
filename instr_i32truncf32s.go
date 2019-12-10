package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI32TruncF32s struct {
	opcode Opcode
}

func NewInstrI32TruncF32s() *InstrI32TruncF32s {
	return &InstrI32TruncF32s{
		opcode: OpcodeI32TruncF32s,
	}
}

func ParseInstrI32TruncF32s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32TruncF32s, error) {
	return &InstrI32TruncF32s{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32TruncF32s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32TruncF32s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF32, ValTypeI32, func(v *Val) (*Val, error) {
		return NewValI32(uint32(int32(math.Trunc(float64(v.MustGetF32()))))), nil
	})
}

func (instr *InstrI32TruncF32s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.trunc_f32_s"),
	}, nil
}
