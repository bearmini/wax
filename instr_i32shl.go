package wax

import (
	"context"
	"fmt"
)

type InstrI32Shl struct {
	opcode Opcode
}

func ParseInstrI32Shl(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Shl, error) {
	return &InstrI32Shl{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Shl) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Shl) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		shift := (v2.MustGetI32() & 0x1f)
		return NewValI32(v1.MustGetI32() << shift), nil
	})
}

func (instr *InstrI32Shl) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.shl"),
	}, nil
}
