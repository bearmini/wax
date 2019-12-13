package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrMemorySize struct {
	opcode Opcode
}

func ParseInstrMemorySize(opcode Opcode, ber *BinaryEncodingReader) (*InstrMemorySize, error) {
	x, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	if x != 0x00 {
		return nil, errors.Errorf("expected 0x00 but found %#2x", x)
	}

	return &InstrMemorySize{
		opcode: opcode,
	}, nil
}

func (instr *InstrMemorySize) Opcode() Opcode {
	return instr.opcode
}

/*
memory.size
1. Let F be the current frame.
2. Assert: due to validation, F.module.memaddrs[0] exists.
3. Let a be the memory address F.module.memaddrs[0].
4. Assert: due to validation, S.mems[a] exists.
5. Let mem be the memory instance S.mems[a].
6. Let sz be the length of mem.data divided by the page size.
7. Push the value i32.const sz to the stack.
*/
func (instr *InstrMemorySize) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
	// 1. Let F be the current frame.
	f := rt.Stack.GetCurrentFrame()

	// 2. Assert: due to validation, F.module.memaddrs[0] exists.
	err := f.Module.AssertMemAddrExists(0)
	if err != nil {
		return nil, err
	}

	// 3. Let a be the memory address F.module.memaddrs[0].
	a := f.Module.MemAddrs[0]

	// 4. Assert: due to validation, S.mems[a] exists.
	err = rt.Store.AssertMemInstExists(a)
	if err != nil {
		return nil, err
	}

	// 5. Let mem be the memory instance S.mems[a].
	mem := rt.Store.Mems[a]

	// 6. Let sz be the length of mem.data divided by the page size.
	sz := uint32(len(mem.Data) / PageSize)

	// 7. Push the value i32.const sz to the stack.
	return nil, rt.Stack.PushValue(NewValI32(sz))
}

func (instr *InstrMemorySize) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, 0x00),
		mnemonic: fmt.Sprintf("memory.size 0x00"),
	}, nil
}
