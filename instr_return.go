package wax

import (
	"context"

	"github.com/pkg/errors"
)

type InstrReturn struct {
	Opcode Opcode
}

func ParseInstrReturn(opcode Opcode, ber *BinaryEncodingReader) (*InstrReturn, error) {
	return &InstrReturn{Opcode: opcode}, nil
}

func (instr *InstrReturn) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrReturn) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: "return",
	}, nil
}
