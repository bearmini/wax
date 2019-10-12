package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrI32Divs struct {
	opcode Opcode
}

func ParseInstrI32Divs(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Divs, error) {
	return &InstrI32Divs{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Divs) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Divs) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI32, func(v1, v2 *Val) (*Val, error) {
		i1 := int32(v1.MustGetI32())
		i2 := int32(v2.MustGetI32())
		if i2 == 0 {
			return nil, errors.New("div by 0")
		}
		return NewValI32(uint32(i1 / i2)), nil
	})
}

func (instr *InstrI32Divs) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.div_s"),
	}, nil
}
