package wax

import (
	"context"
	"fmt"
)

type InstrI32Or struct {
	opcode Opcode
}

func ParseInstrI32Or(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Or, error) {
	return &InstrI32Or{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Or) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Or) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() | v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Or) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.or"),
	}, nil
}
