package wax

import (
	"context"
	"fmt"
)

type InstrI32Shrs struct {
	opcode Opcode
}

func ParseInstrI32Shrs(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Shrs, error) {
	return &InstrI32Shrs{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Shrs) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Shrs) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		return NewValI32(uint32(int32(v1.MustGetI32()) >> v2.MustGetI32())), nil
	})
}

func (instr *InstrI32Shrs) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.shr_s"),
	}, nil
}
