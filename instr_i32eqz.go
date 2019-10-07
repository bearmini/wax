package wax

import (
	"context"
	"fmt"
)

type InstrI32Eqz struct {
	Opcode Opcode
}

func ParseInstrI32Eqz(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Eqz, error) {
	return &InstrI32Eqz{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Eqz) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, testop(rt, ValTypeI32, func(v *Val) (*Val, error) {
		var result = uint32(0)
		if v.MustGetI32() == uint32(0) {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI32Eqz) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.eqz"),
	}, nil
}
