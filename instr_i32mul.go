package wax

import (
	"context"
	"fmt"
)

type InstrI32Mul struct {
	opcode Opcode
}

func NewInstrI32Mul() *InstrI32Mul {
	return &InstrI32Mul{
		opcode: OpcodeI32Mul,
	}
}

func ParseInstrI32Mul(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Mul, error) {
	return &InstrI32Mul{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Mul) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Mul) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() * v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Mul) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.mul"),
	}, nil
}
