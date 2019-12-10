package wax

import (
	"context"
	"fmt"
)

type InstrF64Neg struct {
	opcode Opcode
}

func NewInstrF64Neg() *InstrF64Neg {
	return &InstrF64Neg{
		opcode: OpcodeF64Neg,
	}
}

func ParseInstrF64Neg(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Neg, error) {
	return &InstrF64Neg{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Neg) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Neg) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(-v1.MustGetF64()), nil
	})
}

func (instr *InstrF64Neg) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.neg"),
	}, nil
}
