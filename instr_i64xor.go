package wax

import (
	"context"
	"fmt"
)

type InstrI64Xor struct {
	opcode Opcode
}

func ParseInstrI64Xor(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Xor, error) {
	return &InstrI64Xor{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Xor) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Xor) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() ^ v2.MustGetI64()), nil
	})
}

func (instr *InstrI64Xor) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.xor"),
	}, nil
}
