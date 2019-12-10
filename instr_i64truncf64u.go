package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI64TruncF64u struct {
	opcode Opcode
}

func NewInstrI64TruncF64u() *InstrI64TruncF64u {
	return &InstrI64TruncF64u{
		opcode: OpcodeI64TruncF64u,
	}
}

func ParseInstrI64TruncF64u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64TruncF64u, error) {
	return &InstrI64TruncF64u{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64TruncF64u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64TruncF64u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF64, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(math.Trunc(v.MustGetF64()))), nil
	})
}

func (instr *InstrI64TruncF64u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.trunc_f64_u"),
	}, nil
}
