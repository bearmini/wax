package wax

import (
	"context"
	"fmt"
)

type InstrF64ConvertI32u struct {
	opcode Opcode
}

func NewInstrF64ConvertI32u() *InstrF64ConvertI32u {
	return &InstrF64ConvertI32u{
		opcode: OpcodeF64ConvertI32u,
	}
}

func ParseInstrF64ConvertI32u(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64ConvertI32u, error) {
	return &InstrF64ConvertI32u{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64ConvertI32u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64ConvertI32u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(float64(v1.MustGetI32())), nil
	})
}

func (instr *InstrF64ConvertI32u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.convert_i32_u"),
	}, nil
}
