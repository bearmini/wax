package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64Abs struct {
	opcode Opcode
}

func NewInstrF64Abs() *InstrF64Abs {
	return &InstrF64Abs{
		opcode: OpcodeF64Abs,
	}
}

func ParseInstrF64Abs(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Abs, error) {
	return &InstrF64Abs{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Abs) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Abs) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(math.Abs(v1.MustGetF64())), nil
	})
}

func (instr *InstrF64Abs) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.abs"),
	}, nil
}
