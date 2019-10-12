package wax

import (
	"context"
	"fmt"
)

type InstrI64Mul struct {
	opcode Opcode
}

func ParseInstrI64Mul(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Mul, error) {
	return &InstrI64Mul{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Mul) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Mul) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() * v2.MustGetI64()), nil
	})
}

func (instr *InstrI64Mul) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.mul"),
	}, nil
}
