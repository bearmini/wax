package wax

import (
	"context"
	"fmt"
)

type InstrI64Ne struct {
	opcode Opcode
}

func ParseInstrI64Ne(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Ne, error) {
	return &InstrI64Ne{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Ne) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Ne) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		var result = uint32(0)
		if v1.MustGetI64() != v2.MustGetI64() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI64Ne) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.ne"),
	}, nil
}
