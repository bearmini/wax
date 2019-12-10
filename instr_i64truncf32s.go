package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI64TruncF32s struct {
	opcode Opcode
}

func NewInstrI64TruncF32s() *InstrI64TruncF32s {
	return &InstrI64TruncF32s{
		opcode: OpcodeI64TruncF32s,
	}
}

func ParseInstrI64TruncF32s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64TruncF32s, error) {
	return &InstrI64TruncF32s{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64TruncF32s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64TruncF32s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF32, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(int64(math.Trunc(float64(v.MustGetF32()))))), nil
	})
}

func (instr *InstrI64TruncF32s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.trunc_f32_s"),
	}, nil
}
