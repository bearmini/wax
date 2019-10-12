package wax

import (
	"context"
	"fmt"
)

type InstrI32Load struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Load(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Load, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Load{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Load) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Load) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, load(rt, ValTypeI32, instr.MemArg)
}

func (instr *InstrI32Load) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.load a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
