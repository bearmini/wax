package wax

import (
	"context"
	"fmt"
)

type InstrI32Or struct {
	Opcode Opcode
}

func ParseInstrI32Or(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Or, error) {
	return &InstrI32Or{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Or) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() | v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Or) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.or"),
	}, nil
}
