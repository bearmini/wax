package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrBr struct {
	opcode        Opcode
	LabelIdx      LabelIdx
	LabelIdxBytes []byte
}

func NewInstrBr(labelIdx LabelIdx, labelIdxBytes []byte) *InstrBr {
	return &InstrBr{
		opcode:        OpcodeBr,
		LabelIdx:      labelIdx,
		LabelIdxBytes: labelIdxBytes,
	}
}

func ParseInstrBr(opcode Opcode, ber *BinaryEncodingReader) (*InstrBr, error) {
	l64, lBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	l := LabelIdx(l64)

	return &InstrBr{
		opcode:        opcode,
		LabelIdx:      l,
		LabelIdxBytes: lBytes,
	}, nil
}

func (instr *InstrBr) Opcode() Opcode {
	return instr.opcode
}

/*
br l
1.  Assert: due to validation, the stack contains at least l + 1 labels.
2.  Let L be the l-th label appearing on the stack, starting from the top and counting from zero.
3.  Let n be the arity of L.
4.  Assert: due to validation, there are at least n values on the top of the stack.
5.  Pop the values val n from the stack.
6.  Repeat l + 1  times:
			(a)  While the top of the stack is a value, do:
						 i.  Pop the value from the stack.
			(b)  Assert: due to validation, the top of the stack now is a label.
			(c)  Pop the label from the stack.
7.  Push the values val n to the stack.
8.  Jump to the continuation of L.
*/
func (instr *InstrBr) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1.  Assert: due to validation, the stack contains at least l + 1 labels.
	if uint32(rt.Stack.CountLabels()) < uint32(instr.LabelIdx+1) {
		return nil, errors.New("not enough labels")
	}

	// 2.  Let L be the l-th label appearing on the stack, starting from the top and counting from zero.
	l, err := rt.Stack.GetLabelAt(instr.LabelIdx)
	if err != nil {
		return nil, err
	}

	// 3.  Let n be the arity of L.
	n := l.Arity

	// 4.  Assert: due to validation, there are at least n values on the top of the stack.
	if uint32(rt.Stack.CountValuesOnTop()) < n {
		return nil, errors.New("not enough values")
	}

	// 5.  Pop the values val n from the stack.
	vals, err := rt.Stack.PopValues(int(n))
	if err != nil {
		return nil, err
	}

	// 6.  Repeat l + 1  times:
	for i := uint32(0); i < uint32(instr.LabelIdx+1); i++ {
		// (a)  While the top of the stack is a value, do:
		for rt.Stack.Top().Value != nil {
			// i.  Pop the value from the stack.
			_, err := rt.Stack.PopValue()
			if err != nil {
				return nil, err
			}
		}

		// (b)  Assert: due to validation, the top of the stack now is a label.
		if rt.Stack.Top().Label == nil {
			return nil, errors.New("label expected")
		}

		// (c)  Pop the label from the stack.
		_, err := rt.Stack.PopLabel()
		if err != nil {
			return nil, err
		}
	}

	// 7.  Push the values val n to the stack.
	err = rt.Stack.PushValuesBack(vals)
	if err != nil {
		return nil, err
	}

	// 8.  Jump to the continuation of L.
	if instr.LabelIdx == 0 {
		return l, nil
	}

	ewj := NewEndWithJump(uint32(instr.LabelIdx + 1))
	return l, ewj
}

func (instr *InstrBr) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.LabelIdxBytes...),
		mnemonic: fmt.Sprintf("br labelidx:%08x", instr.LabelIdx),
	}, nil
}
