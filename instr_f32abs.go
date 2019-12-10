package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32Abs struct {
	opcode Opcode
}

func NewInstrF32Abs() *InstrF32Abs {
	return &InstrF32Abs{
		opcode: OpcodeF32Abs,
	}
}

func ParseInstrF32Abs(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Abs, error) {
	return &InstrF32Abs{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Abs) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Abs) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(math.Abs(float64(v1.MustGetF32())))), nil
	})
}

func (instr *InstrF32Abs) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.abs"),
	}, nil
}
