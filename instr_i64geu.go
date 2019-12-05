package wax

import (
	"context"
	"fmt"
)

type InstrI64Geu struct {
	opcode Opcode
}

func ParseInstrI64Geu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Geu, error) {
	return &InstrI64Geu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Geu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Geu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		result := uint32(0)
		if v1.MustGetI64() >= v2.MustGetI64() {
			result = uint32(1)
		}
		return NewValI32(result), nil
	})
}

func (instr *InstrI64Geu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.ge_u"),
	}, nil
}
