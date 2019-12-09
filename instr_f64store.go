package wax

import (
	"context"
	"fmt"
)

type InstrF64Store struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func NewInstrF64Store(memArg MemArg, memArgBytes []byte) *InstrF64Store {
	return &InstrF64Store{
		opcode:      OpcodeF64Store,
		MemArg:      memArg,
		MemArgBytes: memArgBytes,
	}
}

func ParseInstrF64Store(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Store, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrF64Store{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrF64Store) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Store) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, store(rt, ValTypeF64, instr.MemArg)
}

func (instr *InstrF64Store) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("f64.store a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
