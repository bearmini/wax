package wax

import (
	"context"
	"fmt"
)

type InstrF64Add struct {
	opcode Opcode
}

func NewInstrF64Add() *InstrF64Add {
	return &InstrF64Add{
		opcode: OpcodeF64Add,
	}
}

func ParseInstrF64Add(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Add, error) {
	return &InstrF64Add{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Add) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Add) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		return NewValF64(v1.MustGetF64() + v2.MustGetF64()), nil
	})
}

func (instr *InstrF64Add) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.add"),
	}, nil
}
