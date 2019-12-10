package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64ReinterpretI64 struct {
	opcode Opcode
}

func NewInstrF64ReinterpretI64() *InstrF64ReinterpretI64 {
	return &InstrF64ReinterpretI64{
		opcode: OpcodeF64ReinterpretI64,
	}
}

func ParseInstrF64ReinterpretI64(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64ReinterpretI64, error) {
	return &InstrF64ReinterpretI64{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64ReinterpretI64) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64ReinterpretI64) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeI64, ValTypeF64, func(v *Val) (*Val, error) {
		return NewValF64(math.Float64frombits(v.MustGetI64())), nil
	})
}

func (instr *InstrF64ReinterpretI64) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.reinterpret_i64"),
	}, nil
}
