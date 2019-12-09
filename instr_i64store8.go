package wax

import (
	"context"
	"fmt"
)

type InstrI64Store8 struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func NewInstrI64Store8(memArg MemArg, memArgBytes []byte) *InstrI64Store8 {
	return &InstrI64Store8{
		opcode:      OpcodeI64Store8,
		MemArg:      memArg,
		MemArgBytes: memArgBytes,
	}
}

func ParseInstrI64Store8(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Store8, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Store8{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Store8) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Store8) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, storeN(rt, ValTypeI64, 8, instr.MemArg)
}

func (instr *InstrI64Store8) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.store8 a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
