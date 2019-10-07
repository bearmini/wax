package wax

import (
	"context"
	"fmt"
)

type InstrI32Load16u struct {
	Opcode      Opcode
	MemArg      MemArg
	MemArgBytes []byte
}

func ParseInstrI32Load16u(opcode Opcode, ber *BinaryEncodingReader) (*InstrI32Load16u, error) {
	ma, maBytes, err := ParseMemArg(ber)
	if err != nil {
		return nil, err
	}

	return &InstrI32Load16u{
		Opcode:      opcode,
		MemArg:      *ma,
		MemArgBytes: maBytes,
	}, nil
}

func (instr *InstrI32Load16u) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, loadN(rt, ValTypeI32, 16, "u", instr.MemArg)
}

func (instr *InstrI32Load16u) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.MemArgBytes...),
		mnemonic: fmt.Sprintf("i32.load16_u a:%08x o:%08x", instr.MemArg.Align, instr.MemArg.Offset),
	}, nil
}
