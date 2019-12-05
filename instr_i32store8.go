package wax

import (
	"context"
	"fmt"
)

type InstrI32Store8 struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func NewInstrI32Store8(memArg MemArg, memArgBytes []byte) *InstrI32Store8 {
	return &InstrI32Store8{
		opcode:      OpcodeI32Store8,
		MemArg:      memArg,
		MemArgBytes: memArgBytes,
	}
}

func ParseInstrI32Store8(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Store8, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Store8{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Store8) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Store8) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, storeN(rt, ValTypeI32, 8, instr.MemArg)
}

func (instr *InstrI32Store8) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.store8 a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
