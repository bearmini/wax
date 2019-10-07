package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrIf struct {
	Opcode       Opcode
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
		Opcode:       opcode,
		BlockType:    *bt,
		Instructions: in,
		ElseClause:   elseClause, // else clause
	}, nil
}

func (instr *InstrIf) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrIf) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.Opcode)},
		mnemonic: fmt.Sprintf("if bt:%02x", instr.BlockType),
	}, nil
}
