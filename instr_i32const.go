package wax

import (
	"context"
	"fmt"
)

type InstrI32Const struct {
	Opcode Opcode
	N      uint32
	NBytes []byte
}

func ParseInstrI32Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Const, error) {
	n64, nBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}

	return &InstrI32Const{
		Opcode: opcode,
		N:      uint32(n64),
		NBytes: nBytes,
	}, nil
}

func (instr *InstrI32Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValI32(instr.N))
}

func (instr *InstrI32Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("i32.const %08x", instr.N),
	}, nil
}
