package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrElse struct {
	opcode       Opcode
	Instructions []Instr
}

func ParseInstrElse(opcode Opcode, ber *BinaryEncodingReader) (*InstrElse, error) {
	in := make([]Instr, 0)
	for {
		i, err := ParseInstr(ber)
		if err != nil {
			return nil, err
		}
		if _, ok := i.(*InstrEnd); ok {
			break
		}
		in = append(in, i)
	}

	return &InstrElse{
		opcode:       opcode,
		Instructions: in,
	}, nil
}

func (instr *InstrElse) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrElse) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrElse) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   []byte{byte(instr.opcode)},
		mnemonic: fmt.Sprintf("else"),
	}, nil
}
