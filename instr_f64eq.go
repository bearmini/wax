package wax

import (
	"context"
	"fmt"
)

type InstrF64Eq struct {
	opcode Opcode
}

func ParseInstrF64Eq(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Eq, error) {
	return &InstrF64Eq{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64Eq) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Eq) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetF64() == v2.MustGetF64() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrF64Eq) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.eq"),
	}, nil
}
