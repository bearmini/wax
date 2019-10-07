package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

/*
0x03 bt:blocktype (in:instr)* 0x0B
*/
type InstrLoop struct {
	Opcode       Opcode
	BlockType    BlockType
	Instructions []Instr
}

func ParseInstrLoop(opcode Opcode, ber *BinaryEncodingReader) (*InstrLoop, error) {
	bt, err := ParseBlockType(ber)
	if err != nil {
		return nil, err
	}

	in := make([]Instr, 0)
	for {
		i, err := ParseInstr(ber)
		if err != nil {
			return nil, err
		}
		in = append(in, i)
		if _, end := i.(*InstrEnd); end {
			break
		}
	}

	return &InstrLoop{
		Opcode:       opcode,
		BlockType:    *bt,
		Instructions: in,
	}, nil
}

/*
1.  Assert: due to validation, expand F(blocktype) is defined.
2.  Let[tm1]->[tn2] be the function type expand F(blocktype).
3.  Let L be the label whose arity is m and whose continuation is the start of the loop.
4.  Enterthe block instr* with label L
*/
func (instr *InstrLoop) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return nil, errors.New("frame not found")
	}

	// 1. Assert: due to validation, expand F(blocktype) is defined.
	//       expand F(typeidx)    = F.module.types[typeidx]
	//       expand F([valtype?]) = [] -> [valtype?]
	// 2. Let [tm1]→[tn2] be the function type expand F(blocktype).

	m := 0
	if instr.BlockType != 0x40 {
		m = 1
	}

	// 3. Let L be the label whose arity is m and whose continuation is the start of the block.
	l := Label{
		Arity: uint32(m),
		Instr: instr.Instructions,
	}

	// 4. Enter the block instr* with label L.
	err := rt.enterInstructionsWithLabel(ctx, &l, instr.Instructions)
	if err != nil {
		return nil, err
	}

	/*  
	exitInstructionsWithLabel() should be called only When the end of a block is reached without a jump.
	The loop instruction always end with a branch (jump). So we don't need call exitInstructionsWithLabel().
	err = rt.exitInstructionsWithLabel()
	if err != nil {
		return nil, err
	}
	*/
	return nil, nil
}

func (instr *InstrLoop) Disassemble() (*disasmLineComponents, error) {
	nest := []*disasmLineComponents{}
	for _, instr := range instr.Instructions {
		d, err := instr.Disassemble()
		if err != nil {
			return nil, err
		}
		nest = append(nest, d)
	}

	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("loop bt:%02x", instr.BlockType),
		nest:     nest,
	}, nil
}
