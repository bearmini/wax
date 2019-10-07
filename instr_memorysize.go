package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrMemorySize struct {
	Opcode Opcode
}

func ParseInstrMemorySize(opcode Opcode, ber *BinaryEncodingReader) (*InstrMemorySize, error) {
	x, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	if x != 0x00 {
		return nil, errors.Errorf("expected 0x00 but found %#2x", x)
	}

	return &InstrMemorySize{
		Opcode: opcode,
	}, nil
}

func (instr *InstrMemorySize) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrMemorySize) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, 0x00),
		mnemonic: fmt.Sprintf("memory.size 0x00"),
	}, nil
}
