package wax

import (
	"context"
	"fmt"
)

type InstrI64ExtendI32s struct {
	opcode Opcode
}

func NewInstrI64ExtendI32s() *InstrI64ExtendI32s {
	return &InstrI64ExtendI32s{
		opcode: OpcodeI64ExtendI32s,
	}
}

func ParseInstrI64ExtendI32s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64ExtendI32s, error) {
	return &InstrI64ExtendI32s{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64ExtendI32s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64ExtendI32s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(uint64(int64(int32(v.MustGetI32())))), nil
	})
}

func (instr *InstrI64ExtendI32s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.extend_i32_s"),
	}, nil
}
