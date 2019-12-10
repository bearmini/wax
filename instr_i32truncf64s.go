package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI32TruncF64s struct {
	opcode Opcode
}

func NewInstrI32TruncF64s() *InstrI32TruncF64s {
	return &InstrI32TruncF64s{
		opcode: OpcodeI32TruncF64s,
	}
}

func ParseInstrI32TruncF64s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32TruncF64s, error) {
	return &InstrI32TruncF64s{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32TruncF64s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32TruncF64s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF64, ValTypeI32, func(v *Val) (*Val, error) {
		return NewValI32(uint32(int32(math.Trunc(v.MustGetF64())))), nil
	})
}

func (instr *InstrI32TruncF64s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.trunc_f64_s"),
	}, nil
}
