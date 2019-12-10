package wax

import (
	"context"
	"fmt"
)

type InstrF64Sub struct {
	opcode Opcode
}

func ParseInstrF64Sub(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Sub, error) {
	return &InstrF64Sub{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Sub) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Sub) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		return NewValF64(v1.MustGetF64() - v2.MustGetF64()), nil
	})
}

func (instr *InstrF64Sub) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.sub"),
	}, nil
}
