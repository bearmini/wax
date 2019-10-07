package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrGlobalGet struct {
	Opcode         Opcode
	GlobalIdx      GlobalIdx
	GlobalIdxBytes []byte
}

func ParseInstrGlobalGet(opcode Opcode, ber *BinaryEncodingReader) (*InstrGlobalGet, error) {
	x64, xBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	x := GlobalIdx(x64)

	return &InstrGlobalGet{
		Opcode:         opcode,
		GlobalIdx:      x,
		GlobalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrGlobalGet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrGlobalGet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.GlobalIdxBytes...),
		mnemonic: fmt.Sprintf("global.get %08x", instr.GlobalIdx),
	}, nil
}
