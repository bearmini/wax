package wax

import (
	"context"
	"fmt"
)

type InstrI64Load8s struct {
	opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI64Load8s(opcode Opcode, ber *BinaryEncodingReader) (*InstrI64Load8s, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI64Load8s{
		opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI64Load8s) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrI64Load8s) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, loadN(rt, ValTypeI64, 8, "s", instr.MemArg)
}

func (instr *InstrI64Load8s) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i64.load8_s a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
