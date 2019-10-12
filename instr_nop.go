package wax

import "context"

type InstrNop struct {
	opcode Opcode
}

func ParseInstrNop(opcode Opcode, ber *BinaryEncodingReader) (*InstrNop, error) {
	return &InstrNop{opcode: opcode}, nil
}

func (instr *InstrNop) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrNop) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, nil
}

func (instr *InstrNop) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: "nop",
	}, nil
}
