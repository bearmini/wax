package wax

import (
	"context"
	"fmt"
)

type InstrGlobalSet struct {
	opcode         Opcode
	GlobalIdx      GlobalIdx
	GlobalIdxBytes []byte
}

func ParseInstrGlobalSet(opcode Opcode, ber *BinaryEncodingReader) (*InstrGlobalSet, error) {
	x64, xBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	x := GlobalIdx(x64)

	return &InstrGlobalSet{
		opcode:         opcode,
		GlobalIdx:      x,
		GlobalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrGlobalSet) Opcode() Opcode {
	return instr.opcode
}

/*
global.set 𝑥

1. Let F be the current frame.
2. Assert: due to validation, F.module.globaladdrs[x] exists.
3. Let a be the global address F.module.globaladdrs[x].
4. Assert: due to validation, S.globals[a] exists.
5. Let glob be the global instance S.globals[a].
6. Assert: due to validation, a value is on the top of the stack.
7. Pop the value val from the stack.
8. Replace glob.value with the value val.
*/
func (instr *InstrGlobalSet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1. Let F be the current frame.
	f := rt.Stack.GetCurrentFrame()

	// 2. Assert: due to validation, F.module.globaladdrs[x] exists.
	err := f.Module.AssertGlobalAddrExists(instr.GlobalIdx)
	if err != nil {
		return nil, err
	}

	// 3. Let a be the global address F.module.globaladdrs[x].
	a := f.Module.GlobalAddrs[instr.GlobalIdx]

	// 4. Assert: due to validation, S.globals[a] exists.
	err = rt.Store.AssertGlobalInstExists(a)
	if err != nil {
		return nil, err
	}

	// 5. Let glob be the global instance S.globals[a].
	glob := rt.Store.Globals[a]

	// 6. Assert: due to validation, a value is on the top of the stack.
	err = rt.Stack.AssertTopIsValue()
	if err != nil {
		return nil, err
	}

	// 7. Pop the value val from the stack.
	val, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	// 8. Replace glob.value with the value val.
	glob.Value = *val
	rt.Store.Globals[a] = glob

	return nil, nil
}

func (instr *InstrGlobalSet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.GlobalIdxBytes...),
		mnemonic: fmt.Sprintf("global.set %08x", instr.GlobalIdx),
	}, nil
}
