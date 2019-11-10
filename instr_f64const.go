package wax

import (
	"context"
	"fmt"
)

type InstrF64Const struct {
	opcode Opcode
	N      uint64
	NBytes []byte
}

func ParseInstrF64Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Const, error) {
	n64, nBytes, err := ber.ReadVaruintN(64)
	if err != nil {
		return nil, err
	}

	return &InstrF64Const{
		opcode: opcode,
		N:      n64,
		NBytes: nBytes,
	}, nil
}

func (instr *InstrF64Const) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValF64(instr.N))
}

func (instr *InstrF64Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("f64.const %f (%016x)", instr.N, instr.N),
	}, nil
}
