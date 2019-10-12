package wax

import (
	"context"
	"fmt"
)

type InstrI32And struct {
	opcode Opcode
}

func ParseInstrI32And(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32And, error) {
	return &InstrI32And{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32And) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32And) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() & v2.MustGetI32()), nil
	})
}

func (instr *InstrI32And) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.and"),
	}, nil
}
