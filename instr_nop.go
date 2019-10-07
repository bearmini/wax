package wax

import "context"

type InstrNop struct {
	Opcode Opcode
}

func ParseInstrNop(opcode Opcode, ber *BinaryEncodingReader) (*InstrNop, error) {
	return &InstrNop{Opcode: opcode}, nil
}

func (instr *InstrNop) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, nil
}

func (instr *InstrNop) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: "nop",
	}, nil
}
