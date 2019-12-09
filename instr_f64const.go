package wax

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/pkg/errors"
)

type InstrF64Const struct {
	opcode Opcode
	N      float64
	NBytes []byte
}

func ParseInstrF64Const(opcode Opcode, ber *BinaryEncodingReader) (*InstrF64Const, error) {
	b := make([]byte, 8)
	n, err := ber.Read(b)
	if err != nil {
		return nil, err
	}
	if n != len(b) {
		return nil, errors.New("insufficient data")
	}

	return &InstrF64Const{
		opcode: opcode,
		N:      math.Float64frombits(binary.LittleEndian.Uint64(b)),
		NBytes: b,
	}, nil
}

func (instr *InstrF64Const) Opcode() Opcode {
	return instr.opcode
}

func (instr *InstrF64Const) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	return nil, rt.Stack.PushValue(NewValF64(instr.N))
}

func (instr *InstrF64Const) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.NBytes...),
		mnemonic: fmt.Sprintf("f64.const %f (%016x)", instr.N, math.Float64bits(instr.N)),
	}, nil
}
