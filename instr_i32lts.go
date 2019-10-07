package wax

import (
	"context"
	"fmt"
)

type InstrI32Lts struct {
	Opcode Opcode
}

func ParseInstrI32Lts(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Lts, error) {
	return &InstrI32Lts{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Lts) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetI32() < v2.MustGetI32() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Lts) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.lt_s"),
	}, nil
}
