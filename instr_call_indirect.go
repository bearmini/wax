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
	t64, tBytes, err := ber.ReadVaruint()
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
	// 1. Let F be the current frame.
	f := rt.Stack.GetCurrentFrame()

	// 2. Assert: due to validation, F.module.tableaddrs[0] exists.
	if len(f.Module.TableAddrs) < 1 {
		return nil, errors.New("table addrs not found")
	}

	// 3. Let ta be the table address F.module.tableaddrs[0].
	ta := f.Module.TableAddrs[0]

	// 4. Assert: due to validation, S.tables[ta] exists.
	if len(rt.Store.Tables) <= int(ta) {
		return nil, errors.New("table not found")
	}

	// 5. Let tab be the table instance S.tables[ta].
	tab := rt.Store.Tables[ta]

	// 6. Assert: due to validation, F.module.types[x] exists.
	if uint32(len(f.Module.Types)) <= uint32(instr.TypeIdx) {
		return nil, errors.New("type not found")
	}

	// 7. Let ft expect be the function type F.module.types[x].
	ft := f.Module.Types[instr.TypeIdx]

	// 8. Assert: due to validation, a value with value type i32 is on the top of the stack.
	err := rt.Stack.AssertTopIsValueI32()
	if err != nil {
		return nil, err
	}

	// 9. Pop the value i32.const i from the stack.
	iv, err := rt.Stack.PopValue()
	if err != nil {
		return nil, err
	}
	i := iv.MustGetI32()

	// 10. If i is not smaller than the length of tab.elem, then:
	if i > uint32(len(tab.Elem)) {
		// (a) Trap.
		return nil, errors.New("out of range")
	}

	// 11. If tab.elem[i] is uninitialized, then:
	if tab.Elem[i] == nil {
		// (a) Trap.
		return nil, errors.New("table element is not initialized")
	}

	// 12. Let a be the function address tab.elem[i].
	a := tab.Elem[i]

	// 13. Assert: due to validation, S.funcs[𝑎] exists.
	if uint32(len(rt.Store.Funcs)) <= uint32(*a) {
		return nil, errors.New("func not found")
	}

	// 14. Let f be the function instance S.funcs[a].
	fn := rt.Store.Funcs[*a]

	// 15. Let ftactual be the function type f.type.
	ftActual := fn.Type

	// 16. If ftactual and ft expect differ, then:
	if !ftActual.EqualsTo(ft) {
		// (a) Trap.
		return nil, errors.New("unexpected func type")
	}

	// 17. Invoke the function instance at address a.
	return nil, rt.InvokeFuncAddr(ctx, *a)
}

func (instr *InstrCallIndirect) Disassemble() (*disasmLineComponents, error) {
	return &disasmLineComponents{
		binary:   append([]byte{byte(instr.opcode)}, instr.TypeIdxBytes...),
		mnemonic: fmt.Sprintf("call typeidx:%08x 0x00", instr.TypeIdx),
	}, nil
}
