package wax

import (
	"context"
	"fmt"
)

type InstrI64Load16s struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI64Load16s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Load16s, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Load16s{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Load16s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Load16s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, loadN(rt, ValTypeI64, 16, "s", instr.MemArg)
}

func (instr *InstrI64Load16s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.load16_s a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
