package wax

import (
	"context"
	"fmt"
)

type InstrI64Shru struct {
	Opcode Opcode
}

func ParseInstrI64Shru(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Shru, error) {
	return &InstrI64Shru{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI64Shru) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() >> v2.MustGetI64()), nil
	})
}

func (instr *InstrI64Shru) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i64.shr_u"),
	}, nil
}
