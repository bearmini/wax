package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI64TruncF64s struct {
	opcode Opcode
}

func NewInstrI64TruncF64s() *InstrI64TruncF64s {
	return &InstrI64TruncF64s{
		opcode: OpcodeI64TruncF64s,
	}
}

func ParseInstrI64TruncF64s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64TruncF64s, error) {
	return &InstrI64TruncF64s{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64TruncF64s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64TruncF64s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF64, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(int64(math.Trunc(v.MustGetF64())))), nil
	})
}

func (instr *InstrI64TruncF64s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.trunc_f64_s"),
	}, nil
}
