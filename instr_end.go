package wax

import "context"

type InstrEnd struct {
	opcode Opcode
}

func NewInstrEnd() *InstrEnd {
	return &InstrEnd{
		opcode: OpcodeEnd,
	}
}

func ParseInstrEnd(opcode Opcode, ber *BinaryEncodingReader) (*InstrEnd, error) {
	return &InstrEnd{
		opcode: opcode,
	}, nil
}

func (instr *InstrEnd) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrEnd) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, nil
}

func (instr *InstrEnd) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: "end",
	}, nil
}
