package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI32TruncF32u struct {
	opcode Opcode
}

func NewInstrI32TruncF32u() *InstrI32TruncF32u {
	return &InstrI32TruncF32u{
		opcode: OpcodeI32TruncF32u,
	}
}

func ParseInstrI32TruncF32u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32TruncF32u, error) {
	return &InstrI32TruncF32u{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32TruncF32u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32TruncF32u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF32, ValTypeI32, func(v *Val) (*Val, error) {
		return NewValI32(uint32(math.Trunc(float64(v.MustGetF32())))), nil
	})
}

func (instr *InstrI32TruncF32u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.trunc_f32_u"),
	}, nil
}
