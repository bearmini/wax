package wax

import (
	"context"
	"fmt"
)

type InstrI64Leu struct {
	opcode Opcode
}

func ParseInstrI64Leu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Leu, error) {
	return &InstrI64Leu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Leu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Leu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, relop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		result := uint64(0)
		if v1.MustGetI64() <= v2.MustGetI64() {
			result = uint64(1)
		}
		return NewValI64(result), nil
	})
}

func (instr *InstrI64Leu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.le_u"),
	}, nil
}
