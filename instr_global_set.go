package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrGlobalSet struct {
	Opcode         Opcode
	GlobalIdx      GlobalIdx
	GlobalIdxBytes []byte
}

func ParseInstrGlobalSet(opcode Opcode, ber *BinaryEncodingReader) (*InstrGlobalSet, error) {
	x64, xBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	x := GlobalIdx(x64)

	return &InstrGlobalSet{
		Opcode:         opcode,
		GlobalIdx:      x,
		GlobalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrGlobalSet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrGlobalSet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.GlobalIdxBytes...),
		mnemonic: fmt.Sprintf("global.set %08x", instr.GlobalIdx),
	}, nil
}
