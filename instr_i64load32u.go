package wax

import (
	"context"
	"fmt"
)

type InstrI64Load32u struct {
	Opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI64Load32u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Load32u, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Load32u{
		Opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Load32u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, loadN(rt, ValTypeI64, 32, "u", instr.MemArg)
}

func (instr *InstrI64Load32u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.load32_u a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
