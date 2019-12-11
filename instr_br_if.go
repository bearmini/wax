package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrBrIf struct {
	opcode        Opcode
	LabelIdx      LabelIdx
	LabelIdxBytes []byte
}

func NewInstrBrIf(labelIdx LabelIdx, labelIdxBytes []byte) *InstrBrIf {
	return &InstrBrIf{
		opcode:        OpcodeBrIf,
		LabelIdx:      labelIdx,
		LabelIdxBytes: labelIdxBytes,
	}
}

func ParseInstrBrIf(opcode Opcode, ber *BinaryEncodingReader) (*InstrBrIf, error) {
	l64, lBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	l := LabelIdx(l64)

	return &InstrBrIf{
		opcode:        opcode,
		LabelIdx:      l,
		LabelIdxBytes: lBytes,
	}, nil
}

func (instr *InstrBrIf) Opcode() Opcode {
	return instr.opcode
}

/*
br_if l

1.  Assert: due to validation, a value of value type i32 is on the top of the stack.
2.  Pop the value i32.const c from the stack.
3.  If c is non-zero, then:
			(a) Execute the instruction (br l).
4.  Else:
      (a) Do nothing.
*/
func (instr *InstrBrIf) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1.  Assert: due to validation, a value of value type i32 is on the top of the stack.
	top := rt.Stack.Top()
	if top == nil || top.Value == nil {
		return nil, errors.New("i32 value must be on top of the stack")
	}
	t, err := top.Value.GetType()
	if err != nil {
		return nil, err
	}
	if *t != ValTypeI32 {
		return nil, errors.New("i32 value must be on top of the stack")
	}

	// 2.  Pop the value i32.const c from the stack.
	c, err := rt.Stack.PopValue()

	// 3.  If c is non-zero, then:
	if c.MustGetI32() != 0 {
		// (a) Execute the instruction (br l).
		ib := InstrBr{LabelIdx: instr.LabelIdx}
		return ib.Perform(ctx, rt)
	}

	// 4.  Else:
	//    (a) Do nothing.
	return nil, nil
}

func (instr *InstrBrIf) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.LabelIdxBytes...),
		mnemonic: fmt.Sprintf("br_if labelidx:%08x", instr.LabelIdx),
	}, nil
}
