package wax

import (
	"context"

	"github.com/pkg/errors"
)

type InstrReturn struct {
	opcode Opcode
}

func ParseInstrReturn(opcode Opcode, ber *BinaryEncodingReader) (*InstrReturn, error) {
	return &InstrReturn{opcode: opcode}, nil
}

func (instr *InstrReturn) Opcode() Opcode {
	return instr.opcode
}

/*
1.  Let F be the current frame.
2.  Let n be the arity of F.
3.  Assert: due to validation, there are at least n values on the top of the stack.
4.  Pop the results val^n from the stack.
5.  Assert: due to validation, the stack contains at least one frame.
6.  While the top of the stack is not a frame, do:
  (a)  Pop the top element from the stack.
7.  Assert: the top of the stack is the frame F.
8.  Pop the frame from the stack.
9.  Push val^n to the stack.
10.  Jump to the instruction after the original call that pushed the frame
*/
func (instr *InstrReturn) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1.  Let F be the current frame.
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return nil, errors.New("frame not found")
	}

	// 2.  Let n be the arity of F.
	n := f.Arity

	// 3.  Assert: due to validation, there are at least n values on the top of the stack.
	if uint32(rt.Stack.CountValuesOnTop()) < n {
		return nil, errors.New("insufficient number of values on the top of the stack")
	}

	// 4.  Pop the results val^n from the stack.
	vals, err := rt.Stack.PopValues(int(n))
	if err != nil {
		return nil, err
	}

	// 5.  Assert: due to validation, the stack contains at least one frame.
	if rt.Stack.CountFrames() < 1 {
		return nil, errors.New("insufficient number of frames in the stack")
	}

	// 6.  While the top of the stack is not a frame, do:
	for rt.Stack.Top().Frame == nil {
		// (a)  Pop the top element from the stack.
		_, err = rt.Stack.Pop()
		if err != nil {
			return nil, err
		}
	}

	// 7.  Assert: the top of the stack is the frame F.
	if rt.Stack.Top().Frame != f {
		return nil, errors.New("the top of the stack must be the current frame")
	}

	// 8.  Pop the frame from the stack.
	_, err = rt.Stack.Pop()
	if err != nil {
		return nil, err
	}

	// 9.  Push val^n to the stack.
	err = rt.Stack.PushValuesBack(vals)
	if err != nil {
		return nil, err
	}

	// 10.  Jump to the instruction after the original call that pushed the frame
	return nil, NewEndWithReturn() 
}

func (instr *InstrReturn) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: "return",
	}, nil
}
