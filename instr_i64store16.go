package wax

import (
	"context"
	"fmt"
)

type InstrI64Store16 struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func NewInstrI64Store16(memArg MemArg, memArgBytes []byte) *InstrI64Store16 {
	return &InstrI64Store16{
		opcode:      OpcodeI64Store16,
		MemArg:      memArg,
		MemArgBytes: memArgBytes,
	}
}

func ParseInstrI64Store16(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Store16, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Store16{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Store16) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Store16) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, storeN(rt, ValTypeI64, 16, instr.MemArg)
}

func (instr *InstrI64Store16) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.store16 a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
