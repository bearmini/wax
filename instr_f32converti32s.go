package wax

import (
	"context"
	"fmt"
)

type InstrF32ConvertI32s struct {
	opcode Opcode
}

func NewInstrF32ConvertI32s() *InstrF32ConvertI32s {
	return &InstrF32ConvertI32s{
		opcode: OpcodeF32ConvertI32s,
	}
}

func ParseInstrF32ConvertI32s(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32ConvertI32s, error) {
	return &InstrF32ConvertI32s{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32ConvertI32s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32ConvertI32s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(int32(v1.MustGetI32()))), nil
	})
}

func (instr *InstrF32ConvertI32s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.convert_i32_s"),
	}, nil
}
