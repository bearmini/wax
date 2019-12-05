package wax

import (
	"context"
	"fmt"
)

type InstrI64Or struct {
	opcode Opcode
}

func ParseInstrI64Or(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Or, error) {
	return &InstrI64Or{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Or) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Or) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() | v2.MustGetI64()), nil
	})
}

func (instr *InstrI64Or) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.or"),
	}, nil
}
