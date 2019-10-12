package wax

import (
	"context"

	"github.com/pkg/errors"
)

type InstrSelect struct {
	opcode Opcode
}

func ParseInstrSelect(opcode Opcode, ber *BinaryEncodingReader) (*InstrSelect, error) {
	return &InstrSelect{
		opcode: opcode,
	}, nil
}

func (instr *InstrSelect) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrSelect) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	v1, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	v2, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	v3, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	t3, err := v3.GetType()
	switch *t3 {
	case ValTypeI32:
		if v3.MustGetI32() == 0 {
			rt.Stack.PushValue(v1)
		} else {
			rt.Stack.PushValue(v2)
		}
	default:
		return nil, errors.New("not implemented")
	}

	return nil, nil
}

func (instr *InstrSelect) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}),
		mnemonic: "select",
	}, nil
}
