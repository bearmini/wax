package wax

import (
	"context"
	"fmt"
)

type InstrI32Load8u struct {
	Opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Load8u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Load8u, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Load8u{
		Opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Load8u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, loadN(rt, ValTypeI32, 8, "u", instr.MemArg)
}

func (instr *InstrI32Load8u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.load8_u a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
