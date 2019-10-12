package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrCall struct {
	opcode       Opcode
	FuncIdx      FuncIdx
	FuncIdxBytes []byte
}

func ParseInstrCall(opcode Opcode, ber *BinaryEncodingReader) (*InstrCall, error) {
	f64, fBytes, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	f := FuncIdx(f64)

	return &InstrCall{
		opcode:       opcode,
		FuncIdx:      f,
		FuncIdxBytes: fBytes,
	}, nil
}

func (instr *InstrCall) Opcode() Opcode {
	return instr.opcode
}

/*
call x
1. Let F be the current frame.
2. Assert: due to validation, F.module.funcaddrs[x] exists.
3. Let a be the function address F.module.funcaddrs[x].
4. Invoke the function instance at address a.

	F; (call x) ˓→ F; (invoke a) (if F.module.funcaddrs[x] = a)
*/
func (instr *InstrCall) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1.  Let f be the current frame.
	f := rt.Stack.GetCurrentFrame()

	// 2. Assert: due to validation, F.module.funcaddrs[x] exists.
	if uint32(len(f.Module.FuncAddrs)) <= uint32(instr.FuncIdx) {
		return nil, errors.New("out of range")
	}

	// 3. Let a be the function address F.module.funcaddrs[x].
	a := f.Module.FuncAddrs[instr.FuncIdx]

	// 4. Invoke the function instance at address a.
	return nil, rt.InvokeFuncAddr(ctx, a)
}

func (instr *InstrCall) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.FuncIdxBytes...),
		mnemonic: fmt.Sprintf("call funcidx:%08x", instr.FuncIdx),
	}, nil
}
