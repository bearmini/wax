package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrGlobalGet struct {
	opcode         Opcode
	GlobalIdx      GlobalIdx
	GlobalIdxBytes []byte
}

func ParseInstrGlobalGet(opcode Opcode, ber *BinaryEncodingReader) (*InstrGlobalGet, error) {
	x64, xBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	x := GlobalIdx(x64)

	return &InstrGlobalGet{
		opcode:         opcode,
		GlobalIdx:      x,
		GlobalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrGlobalGet) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrGlobalGet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrGlobalGet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.GlobalIdxBytes...),
		mnemonic: fmt.Sprintf("global.get %08x", instr.GlobalIdx),
	}, nil
}
