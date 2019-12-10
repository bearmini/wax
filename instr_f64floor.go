package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64Floor struct {
	opcode Opcode
}

func NewInstrF64Floor() *InstrF64Floor {
	return &InstrF64Floor{
		opcode: OpcodeF64Floor,
	}
}

func ParseInstrF64Floor(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Floor, error) {
	return &InstrF64Floor{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Floor) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Floor) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeF64, func(v1 *Val) (*Val, error) {
		return NewValF64(math.Floor(v1.MustGetF64())), nil
	})
}

func (instr *InstrF64Floor) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.floor"),
	}, nil
}
