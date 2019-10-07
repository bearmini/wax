package wax

import (
	"context"
	"fmt"
)

type InstrI32Leu struct {
	Opcode Opcode
}

func ParseInstrI32Leu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Leu, error) {
	return &InstrI32Leu{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Leu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetI32() <= v2.MustGetI32() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Leu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.le_u"),
	}, nil
}
