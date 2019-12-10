package wax

import (
	"context"
	"fmt"
)

type InstrF64Max struct {
	opcode Opcode
}

func ParseInstrF64Max(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Max, error) {
	return &InstrF64Max{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Max) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Max) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF64()
		f2 := v2.MustGetF64()
		if f1 >= f2 {
			return v1, nil
		}
		return v2, nil
	})
}

func (instr *InstrF64Max) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.max"),
	}, nil
}
