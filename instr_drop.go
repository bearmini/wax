package wax

import "context"

type InstrDrop struct {
	opcode Opcode
}

func ParseInstrDrop(opcode Opcode, ber *BinaryEncodingReader) (*InstrDrop, error) {
	return &InstrDrop{
		opcode: opcode,
	}, nil
}

func (instr *InstrDrop) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrDrop) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	_, err := rt.Stack.PopValue()
	return nil, err
}

func (instr *InstrDrop) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}),
		mnemonic: "drop",
	}, nil
}
