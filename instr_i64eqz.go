package wax

import (
	"context"
	"fmt"
)

type InstrI64Eqz struct {
	opcode Opcode
}

func ParseInstrI64Eqz(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Eqz, error) {
	return &InstrI64Eqz{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Eqz) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Eqz) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, testop(rt, ValTypeI64, func(v *Val) (*Val, error) {
		var result = uint32(0)
		if v.MustGetI64() == uint64(0) {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI64Eqz) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.eqz"),
	}, nil
}
