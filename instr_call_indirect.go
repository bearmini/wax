package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrCallIndirect struct {
	opcode       Opcode
	TypeIdx      TypeIdx
	TypeIdxBytes []byte
}

func ParseInstrCallIndirect(opcode Opcode, ber *BinaryEncodingReader) (*InstrCallIndirect, error) {
	t64, tBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	t := TypeIdx(t64)

	b, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	if b != 0x00 {
		return nil, errors.New("invalid value")
	}

	return &InstrCallIndirect{
		opcode:       opcode,
		TypeIdx:      t,
		TypeIdxBytes: tBytes,
	}, nil
}

func (instr *InstrCallIndirect) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrCallIndirect) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, errors.New("not implemented")
}

func (instr *InstrCallIndirect) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.TypeIdxBytes...),
		mnemonic: fmt.Sprintf("call typeidx:%08x 0x00", instr.TypeIdx),
	}, nil
}
