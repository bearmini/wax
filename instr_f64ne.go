package wax

import (
	"context"
	"fmt"
)

type InstrF64Ne struct {
	opcode Opcode
}

func ParseInstrF64Ne(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Ne, error) {
	return &InstrF64Ne{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Ne) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Ne) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetF64() != v2.MustGetF64() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrF64Ne) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.ne"),
	}, nil
}
