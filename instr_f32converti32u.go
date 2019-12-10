package wax

import (
	"context"
	"fmt"
)

type InstrF32ConvertI32u struct {
	opcode Opcode
}

func NewInstrF32ConvertI32u() *InstrF32ConvertI32u {
	return &InstrF32ConvertI32u{
		opcode: OpcodeF32ConvertI32u,
	}
}

func ParseInstrF32ConvertI32u(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32ConvertI32u, error) {
	return &InstrF32ConvertI32u{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32ConvertI32u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32ConvertI32u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(v1.MustGetI32())), nil
	})
}

func (instr *InstrF32ConvertI32u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.convert_i32_u"),
	}, nil
}
