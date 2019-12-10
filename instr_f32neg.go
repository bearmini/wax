package wax

import (
	"context"
	"fmt"
)

type InstrF32Neg struct {
	opcode Opcode
}

func NewInstrF32Neg() *InstrF32Neg {
	return &InstrF32Neg{
		opcode: OpcodeF32Neg,
	}
}

func ParseInstrF32Neg(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Neg, error) {
	return &InstrF32Neg{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Neg) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Neg) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(-v1.MustGetF32())), nil
	})
}

func (instr *InstrF32Neg) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.neg"),
	}, nil
}
