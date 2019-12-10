package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32Ceil struct {
	opcode Opcode
}

func NewInstrF32Ceil() *InstrF32Ceil {
	return &InstrF32Ceil{
		opcode: OpcodeF32Ceil,
	}
}

func ParseInstrF32Ceil(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Ceil, error) {
	return &InstrF32Ceil{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Ceil) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Ceil) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(math.Ceil(float64(v1.MustGetF32())))), nil
	})
}

func (instr *InstrF32Ceil) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.ceil"),
	}, nil
}
