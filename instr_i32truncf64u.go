package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI32TruncF64u struct {
	opcode Opcode
}

func NewInstrI32TruncF64u() *InstrI32TruncF64u {
	return &InstrI32TruncF64u{
		opcode: OpcodeI32TruncF64u,
	}
}

func ParseInstrI32TruncF64u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32TruncF64u, error) {
	return &InstrI32TruncF64u{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32TruncF64u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32TruncF64u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF64, ValTypeI32, func(v *Val) (*Val, error) {
		return NewValI32(uint32(math.Trunc(v.MustGetF64()))), nil
	})
}

func (instr *InstrI32TruncF64u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.trunc_f64_u"),
	}, nil
}
