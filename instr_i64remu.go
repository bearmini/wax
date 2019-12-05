package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrI64Remu struct {
	opcode Opcode
}

func ParseInstrI64Remu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Remu, error) {
	return &InstrI64Remu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Remu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Remu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		i1 := uint64(v1.MustGetI64())
		i2 := uint64(v2.MustGetI64())
		if i2 == 0 {
			return nil, errors.New("integer divide by zero")
		}
		return NewValI64(uint64(i1 % i2)), nil
	})
}

func (instr *InstrI64Remu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.rem_u"),
	}, nil
}
