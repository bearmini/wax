package wax

import (
	"context"
	"fmt"
)

type InstrF32DemoteF64 struct {
	opcode Opcode
}

func NewInstrF32DemoteF64() *InstrF32DemoteF64 {
	return &InstrF32DemoteF64{
		opcode: OpcodeF32DemoteF64,
	}
}

func ParseInstrF32DemoteF64(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32DemoteF64, error) {
	return &InstrF32DemoteF64{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32DemoteF64) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32DemoteF64) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF64, ValTypeF32, func(v1 *Val) (*Val, error) {
		return NewValF32(float32(v1.MustGetF64())), nil
	})
}

func (instr *InstrF32DemoteF64) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.demote_f64"),
	}, nil
}
