package wax

import (
	"context"
	"fmt"
)

type InstrF32ConvertI64s struct {
	opcode Opcode
}

func NewInstrF32ConvertI64s() *InstrF32ConvertI64s {
	return &InstrF32ConvertI64s{
		opcode: OpcodeF32ConvertI64s,
	}
}

func ParseInstrF32ConvertI64s(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32ConvertI64s, error) {
	return &InstrF32ConvertI64s{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32ConvertI64s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32ConvertI64s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI64, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(int64(v1.MustGetI64()))), nil
	})
}

func (instr *InstrF32ConvertI64s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.convert_i64_s"),
	}, nil
}
