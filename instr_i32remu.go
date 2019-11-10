package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrI32Remu struct {
	opcode Opcode
}

func ParseInstrI32Remu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Remu, error) {
	return &InstrI32Remu{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Remu) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Remu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		i1 := uint32(v1.MustGetI32())
		i2 := uint32(v2.MustGetI32())
		if i2 == 0 {
			return nil, errors.New("integer divide by zero")
		}
		return NewValI32(uint32(i1 % i2)), nil
	})
}

func (instr *InstrI32Remu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.rem_u"),
	}, nil
}
