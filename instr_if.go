package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrIf struct {
	opcode       Opcode
	BlockType    BlockType
	Instructions []Instr
	ElseClause   *InstrElse
}

func ParseInstrIf(opcode Opcode, ber *BinaryEncodingReader) (*InstrIf, error) {
	bt, err := ParseBlockType(ber)
	if err != nil {
		return nil, err
	}

	var elseClause *InstrElse

	in := make([]Instr, 0)
	for {
		i, err := ParseInstr(ber)
		if err != nil {
			return nil, err
		}
		if _, ok := i.(*InstrEnd); ok {
			break
		}
		if e, ok := i.(*InstrElse); ok {
			elseClause = e
			break
		}
		in = append(in, i)
	}

	return &InstrIf{
		opcode:       opcode,
		BlockType:    *bt,
		Instructions: in,
		ElseClause:   elseClause, // else clause
	}, nil
}

func (instr *InstrIf) Opcode() Opcode {
	return instr.opcode
}

/*
if [t^?] instr^*_1 else instr^*_2 end

1. Assert: due to validation, a value of value type i32 is on the top of the stack.
2. Pop the value i32.const c from the stack.
3. Let n be the arity |t^?| of the result type t^?.
4. Let L be the label whose arity is n and whose continuation is the end of the if instruction.
5. If c is non-zero, then:
  (a) Enter the block instr^*_1 with label L.
6. Else:
  (a) Enter the block instr^*_2 with label L.

(i32.const c) if [t^n] instr^*_1 else instr^*_2 end -> label_n{ε} instr^*_1 end (if c != 0)
(i32.const c) if [t^n] instr^*_1 else instr^*_2 end -> label_n{ε} instr^*_2 end (if c == 0)
*/
func (instr *InstrIf) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1. Assert: due to validation, a value of value type i32 is on the top of the stack.
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

	//2. Pop the value i32.const c from the stack.
	c, err := rt.Stack.PopValue()

	//3. Let n be the arity |t^?| of the result type t^?.
	n := 0
	if instr.BlockType != 0x40 {
		n = 1
	}

	//4. Let L be the label whose arity is n and whose continuation is the end of the if instruction.
	l := &Label{
		Arity: uint32(n),
		Instr: []Instr{},
	}

	//5. If c is non-zero, then:
	if c.MustGetI32() != 0 {
		//  (a) Enter the block instr^*_1 with label L.
		err = rt.enterInstructionsWithLabel(ctx, l, instr.Instructions)
		if err != nil {
			return nil, err
		}

		return nil, rt.exitInstructionsWithLabel()
	}

	//6. Else:
	//  (a) Enter the block instr^*_2 with label L.
	instr2 := []Instr{}
	if instr.ElseClause != nil {
		instr2 = instr.ElseClause.Instructions
	}
	err = rt.enterInstructionsWithLabel(ctx, l, instr2)
	if err != nil {
		return nil, err
	}

	return nil, rt.exitInstructionsWithLabel()
}

func (instr *InstrIf) Disassemble() (*disasmLineComponents, error) {
	nest := []*disasmLineComponents{}
	for _, instr := range instr.Instructions {
		d, err := instr.Disassemble()
		if err != nil {
			return nil, err
		}
		nest = append(nest, d)
	}
	if instr.ElseClause != nil {
		d, err := instr.ElseClause.Disassemble()
		if err != nil {
			return nil, err
		}
		nest = append(nest, d)
	}

	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("if bt:%02x", instr.BlockType),
		nest:     nest,
	}, nil
}
