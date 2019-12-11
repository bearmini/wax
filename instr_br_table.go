package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrBrTable struct {
	opcode  Opcode
	Ls      []LabelIdx
	LsBytes []byte
	Ln      LabelIdx
	LnBytes []byte
}

func ParseInstrBrTable(opcode Opcode, ber *BinaryEncodingReader) (*InstrBrTable, error) {
	n, nBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}

	ls := make([]LabelIdx, 0)
	lsBytes := make([]byte, 0)

	for i := uint64(0); i < n; i++ {
		l64, lBytes, err := ber.ReadVaruint()
		if err != nil {
			return nil, err
		}
		l := LabelIdx(l64)
		ls = append(ls, l)
		lsBytes = append(lsBytes, lBytes...)
	}

	ln64, lnBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	ln := LabelIdx(ln64)

	return &InstrBrTable{
		opcode:  opcode,
		Ls:      ls,
		LsBytes: append(nBytes, lsBytes...),
		Ln:      ln,
		LnBytes: lnBytes,
	}, nil
}

func (instr *InstrBrTable) Opcode() Opcode {
	return instr.opcode
}

/*
br_table 𝑙* lN

1. Assert: due to validation, a value of value type i32 is on the top of the stack.
2. Pop the value i32.const i from the stack.
3. If i is smaller than the length of l*, then:
  (a) Let li be the label l*[i].
  (b) Execute the instruction (br li).
4. Else:
  (a) Execute the instruction (br lN).

      (i32.const i) (br_table l*l𝑁) → (br li) (if l*[i] = li)
      (i32.const i) (br_table l*lN) → (br lN) (if |l*| <= i)
*/
func (instr *InstrBrTable) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1.  Assert: due to validation, a value of value type i32 is on the top of the stack.
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

	// 2.  Pop the value i32.const i from the stack.
	iv, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}
	i := iv.MustGetI32()

	// 3. If i is smaller than the length of l*, then:
	if i < uint32(len(instr.Ls)) {
		// (a) Let li be the label l*[i].
		li := instr.Ls[i]

		// (b) Execute the instruction (br li).
		ib := InstrBr{LabelIdx: li}
		return ib.Perform(ctx, rt)
	}

	// 4. Else:
	//   (a) Execute the instruction (br lN).
	ib := InstrBr{LabelIdx: instr.Ln}
	return ib.Perform(ctx, rt)
}

func (instr *InstrBrTable) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append(append([]byte{byte(instr.opcode)}, instr.LsBytes...), instr.LnBytes...),
		mnemonic: fmt.Sprintf("br_table l*:%v, lN:%x", instr.Ls, instr.Ln),
	}, nil
}
