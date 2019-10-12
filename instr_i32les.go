package wax

import (
	"context"
	"fmt"
)

type InstrI32Les struct {
	opcode Opcode
}

func ParseInstrI32Les(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Les, error) {
	return &InstrI32Les{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Les) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Les) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if int32(v1.MustGetI32()) <= int32(v2.MustGetI32()) {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Les) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.le_s"),
	}, nil
}
