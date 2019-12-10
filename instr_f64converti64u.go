package wax

import (
	"context"
	"fmt"
)

type InstrF64ConvertI64u struct {
	opcode Opcode
}

func NewInstrF64ConvertI64u() *InstrF64ConvertI64u {
	return &InstrF64ConvertI64u{
		opcode: OpcodeF64ConvertI64u,
	}
}

func ParseInstrF64ConvertI64u(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64ConvertI64u, error) {
	return &InstrF64ConvertI64u{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64ConvertI64u) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64ConvertI64u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI64, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(float64(v1.MustGetI64())), nil
	})
}

func (instr *InstrF64ConvertI64u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.convert_i64_u"),
	}, nil
}
