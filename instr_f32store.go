package wax

import (
	"context"
	"fmt"
)

type InstrF32Store struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func NewInstrF32Store(memArg MemArg, memArgBytes []byte) *InstrF32Store {
	return &InstrF32Store{
		opcode:      OpcodeF32Store,
		MemArg:      memArg,
		MemArgBytes: memArgBytes,
	}
}

func ParseInstrF32Store(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Store, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrF32Store{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrF32Store) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Store) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, store(rt, ValTypeF32, instr.MemArg)
}

func (instr *InstrF32Store) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("f32.store a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
