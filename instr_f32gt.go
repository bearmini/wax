package wax

import (
	"context"
	"fmt"
)

type InstrF32Gt struct {
	opcode Opcode
}

func ParseInstrF32Gt(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Gt, error) {
	return &InstrF32Gt{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Gt) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Gt) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetF32() > v2.MustGetF32() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrF32Gt) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.gt"),
	}, nil
}
