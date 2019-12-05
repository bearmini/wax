package wax

import (
	"context"
	"fmt"
)

type InstrI64Shl struct {
	opcode Opcode
}

func ParseInstrI64Shl(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Shl, error) {
	return &InstrI64Shl{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Shl) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Shl) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		shift := (v2.MustGetI64() & 0x3f)
		return NewValI64(v1.MustGetI64() << shift), nil
	})
}

func (instr *InstrI64Shl) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.shl"),
	}, nil
}
