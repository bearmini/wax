package wax

import "context"

type InstrEnd struct {
	Opcode Opcode
}

func ParseInstrEnd(opcode Opcode, ber *BinaryEncodingReader) (*InstrEnd, error) {
	return &InstrEnd{Opcode: opcode}, nil
}

func (instr *InstrEnd) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, nil
}

func (instr *InstrEnd) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: "end",
	}, nil
}
