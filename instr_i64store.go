package wax

import (
	"context"
	"fmt"
)

type InstrI64Store struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI64Store(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Store, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Store{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Store) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Store) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, store(rt, ValTypeI64, instr.MemArg)
}

func (instr *InstrI64Store) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.store a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
