package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32Trunc struct {
	opcode Opcode
}

func NewInstrF32Trunc() *InstrF32Trunc {
	return &InstrF32Trunc{
		opcode: OpcodeF32Trunc,
	}
}

func ParseInstrF32Trunc(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Trunc, error) {
	return &InstrF32Trunc{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Trunc) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Trunc) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(math.Trunc(float64(v1.MustGetF32())))), nil
	})
}

func (instr *InstrF32Trunc) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.trunc"),
	}, nil
}
