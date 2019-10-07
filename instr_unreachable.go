package wax

import (
	"context"

	"github.com/pkg/errors"
)

type InstrUnreachable struct {
	Opcode Opcode
}

func ParseInstrUnreachable(opcode Opcode, ber *BinaryEncodingReader) (*InstrUnreachable, error) {
	return &InstrUnreachable{Opcode: opcode}, nil
}

func (instr *InstrUnreachable) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("reached unreachable")
}

func (instr *InstrUnreachable) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: "unreachable",
	}, nil
}
