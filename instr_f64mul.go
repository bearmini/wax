package wax

import (
	"context"
	"fmt"
)

type InstrF64Mul struct {
	opcode Opcode
}

func NewInstrF64Mul() *InstrF64Mul {
	return &InstrF64Mul{
		opcode: OpcodeF64Mul,
	}
}

func ParseInstrF64Mul(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Mul, error) {
	return &InstrF64Mul{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Mul) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Mul) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		return NewValF64(v1.MustGetF64() * v2.MustGetF64()), nil
	})
}

func (instr *InstrF64Mul) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.mul"),
	}, nil
}
