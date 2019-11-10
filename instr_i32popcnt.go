package wax

import (
	"context"
	"fmt"
)

type InstrI32Popcnt struct {
	opcode Opcode
}

func ParseInstrI32Popcnt(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Popcnt, error) {
	return &InstrI32Popcnt{
		opcode: opcode,
	}, nil
}

func (instr *InstrI32Popcnt) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Popcnt) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeI32, func(v1 *Val) (*Val, error) {
		x := v1.MustGetI32()
		count := uint32(0)
		mask := uint32(0x00000001)
		for mask != 0 {
			if (x & mask) != 0 {
				count++
			}
			mask <<= 1
		}

		return NewValI32(count), nil
	})
}

func (instr *InstrI32Popcnt) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i32.popcnt"),
	}, nil
}
