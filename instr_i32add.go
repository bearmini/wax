package wax

import (
	"context"
	"fmt"
)

type InstrI32Add struct {
	opcode Opcode
}

func ParseInstrI32Add(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Add, error) {
	return &InstrI32Add{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Add) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Add) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() + v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Add) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.add"),
	}, nil
}
