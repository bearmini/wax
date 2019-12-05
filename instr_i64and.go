package wax

import (
	"context"
	"fmt"
)

type InstrI64And struct {
	opcode Opcode
}

func ParseInstrI64And(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64And, error) {
	return &InstrI64And{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64And) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64And) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		return NewValI64(v1.MustGetI64() & v2.MustGetI64()), nil
	})
}

func (instr *InstrI64And) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.and"),
	}, nil
}
