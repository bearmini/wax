package wax

import (
	"context"
	"fmt"
)

type InstrI64ExtendI32u struct {
	opcode Opcode
}

func NewInstrI64ExtendI32u() *InstrI64ExtendI32u {
	return &InstrI64ExtendI32u{
		opcode: OpcodeI64ExtendI32u,
	}
}

func ParseInstrI64ExtendI32u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64ExtendI32u, error) {
	return &InstrI64ExtendI32u{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64ExtendI32u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64ExtendI32u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(v.MustGetI32())), nil
	})
}

func (instr *InstrI64ExtendI32u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.extend_i32_u"),
	}, nil
}
