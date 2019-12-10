package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32Nearest struct {
	opcode Opcode
}

func NewInstrF32Nearest() *InstrF32Nearest {
	return &InstrF32Nearest{
		opcode: OpcodeF32Nearest,
	}
}

func ParseInstrF32Nearest(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Nearest, error) {
	return &InstrF32Nearest{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Nearest) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Nearest) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(math.Round(float64(v1.MustGetF32())))), nil
	})
}

func (instr *InstrF32Nearest) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.nearest"),
	}, nil
}
