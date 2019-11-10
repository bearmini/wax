package wax

import (
	"context"
	"fmt"
)

type InstrI32Rotr struct {
	opcode Opcode
}

func ParseInstrI32Rotr(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Rotr, error) {
	return &InstrI32Rotr{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Rotr) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Rotr) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		k := v2.MustGetI32() % 32
		i1 := v1.MustGetI32()
		return NewValI32((i1 >> k) | (i1 << (32 - k))), nil
	})
}

func (instr *InstrI32Rotr) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.rotr"),
	}, nil
}
