package wax

import (
	"context"
	"fmt"
)

type InstrI64Ctz struct {
	opcode Opcode
}

func ParseInstrI64Ctz(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Ctz, error) {
	return &InstrI64Ctz{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Ctz) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Ctz) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeI64, func(v1 *Val) (*Val, error) {
		x := v1.MustGetI64()
		count := uint64(0)
		mask := uint64(0x0000000000000001)
		for mask > 0 {
			if (x & mask) != 0 {
				return NewValI64(count), nil
			}

			count++
			mask <<= 1
		}

		return NewValI64(count), nil
	})
}

func (instr *InstrI64Ctz) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.ctz"),
	}, nil
}
