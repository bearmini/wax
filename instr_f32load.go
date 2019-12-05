package wax

import (
	"context"
	"fmt"
)

type InstrF32Load struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrF32Load(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Load, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrF32Load{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrF32Load) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Load) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, load(rt, ValTypeF32, instr.MemArg)
}

func (instr *InstrF32Load) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("f32.load a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
