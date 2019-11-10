package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrI32Rems struct {
	opcode Opcode
}

func ParseInstrI32Rems(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Rems, error) {
	return &InstrI32Rems{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Rems) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Rems) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		i1 := int32(v1.MustGetI32())
		i2 := int32(v2.MustGetI32())
		if i2 == 0 {
			return nil, errors.New("integer divide by zero")
		}
		return NewValI32(uint32(i1 % i2)), nil
	})
}

func (instr *InstrI32Rems) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.rem_s"),
	}, nil
}
