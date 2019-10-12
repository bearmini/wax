package wax

import (
	"context"
	"fmt"
)

type InstrI64Load struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI64Load(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Load, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Load{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Load) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Load) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, load(rt, ValTypeI64, instr.MemArg)
}

func (instr *InstrI64Load) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.load a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
