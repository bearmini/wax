package wax

import (
	"context"
	"fmt"
)

type InstrI32Gtu struct {
	opcode Opcode
}

func ParseInstrI32Gtu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Gtu, error) {
	return &InstrI32Gtu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Gtu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Gtu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetI32() > v2.MustGetI32() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Gtu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.gt_u"),
	}, nil
}
