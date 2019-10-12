package wax

import (
	"context"
	"fmt"
)

type InstrI32Ctz struct {
	opcode Opcode
}

func ParseInstrI32Ctz(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Ctz, error) {
	return &InstrI32Ctz{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Ctz) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Ctz) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeI32, func(v1 *Val) (*Val, error) {
		x := v1.MustGetI32()
		count := uint32(0)
		mask := uint32(0x00000001)
		for mask > 0 {
			if (x & mask) != 0 {
				return NewValI32(count), nil
			}

			count++
			mask <<= 1
		}

		return NewValI32(count), nil
	})
}

func (instr *InstrI32Ctz) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.ctz"),
	}, nil
}
