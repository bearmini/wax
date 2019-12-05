package wax

import (
	"context"
	"fmt"
)

type InstrI64Lts struct {
	opcode Opcode
}

func ParseInstrI64Lts(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Lts, error) {
	return &InstrI64Lts{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Lts) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Lts) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if int64(v1.MustGetI64()) < int64(v2.MustGetI64()) {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI64Lts) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.lt_s"),
	}, nil
}
