package wax

import (
	"context"
	"fmt"
)

type InstrI32Shl struct {
	Opcode Opcode
}

func ParseInstrI32Shl(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Shl, error) {
	return &InstrI32Shl{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Shl) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() << v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Shl) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.shl"),
	}, nil
}
