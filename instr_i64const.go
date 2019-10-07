package wax

import (
	"context"
	"fmt"
)

type InstrI64Const struct {
	Opcode Opcode
	N      uint64
	NBytes []byte
}

func ParseInstrI64Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Const, error) {
	n64, nBytes, err := ber.ReadVaruintN(64)
	if err != nil {
		return nil, err
	}

	return &InstrI64Const{
		Opcode: opcode,
		N:      n64,
		NBytes: nBytes,
	}, nil
}

func (instr *InstrI64Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValI64(instr.N))
}

func (instr *InstrI64Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("i64.const %016x", instr.N),
	}, nil
}
