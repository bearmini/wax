package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrLocalSet struct {
	opcode        Opcode
	LocalIdx      LocalIdx
	LocalIdxBytes []byte
}

func NewInstrLocalSet(localIdx LocalIdx, localIdxBytes []byte) *InstrLocalGet {
	return &InstrLocalGet{
		opcode:        OpcodeLocalSet,
		LocalIdx:      localIdx,
		LocalIdxBytes: localIdxBytes,
	}
}

func ParseInstrLocalSet(opcode Opcode, ber *BinaryEncodingReader) (*InstrLocalSet, error) {
	x64, xBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	x := LocalIdx(x64)

	return &InstrLocalSet{
		opcode:        opcode,
		LocalIdx:      x,
		LocalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrLocalSet) Opcode() Opcode {
	return instr.opcode
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
		binary:   append([]byte{byte(instr.opcode)}, instr.LocalIdxBytes...),
		mnemonic: fmt.Sprintf("local.set %08x", instr.LocalIdx),
	}, nil
}
