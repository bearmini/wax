package wax

import "context"

type InstrDrop struct {
	Opcode Opcode
}

func ParseInstrDrop(opcode Opcode, ber *BinaryEncodingReader) (*InstrDrop, error) {
	return &InstrDrop{
		Opcode: opcode,
	}, nil
}

func (instr *InstrDrop) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	_, err := rt.Stack.PopValue()
	return nil, err
}

func (instr *InstrDrop) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}),
		mnemonic: "drop",
	}, nil
}
