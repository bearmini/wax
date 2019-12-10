package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64Sqrt struct {
	opcode Opcode
}

func NewInstrF64Sqrt() *InstrF64Sqrt {
	return &InstrF64Sqrt{
		opcode: OpcodeF64Sqrt,
	}
}

func ParseInstrF64Sqrt(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Sqrt, error) {
	return &InstrF64Sqrt{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Sqrt) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Sqrt) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(math.Sqrt(v1.MustGetF64())), nil
	})
}

func (instr *InstrF64Sqrt) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.sqrt"),
	}, nil
}
