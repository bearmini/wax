package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

/*
0x02 bt:blocktype (in:instr)* 0x0B
*/
type InstrBlock struct {
	opcode       Opcode
	BlockType    BlockType
	Instructions []Instr
}

func NewInstrBlock(blockType BlockType, instructions []Instr) *InstrBlock {
	return &InstrBlock{
		opcode:       OpcodeBlock,
		BlockType:    blockType,
		Instructions: instructions,
	}
}

func ParseInstrBlock(opcode Opcode, ber *BinaryEncodingReader) (*InstrBlock, error) {
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

	return &InstrBlock{
		opcode:       opcode,
		BlockType:    *bt,
		Instructions: in,
	}, nil
}

func (instr *InstrBlock) Opcode() Opcode {
	return instr.opcode
}

/*
1. Assert: due to validation, expand F(blocktype) is defined.
2. Let [tm1]→[tn2] be the function type expand F(blocktype).
3. Let L be the label whose arity is n and whose continuation is the end of the block.
4. Enter the block instr* with label L.
*/
func (instr *InstrBlock) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return nil, errors.New("frame not found")
	}

	// 1. Assert: due to validation, expand F(blocktype) is defined.
	//       expand F(typeidx)    = F.module.types[typeidx]
	//       expand F([valtype?]) = [] -> [valtype?]
	// 2. Let [tm1]→[tn2] be the function type expand F(blocktype).

	n := 0
	if instr.BlockType != 0x40 {
		n = 1
	}

	// 3. Let L be the label whose arity is n and whose continuation is the end of the block.
	l := &Label{
		Arity: uint32(n),
		Instr: []Instr{},
	}

	// 4. Enter the block instr* with label L.
	err := rt.enterInstructionsWithLabel(ctx, l, instr.Instructions)
	if err != nil {
		return nil, err
	}

	return nil, rt.exitInstructionsWithLabel()
}

func (instr *InstrBlock) Disassemble() (*disasmLineComponents, error) {
	nest := []*disasmLineComponents{}
	for _, instr := range instr.Instructions {
		d, err := instr.Disassemble()
		if err != nil {
			return nil, err
		}
		nest = append(nest, d)
	}

	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("block bt:%02x", instr.BlockType),
		nest:     nest,
	}, nil
}
