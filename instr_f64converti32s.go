package wax

import (
	"context"
	"fmt"
)

type InstrF64ConvertI32s struct {
	opcode Opcode
}

func NewInstrF64ConvertI32s() *InstrF64ConvertI32s {
	return &InstrF64ConvertI32s{
		opcode: OpcodeF64ConvertI32s,
	}
}

func ParseInstrF64ConvertI32s(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64ConvertI32s, error) {
	return &InstrF64ConvertI32s{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64ConvertI32s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64ConvertI32s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI32, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(float64(int32(v1.MustGetI32()))), nil
	})
}

func (instr *InstrF64ConvertI32s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.convert_i32_s"),
	}, nil
}
