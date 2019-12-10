package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64Nearest struct {
	opcode Opcode
}

func NewInstrF64Nearest() *InstrF64Nearest {
	return &InstrF64Nearest{
		opcode: OpcodeF64Nearest,
	}
}

func ParseInstrF64Nearest(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Nearest, error) {
	return &InstrF64Nearest{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Nearest) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Nearest) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(math.RoundToEven(v1.MustGetF64())), nil
	})
}

func (instr *InstrF64Nearest) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.nearest"),
	}, nil
}
