package wax

import (
	"context"
	"fmt"
)

type InstrI32Gts struct {
	opcode Opcode
}

func NewInstrI32Gts() *InstrI32Gts {
	return &InstrI32Gts{
		opcode: OpcodeI32Gts,
	}
}

func ParseInstrI32Gts(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Gts, error) {
	return &InstrI32Gts{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Gts) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Gts) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if int32(v1.MustGetI32()) > int32(v2.MustGetI32()) {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Gts) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.gt_s"),
	}, nil
}
