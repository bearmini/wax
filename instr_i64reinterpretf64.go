package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrI64ReinterpretF64 struct {
	opcode Opcode
}

func NewInstrI64ReinterpretF64() *InstrI64ReinterpretF64 {
	return &InstrI64ReinterpretF64{
		opcode: OpcodeI64ReinterpretF64,
	}
}

func ParseInstrI64ReinterpretF64(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64ReinterpretF64, error) {
	return &InstrI64ReinterpretF64{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64ReinterpretF64) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64ReinterpretF64) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, cvtop(rt, ValTypeF64, ValTypeI64, func(v *Val) (*Val, error) {
		return NewValI64(math.Float64bits(v.MustGetF64())), nil
	})
}

func (instr *InstrI64ReinterpretF64) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.reinterpret_f64"),
	}, nil
}
