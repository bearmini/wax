package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI32ReinterpretF32 struct {
	opcode Opcode
}

func NewInstrI32ReinterpretF32() *InstrI32ReinterpretF32 {
	return &InstrI32ReinterpretF32{
		opcode: OpcodeI32ReinterpretF32,
	}
}

func ParseInstrI32ReinterpretF32(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32ReinterpretF32, error) {
	return &InstrI32ReinterpretF32{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32ReinterpretF32) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32ReinterpretF32) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF32, ValTypeI32, func(v *Val) (*Val, error) {
		return NewValI32(math.Float32bits(v.MustGetF32())), nil
	})
}

func (instr *InstrI32ReinterpretF32) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.reinterpret_f32"),
	}, nil
}
