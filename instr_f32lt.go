package wax

import (
	"context"
	"fmt"
)

type InstrF32Lt struct {
	opcode Opcode
}

func ParseInstrF32Lt(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Lt, error) {
	return &InstrF32Lt{
		opcode: opcode,
	}, nil
}

func (instr *InstrF32Lt) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Lt) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeF32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetF32() < v2.MustGetF32() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrF32Lt) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f32.lt"),
	}, nil
}
