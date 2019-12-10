package wax

import (
	"context"
	"fmt"
)

type InstrF32Min struct {
	opcode Opcode
}

func ParseInstrF32Min(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Min, error) {
	return &InstrF32Min{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Min) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Min) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF32()
		f2 := v2.MustGetF32()
		if f1 <= f2 {
			return v1, nil
		}
		return v2, nil
	})
}

func (instr *InstrF32Min) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.min"),
	}, nil
}
