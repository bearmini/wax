package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrLocalTee struct {
	opcode        Opcode
	LocalIdx      LocalIdx
	LocalIdxBytes []byte
}

func ParseInstrLocalTee(opcode Opcode, ber *BinaryEncodingReader) (*InstrLocalTee, error) {
	x64, xBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	x := LocalIdx(x64)

	return &InstrLocalTee{
		opcode:        opcode,
		LocalIdx:      x,
		LocalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrLocalTee) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrLocalTee) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	val, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}
	valCopy := *val

	// push the value immediately
	err = rt.Stack.PushValue(&valCopy)

	// do the same with local.set
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return nil, errors.New("no frame found")
	}

	if instr.LocalIdx >= LocalIdx(len(f.Locals)) {
		return nil, errors.New("out of range")
	}

	f.Locals[instr.LocalIdx] = &valCopy
	return nil, nil
}

func (instr *InstrLocalTee) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.LocalIdxBytes...),
		mnemonic: fmt.Sprintf("local.tee %08x", instr.LocalIdx),
	}, nil
}
