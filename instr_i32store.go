package wax

import (
	"context"
	"fmt"
)

type InstrI32Store struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Store(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Store, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Store{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Store) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Store) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, store(rt, ValTypeI32, instr.MemArg)
}

func (instr *InstrI32Store) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.store a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
