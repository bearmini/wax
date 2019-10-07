package wax

import (
	"context"
	"fmt"
)

type InstrI32Store8 struct {
	Opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Store8(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Store8, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Store8{
		Opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Store8) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, storeN(rt, ValTypeI32, 8, instr.MemArg)
}

func (instr *InstrI32Store8) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.store8 a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
