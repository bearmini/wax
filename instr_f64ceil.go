package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64Ceil struct {
	opcode Opcode
}

func NewInstrF64Ceil() *InstrF64Ceil {
	return &InstrF64Ceil{
		opcode: OpcodeF64Ceil,
	}
}

func ParseInstrF64Ceil(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Ceil, error) {
	return &InstrF64Ceil{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Ceil) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Ceil) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(math.Ceil(v1.MustGetF64())), nil
	})
}

func (instr *InstrF64Ceil) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.ceil"),
	}, nil
}
