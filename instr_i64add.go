package wax

import (
	"context"
	"fmt"
)

type InstrI64Add struct {
	opcode Opcode
}

func ParseInstrI64Add(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Add, error) {
	return &InstrI64Add{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Add) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Add) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() + v2.MustGetI64()), nil
	})
}

func (instr *InstrI64Add) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.add"),
	}, nil
}
