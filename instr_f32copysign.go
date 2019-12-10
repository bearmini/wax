package wax

import (
	"context"
	"fmt"
)

type InstrF32CopySign struct {
	opcode Opcode
}

func ParseInstrF32CopySign(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32CopySign, error) {
	return &InstrF32CopySign{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32CopySign) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32CopySign) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF32()
		f2 := v2.MustGetF32()
		if (f1 >= 0 && f2 >= 0) || (f1 < 0 && f2 < 0) {
			return v1, nil
		}
		return NewValF32(-f1), nil
	})
}

func (instr *InstrF32CopySign) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.copysign"),
	}, nil
}
