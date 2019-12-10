package wax

import (
	"context"
	"fmt"
)

type InstrI32WrapI64 struct {
	opcode Opcode
}

func NewInstrI32WrapI64() *InstrI32WrapI64 {
	return &InstrI32WrapI64{
		opcode: OpcodeI32WrapI64,
	}
}

func ParseInstrI32WrapI64(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32WrapI64, error) {
	return &InstrI32WrapI64{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32WrapI64) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32WrapI64) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI64, ValTypeI32, func(v *Val) (*Val, error) {
		return NewValI32(uint32(v.MustGetI64())), nil
	})
}

func (instr *InstrI32WrapI64) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.wrap_i64"),
	}, nil
}
