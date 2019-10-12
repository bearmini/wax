package wax

import (
	"context"
	"fmt"
)

type InstrI32Ltu struct {
	opcode Opcode
}

func ParseInstrI32Ltu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Ltu, error) {
	return &InstrI32Ltu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Ltu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Ltu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetI32() < v2.MustGetI32() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Ltu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.lt_u"),
	}, nil
}
