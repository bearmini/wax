package wax

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/pkg/errors"
)

type InstrF32Const struct {
	opcode Opcode
	N      float32
	NBytes []byte
}

func ParseInstrF32Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrF32Const, error) {
	b := make([]byte, 4)
	n, err := ber.Read(b)
	if err != nil {
		return nil, err
	}
	if n != len(b) {
		return nil, errors.New("insufficient data")
	}

	return &InstrF32Const{
		opcode: opcode,
		N:      math.Float32frombits(binary.LittleEndian.Uint32(b)),
		NBytes: b,
	}, nil
}

func (instr *InstrF32Const) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF32Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValF32(instr.N))
}

func (instr *InstrF32Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("f32.const %f (%08x)", instr.N, math.Float32bits(instr.N)),
	}, nil
}
