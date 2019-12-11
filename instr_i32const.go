package wax

import (
	"context"
	"fmt"
)

type InstrI32Const struct {
	opcode Opcode
	N      uint32
	NBytes []byte
}

func NewInstrI32Const(n uint32, nBytes []byte) *InstrI32Const {
	return &InstrI32Const{
		opcode: OpcodeI32Const,
		N:      n,
		NBytes: nBytes,
	}
}

func ParseInstrI32Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Const, error) {
	n64, nBytes, err := ber.ReadVarint()
	if err != nil {
		return nil, err
	}

	return &InstrI32Const{
		opcode: opcode,
		N:      uint32(n64),
		NBytes: nBytes,
	}, nil
}

func (instr *InstrI32Const) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValI32(instr.N))
}

func (instr *InstrI32Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("i32.const %08x", instr.N),
	}, nil
}
