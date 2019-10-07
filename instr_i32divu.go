package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrI32Divu struct {
	Opcode Opcode
}

func ParseInstrI32Divu(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Divu, error) {
	return &InstrI32Divu{
		Opcode: opcode,
	}, nil
}

func (instr *InstrI32Divu) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		i1 := v1.MustGetI32()
		i2 := v2.MustGetI32()
		if i2 == 0 {
			return nil, errors.New("div by 0")
		}
		return NewValI32(i1 / i2), nil
	})
}

func (instr *InstrI32Divu) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("i32.div_u"),
	}, nil
}
