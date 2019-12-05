package wax

import (
	"context"
	"fmt"
)

type InstrI64Shru struct {
	opcode Opcode
}

func NewInstrI64Shru() *InstrI64Shru {
	return &InstrI64Shru{
		opcode: OpcodeI64Shru,
	}
}

func ParseInstrI64Shru(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Shru, error) {
	return &InstrI64Shru{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Shru) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Shru) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		shift := (v2.MustGetI64() & 0x3f)
		return NewValI64(v1.MustGetI64() >> shift), nil
	})
}

func (instr *InstrI64Shru) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.shr_u"),
	}, nil
}
