package wax

import (
	"context"
	"fmt"
)

type InstrI64Gtu struct {
	opcode Opcode
}

func ParseInstrI64Gtu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Gtu, error) {
	return &InstrI64Gtu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Gtu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Gtu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		result := uint64(0)
		if v1.MustGetI64() > v2.MustGetI64() {
			result = uint64(1)
		}
		return NewValI64(result), nil
	})
}

func (instr *InstrI64Gtu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.gt_u"),
	}, nil
}
