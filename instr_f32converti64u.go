package wax

import (
	"context"
	"fmt"
)

type InstrF32ConvertI64u struct {
	opcode Opcode
}

func NewInstrF32ConvertI64u() *InstrF32ConvertI64u {
	return &InstrF32ConvertI64u{
		opcode: OpcodeF32ConvertI64u,
	}
}

func ParseInstrF32ConvertI64u(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32ConvertI64u, error) {
	return &InstrF32ConvertI64u{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32ConvertI64u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32ConvertI64u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI64, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(v1.MustGetI64())), nil
	})
}

func (instr *InstrF32ConvertI64u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.convert_i64_u"),
	}, nil
}
