package wax

import (
	"context"
	"fmt"
	"math"

	"github.com/pkg/errors"
)

type InstrI64Divs struct {
	opcode Opcode
}

func ParseInstrI64Divs(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Divs, error) {
	return &InstrI64Divs{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Divs) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Divs) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeI64, func(v1, v2 *Val) (*Val, error) {
		i1 := int64(v1.MustGetI64())
		i2 := int64(v2.MustGetI64())
		if i2 == 0 {
			return nil, errors.New("integer divide by zero")
		}
		if i1 == math.MinInt64 && i2 == -1 {
			return nil, errors.New("integer overflow")
		}
		res := int64(i1) / int64(i2)
		return NewValI64(uint64(res)), nil
	})
}

func (instr *InstrI64Divs) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.div_s"),
	}, nil
}
