package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrI64Divu struct {
	opcode Opcode
}

func ParseInstrI64Divu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Divu, error) {
	return &InstrI64Divu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Divu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Divu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		i1 := v1.MustGetI64()
		i2 := v2.MustGetI64()
		if i2 == 0 {
			return nil, errors.New("integer divide by zero")
		}
		return NewValI64(i1 / i2), nil
	})
}

func (instr *InstrI64Divu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.div_u"),
	}, nil
}
