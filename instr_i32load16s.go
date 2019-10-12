package wax

import (
	"context"
	"fmt"
)

type InstrI32Load16s struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Load16s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Load16s, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Load16s{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Load16s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI32Load16s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, loadN(rt, ValTypeI32, 16, "s", instr.MemArg)
}

func (instr *InstrI32Load16s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.load16_s a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
