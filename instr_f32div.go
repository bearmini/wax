package wax

import (
	"context"
	"fmt"
)

type InstrF32Div struct {
	opcode Opcode
}

func ParseInstrF32Div(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Div, error) {
	return &InstrF32Div{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Div) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Div) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF32()
		f2 := v2.MustGetF32()
		res := f1 / f2
		return NewValF32(res), nil
	})
}

func (instr *InstrF32Div) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.div"),
	}, nil
}
