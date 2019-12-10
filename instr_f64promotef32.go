package wax

import (
	"context"
	"fmt"
)

type InstrF64PromoteF32 struct {
	opcode Opcode
}

func NewInstrF64PromoteF32() *InstrF64PromoteF32 {
	return &InstrF64PromoteF32{
		opcode: OpcodeF64PromoteF32,
	}
}

func ParseInstrF64PromoteF32(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64PromoteF32, error) {
	return &InstrF64PromoteF32{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64PromoteF32) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64PromoteF32) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF32, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(float64(v1.MustGetF32())), nil
	})
}

func (instr *InstrF64PromoteF32) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.promote_f32"),
	}, nil
}
