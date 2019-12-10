package wax

import (
	"context"
	"fmt"
)

type InstrF32Add struct {
	opcode Opcode
}

func NewInstrF32Add() *InstrF32Add {
	return &InstrF32Add{
		opcode: OpcodeF32Add,
	}
}

func ParseInstrF32Add(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Add, error) {
	return &InstrF32Add{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Add) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Add) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		return NewValF32(v1.MustGetF32() + v2.MustGetF32()), nil
	})
}

func (instr *InstrF32Add) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.add"),
	}, nil
}
