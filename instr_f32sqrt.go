package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32Sqrt struct {
	opcode Opcode
}

func NewInstrF32Sqrt() *InstrF32Sqrt {
	return &InstrF32Sqrt{
		opcode: OpcodeF32Sqrt,
	}
}

func ParseInstrF32Sqrt(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Sqrt, error) {
	return &InstrF32Sqrt{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Sqrt) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Sqrt) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(math.Sqrt(float64(v1.MustGetF32())))), nil
	})
}

func (instr *InstrF32Sqrt) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.sqrt"),
	}, nil
}
