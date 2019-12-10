package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF32ReinterpretI32 struct {
	opcode Opcode
}

func NewInstrF32ReinterpretI32() *InstrF32ReinterpretI32 {
	return &InstrF32ReinterpretI32{
		opcode: OpcodeF32ReinterpretI32,
	}
}

func ParseInstrF32ReinterpretI32(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32ReinterpretI32, error) {
	return &InstrF32ReinterpretI32{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32ReinterpretI32) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32ReinterpretI32) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeF32, func(v *Val) (*Val, error) {
		return NewValF32(math.Float32frombits(v.MustGetI32())), nil
	})
}

func (instr *InstrF32ReinterpretI32) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.reinterpret_i32"),
	}, nil
}
