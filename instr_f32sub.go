package wax

import (
	"context"
	"fmt"
)

type InstrF32Sub struct {
	opcode Opcode
}

func ParseInstrF32Sub(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Sub, error) {
	return &InstrF32Sub{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Sub) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Sub) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		return NewValF32(v1.MustGetF32() - v2.MustGetF32()), nil
	})
}

func (instr *InstrF32Sub) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.sub"),
	}, nil
}
