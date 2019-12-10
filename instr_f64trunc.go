package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64Trunc struct {
	opcode Opcode
}

func NewInstrF64Trunc() *InstrF64Trunc {
	return &InstrF64Trunc{
		opcode: OpcodeF64Trunc,
	}
}

func ParseInstrF64Trunc(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Trunc, error) {
	return &InstrF64Trunc{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Trunc) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Trunc) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(math.Trunc(v1.MustGetF64())), nil
	})
}

func (instr *InstrF64Trunc) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.trunc"),
	}, nil
}
