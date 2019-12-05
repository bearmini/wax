package wax

import (
	"context"
	"fmt"
)

type InstrF64Load struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrF64Load(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Load, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrF64Load{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrF64Load) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Load) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, load(rt, ValTypeF64, instr.MemArg)
}

func (instr *InstrF64Load) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("f64.load a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
