package wax

import (
	"context"
	"fmt"
)

type InstrI64Rotr struct {
	opcode Opcode
}

func ParseInstrI64Rotr(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Rotr, error) {
	return &InstrI64Rotr{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Rotr) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Rotr) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		k := v2.MustGetI64() % 64
		i1 := v1.MustGetI64()
		return NewValI64((i1 >> k) | (i1 << (64 - k))), nil
	})
}

func (instr *InstrI64Rotr) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.rotr"),
	}, nil
}
