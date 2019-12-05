package wax

import (
	"context"
	"fmt"
)

type InstrI64ExtenduI32 struct {
	opcode Opcode
}

func NewInstrI64ExtenduI32() *InstrI64ExtenduI32 {
	return &InstrI64ExtenduI32{
		opcode: OpcodeI64ExtenduI32,
	}
}

func ParseInstrI64ExtenduI32(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64ExtenduI32, error) {
	return &InstrI64ExtenduI32{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64ExtenduI32) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64ExtenduI32) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(v.MustGetI32())), nil
	})
}

func (instr *InstrI64ExtenduI32) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.extend_i32_u"),
	}, nil
}
