package wax

import (
	"context"
	"fmt"
)

type InstrI64Sub struct {
	opcode Opcode
}

func ParseInstrI64Sub(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Sub, error) {
	return &InstrI64Sub{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Sub) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Sub) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() - v2.MustGetI64()), nil
	})
}

func (instr *InstrI64Sub) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.sub"),
	}, nil
}
