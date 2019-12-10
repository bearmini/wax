package wax

import (
	"context"
	"fmt"
	"math"
)

type InstrF64CopySign struct {
	opcode Opcode
}

func ParseInstrF64CopySign(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64CopySign, error) {
	return &InstrF64CopySign{
		opcode: opcode,
	}, nil
}

func (instr *InstrF64CopySign) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64CopySign) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, binop(rt, ValTypeF64, func(v1, v2 *Val) (*Val, error) {
		f1 := v1.MustGetF64()
		f2 := v2.MustGetF64()
		return NewValF64(math.Copysign(f1, f2)), nil
	})
}

func (instr *InstrF64CopySign) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("f64.copysign"),
	}, nil
}
