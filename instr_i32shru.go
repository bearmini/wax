package wax

import (
	"context"
	"fmt"
)

type InstrI32Shru struct {
	Opcode Opcode
}

func ParseInstrI32Shru(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Shru, error) {
	return &InstrI32Shru{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Shru) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(v1.MustGetI32() >> v2.MustGetI32()), nil
	})
}

func (instr *InstrI32Shru) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.shr_u"),
	}, nil
}
