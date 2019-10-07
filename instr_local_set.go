package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrLocalSet struct {
	Opcode        Opcode
	LocalIdx      LocalIdx
	LocalIdxBytes []byte
}

func ParseInstrLocalSet(opcode Opcode, ber *BinaryEncodingReader) (*InstrLocalSet, error) {
	x64, xBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	x := LocalIdx(x64)

	return &InstrLocalSet{
		Opcode:        opcode,
		LocalIdx:      x,
		LocalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrLocalSet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return nil, errors.New("no frame found")
	}

	if instr.LocalIdx >= LocalIdx(len(f.Locals)) {
		return nil, errors.New("out of range")
	}

	val, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	valCopy := *val
	f.Locals[instr.LocalIdx] = &valCopy

	return nil, nil
}

func (instr *InstrLocalSet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.Opcode)}, instr.LocalIdxBytes...),
		mnemonic: fmt.Sprintf("local.set %08x", instr.LocalIdx),
	}, nil
}
