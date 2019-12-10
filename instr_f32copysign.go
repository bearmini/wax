package wax

import (
	"context"
	"fmt"
	"math"
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
		f1 := float64(v1.MustGetF32())
		f2 := float64(v2.MustGetF32())
		return NewValF32(float32(math.Copysign(f1, f2))), nil
	})
}

func (instr *InstrF32CopySign) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.copysign"),
	}, nil
}
