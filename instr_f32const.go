package wax

import (
	"context"
	"fmt"
)

type InstrF32Const struct {
	opcode Opcode
	N      uint32
	NBytes []byte
}

func ParseInstrF32Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Const, error) {
	n64, nBytes, err := ber.ReadVarintN(32)
	if err != nil {
		return nil, err
	}

	return &InstrF32Const{
		opcode: opcode,
		N:      uint32(n64),
		NBytes: nBytes,
	}, nil
}

func (instr *InstrF32Const) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValF32(instr.N))
}

func (instr *InstrF32Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("f32.const %08x", instr.N),
	}, nil
}
