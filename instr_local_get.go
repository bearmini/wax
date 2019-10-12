package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrLocalGet struct {
	opcode        Opcode
	LocalIdx      LocalIdx
	LocalIdxBytes []byte
}

func ParseInstrLocalGet(opcode Opcode, ber *BinaryEncodingReader) (*InstrLocalGet, error) {
	x64, xBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	x := LocalIdx(x64)

	return &InstrLocalGet{
		opcode:        opcode,
		LocalIdx:      x,
		LocalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrLocalGet) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrLocalGet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return nil, errors.New("no frame found")
	}

	if instr.LocalIdx >= LocalIdx(len(f.Locals)) {
		return nil, errors.New("out of range")
	}

	val := f.Locals[instr.LocalIdx]
	valCopy := *val
	return nil, rt.Stack.PushValue(&valCopy)
}

func (instr *InstrLocalGet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.LocalIdxBytes...),
		mnemonic: fmt.Sprintf("local.get %08x", instr.LocalIdx),
	}, nil
}
