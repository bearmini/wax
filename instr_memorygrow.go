package wax

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type InstrMemoryGrow struct {
	opcode Opcode
}

func ParseInstrMemoryGrow(opcode Opcode, ber *BinaryEncodingReader) (*InstrMemoryGrow, error) {
	x, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	if x != 0x00 {
		return nil, errors.Errorf("expected 0x00 but found %#2x", x)
	}

	return &InstrMemoryGrow{
		opcode: opcode,
	}, nil
}

func (instr *InstrMemoryGrow) Opcode() Opcode {
	return instr.opcode
}

/*
memory.grow
1. Let F be the current frame.
2. Assert: due to validation, F.module.memaddrs[0] exists.
3. Let a be the memory address F.module.memaddrs[0].
4. Assert: due to validation, S.mems[a] exists.
5. Let mem be the memory instance S.mems[a].
6. Let sz be the length of S.mems[a] divided by the page size.
7. Assert: due to validation, a value of value type i32 is on the top of the stack.
8. Pop the value i32.const n from the stack.
9. Either, try growing mem by n pages:
(a) If it succeeds, push the value i32.const sz to the stack.
(b) Else, push the value i32.const (−1) to the stack.
10. Or, push the value i32.const (−1) to the stack.
*/
func (instr *InstrMemoryGrow) Perform(ctx context.Context, rt *Runtime) (*Label, error) {
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

	// 6. Let sz be the length of S.mems[a] divided by the page size.
	sz := uint32(len(mem.Data) / PageSize)

	// 7. Assert: due to validation, a value of value type i32 is on the top of the stack.
	err = rt.Stack.AssertTopIsValueI32()
	if err != nil {
		return nil, err
	}

	// 8. Pop the value i32.const n from the stack.
	n, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}

	// 9. Either, try growing mem by n pages:
	err = mem.TryGrowing(n.MustGetI32(), rt.cfg.maxMemorySizeInPage)

	// (a) If it succeeds, push the value i32.const sz to the stack.
	if err == nil {
		rt.Store.Mems[a] = mem
		return nil, rt.Stack.PushValue(NewValI32(sz))
	}

	// (b) Else, push the value i32.const (−1) to the stack.
	// 10. Or, push the value i32.const (−1) to the stack.
	return nil, rt.Stack.PushValue(NewValI32(0xffffffff))
}

func (instr *InstrMemoryGrow) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, 0x00),
		mnemonic: fmt.Sprintf("memory.grow 0x00"),
	}, nil
}
