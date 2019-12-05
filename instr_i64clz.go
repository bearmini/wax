package wax

import (
	"context"
	"fmt"
)

type InstrI64Clz struct {
	opcode Opcode
}

func ParseInstrI64Clz(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Clz, error) {
	return &InstrI64Clz{
		opcode: opcode,
	}, nil
}

func (instr *InstrI64Clz) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Clz) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, unop(rt, ValTypeI64, func(v1 *Val) (*Val, error) {
		x := v1.MustGetI64()
		count := uint64(0)
		mask := uint64(0x8000000000000000)
		for mask > 0 {
			if (x & mask) != 0 {
				return NewValI64(count), nil
			}

			count++
			mask >>= 1
		}

		return NewValI64(count), nil
	})
}

func (instr *InstrI64Clz) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("i64.clz"),
	}, nil
}
