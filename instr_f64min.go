package wax

import (
	"context"
	"fmt"
)

type InstrF64Min struct {
	opcode Opcode
}

func ParseInstrF64Min(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Min, error) {
	return &InstrF64Min{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Min) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Min) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF64()
		f2 := v2.MustGetF64()
		if f1 <= f2 {
			return v1, nil
		}
		return v2, nil
	})
}

func (instr *InstrF64Min) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.min"),
	}, nil
}
