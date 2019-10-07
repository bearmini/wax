package wax

import (
	"context"
	"fmt"
)

type InstrI32Store16 struct {
	Opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Store16(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Store16, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Store16{
		Opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Store16) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, storeN(rt, ValTypeI32, 16, instr.MemArg)
}

func (instr *InstrI32Store16) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.store16 a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
