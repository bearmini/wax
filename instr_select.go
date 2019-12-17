package wax

import (
	"context"

	"github.com/pkg/errors"
)

type InstrSelect struct {
	opcode Opcode
}

func NewInstrSelect() *InstrSelect {
	return &InstrSelect{
		opcode: OpcodeSelect,
	}
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
	// 1. Assert: due to validation, a value of value type i32 is on the top of the stack.
	err := rt.Stack.AssertTopIsValueI32()
	if err != nil {
		return nil, err
	}

	// 2. Pop the value i32.const 𝑐 from the stack.
	c, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	// 3. Assert: due to validation, two more values (of the same value type) are on the top of the stack.
	vals, err := rt.Stack.TopValues(2)
	if err != nil {
		return nil, err
	}
	t1, err := vals[0].GetType()
	if err != nil {
		return nil, err
	}
	t2, err := vals[1].GetType()
	if err != nil {
		return nil, err
	}
	if *t1 != *t2 {
		return nil, errors.New("types of two values on top of stack do not match")
	}

	// 4. Pop the value val 2 from the stack.
	v2, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	// 5. Pop the value val 1 from the stack.
	v1, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	// 6. If c is not 0, then:
	if c.MustGetI32() != 0 {
		// (a) Push the value val 1 back to the stack.
		err = rt.Stack.PushValue(v1)
	} else {
		// 7. Else:
		// (a) Push the value val 2 back to the stack.
		err = rt.Stack.PushValue(v2)
	}

	return nil, err
}

func (instr *InstrSelect) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}),
		mnemonic: "select",
	}, nil
}
