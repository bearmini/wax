package wax

import (
	"context"
	"fmt"
)

type InstrI32Sub struct {
	Opcode Opcode
}

func ParseInstrI32Sub(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Sub, error) {
	return &InstrI32Sub{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Sub) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() - v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Sub) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.sub"),
	}, nil
}
