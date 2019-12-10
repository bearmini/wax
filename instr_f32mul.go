package wax

import (
	"context"
	"fmt"
)

type InstrF32Mul struct {
	opcode Opcode
}

func NewInstrF32Mul() *InstrF32Mul {
	return &InstrF32Mul{
		opcode: OpcodeF32Mul,
	}
}

func ParseInstrF32Mul(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Mul, error) {
	return &InstrF32Mul{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Mul) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Mul) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		return NewValF32(v1.MustGetF32() * v2.MustGetF32()), nil
	})
}

func (instr *InstrF32Mul) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.mul"),
	}, nil
}
