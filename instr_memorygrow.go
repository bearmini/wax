package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrMemoryGrow struct {
	Opcode Opcode
}

func ParseInstrMemoryGrow(opcode Opcode, ber *BinaryEncodingReader) (*InstrMemoryGrow, error) {
	x, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	if x != 0x00 {
		return nil, errors.Errorf("expected 0x00 but found %#2x", x)
	}

	return &InstrMemoryGrow{
		Opcode: opcode,
	}, nil
}

func (instr *InstrMemoryGrow) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrMemoryGrow) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, 0x00),
		mnemonic: fmt.Sprintf("memory.grow 0x00"),
	}, nil
}
