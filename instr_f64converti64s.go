package wax

import (
	"context"
	"fmt"
)

type InstrF64ConvertI64s struct {
	opcode Opcode
}

func NewInstrF64ConvertI64s() *InstrF64ConvertI64s {
	return &InstrF64ConvertI64s{
		opcode: OpcodeF64ConvertI64s,
	}
}

func ParseInstrF64ConvertI64s(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64ConvertI64s, error) {
	return &InstrF64ConvertI64s{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64ConvertI64s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64ConvertI64s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI64, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(float64(int64(v1.MustGetI64()))), nil
	})
}

func (instr *InstrF64ConvertI64s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.convert_i64_s"),
	}, nil
}
