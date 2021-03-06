package wax

import (
	"context"
	"fmt"
)

type InstrI32Xor struct {
	opcode Opcode
}

func ParseInstrI32Xor(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Xor, error) {
	return &InstrI32Xor{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Xor) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Xor) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() ^ v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Xor) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.xor"),
	}, nil
}
