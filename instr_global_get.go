package wax

import (
	"context"
	"fmt"
)

type InstrGlobalGet struct {
	opcode         Opcode
	GlobalIdx      GlobalIdx
	GlobalIdxBytes []byte
}

func ParseInstrGlobalGet(opcode Opcode, ber *BinaryEncodingReader) (*InstrGlobalGet, error) {
	x64, xBytes, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	x := GlobalIdx(x64)

	return &InstrGlobalGet{
		opcode:         opcode,
		GlobalIdx:      x,
		GlobalIdxBytes: xBytes,
	}, nil
}

func (instr *InstrGlobalGet) Opcode() Opcode {
	return instr.opcode
}

/*
global.get 𝑥

1. Let F be the current frame.
2. Assert: due to validation, F.module.globaladdrs[x] exists.
3. Let a be the global address F.module.globaladdrs[x].
4. Assert: due to validation, S.globals[a] exists.
5. Let glob be the global instance S.globals[a].
6. Let val be the value glob.value.
7. Push the value val to the stack.
*/
func (instr *InstrGlobalGet) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
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

	// 6. Let val be the value glob.value.
	val := glob.Value

	// 7. Push the value val to the stack.
	return nil, rt.Stack.PushValue(&val)
}

func (instr *InstrGlobalGet) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.GlobalIdxBytes...),
		mnemonic: fmt.Sprintf("global.get %08x", instr.GlobalIdx),
	}, nil
}
