package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32Floor struct {
	opcode Opcode
}

func NewInstrF32Floor() *InstrF32Floor {
	return &InstrF32Floor{
		opcode: OpcodeF32Floor,
	}
}

func ParseInstrF32Floor(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Floor, error) {
	return &InstrF32Floor{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Floor) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Floor) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(math.Floor(float64(v1.MustGetF32())))), nil
	})
}

func (instr *InstrF32Floor) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.floor"),
	}, nil
}
