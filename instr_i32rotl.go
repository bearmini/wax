package wax

import (
	"context"
	"fmt"
)

type InstrI32Rotl struct {
	opcode Opcode
}

func ParseInstrI32Rotl(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Rotl, error) {
	return &InstrI32Rotl{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Rotl) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Rotl) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		k := v2.MustGetI32() % 32
		i1 := v1.MustGetI32()
		return NewValI32((i1 << k) | (i1>>(32 - k))), nil
	})
}

func (instr *InstrI32Rotl) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.rotl"),
	}, nil
}
