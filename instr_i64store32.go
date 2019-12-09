package wax

import (
	"context"
	"fmt"
)

type InstrI64Store32 struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func NewInstrI64Store32(memArg MemArg, memArgBytes []byte) *InstrI64Store32 {
	return &InstrI64Store32{
		opcode:      OpcodeI64Store32,
		MemArg:      memArg,
		MemArgBytes: memArgBytes,
	}
}

func ParseInstrI64Store32(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Store32, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Store32{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Store32) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Store32) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, storeN(rt, ValTypeI64, 32, instr.MemArg)
}

func (instr *InstrI64Store32) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.store32 a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
