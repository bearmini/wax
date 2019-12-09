package wax

import (
	"context"
	"fmt"
)

type InstrF64Gt struct {
	opcode Opcode
}

func ParseInstrF64Gt(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Gt, error) {
	return &InstrF64Gt{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Gt) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Gt) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetF64() > v2.MustGetF64() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrF64Gt) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.gt"),
	}, nil
}
