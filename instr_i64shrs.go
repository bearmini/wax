package wax

import (
	"context"
	"fmt"
)

type InstrI64Shrs struct {
	opcode Opcode
}

func ParseInstrI64Shrs(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Shrs, error) {
	return &InstrI64Shrs{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Shrs) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Shrs) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		shift := (v2.MustGetI64() & 0x3f)
		return NewValI64(uint64(int64(v1.MustGetI64()) >> shift)), nil
	})
}

func (instr *InstrI64Shrs) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.shr_s"),
	}, nil
}
