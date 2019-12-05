package wax

import (
	"context"
	"fmt"
)

type InstrI64Les struct {
	opcode Opcode
}

func ParseInstrI64Les(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Les, error) {
	return &InstrI64Les{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Les) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Les) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if int64(v1.MustGetI64()) <= int64(v2.MustGetI64()) {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI64Les) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.le_s"),
	}, nil
}
