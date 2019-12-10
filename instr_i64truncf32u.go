package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI64TruncF32u struct {
	opcode Opcode
}

func NewInstrI64TruncF32u() *InstrI64TruncF32u {
	return &InstrI64TruncF32u{
		opcode: OpcodeI64TruncF32u,
	}
}

func ParseInstrI64TruncF32u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64TruncF32u, error) {
	return &InstrI64TruncF32u{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64TruncF32u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64TruncF32u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF32, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(math.Trunc(float64(v.MustGetF32())))), nil
	})
}

func (instr *InstrI64TruncF32u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.trunc_f32_u"),
	}, nil
}
