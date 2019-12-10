package wax

import (
	"context"
	"fmt"
)

type InstrF64Div struct {
	opcode Opcode
}

func ParseInstrF64Div(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Div, error) {
	return &InstrF64Div{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Div) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Div) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF64()
		f2 := v2.MustGetF64()
		res := f1 / f2
		return NewValF64(res), nil
	})
}

func (instr *InstrF64Div) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.div"),
	}, nil
}
