package wax

import (
	"context"
	"encoding/binary"

	"github.com/pkg/errors"
)

/*
Instructions
https://webassembly.github.io/multi-value/core/binary/instructions.html

Instructions are encoded by opcodes. Each opcode is represented by a single byte,
and is followed by the instruction’s immediate arguments, where present.

The only exception are structured control instructions,
which consist of several opcodes bracketing their nested instruction sequences.
*/
type Instr interface {
	Perform(ctx context.Context, rt *Runtime) (*Label, error)
	Disassemble() (*disasmLineComponents, error)
}

func ParseInstr(ber *BinaryEncodingReader) (Instr, error) {
	// Read opcode byte
	opcodeU8, err := ber.ReadU8()
	if err != nil {
		return nil, err
	}
	opc := Opcode(opcodeU8)

	switch opc {
	// Control Instructions
	case OpcodeUnreachable: // 0x00
		return ParseInstrUnreachable(opc, ber)
	case OpcodeNop: // 0x01
		return ParseInstrNop(opc, ber)

	case OpcodeBlock: // 0x02
		return ParseInstrBlock(opc, ber)
	case OpcodeLoop: // 0x03
		return ParseInstrLoop(opc, ber)
	case OpcodeIf: // 0x04
		return ParseInstrIf(opc, ber)
	case OpcodeElse: // 0x05
		return ParseInstrElse(opc, ber)
	case OpcodeEnd: // 0x0b
		return ParseInstrEnd(opc, ber)
	case OpcodeBr: // 0x0c
		return ParseInstrBr(opc, ber)
	case OpcodeBrIf: // 0x0d
		return ParseInstrBrIf(opc, ber)
	case OpcodeBrTable: //0x0e
		return ParseInstrBrTable(opc, ber)
	case OpcodeReturn: // 0x0f
		return ParseInstrReturn(opc, ber)
	case OpcodeCall: // 0x10
		return ParseInstrCall(opc, ber)
	case OpcodeCallIndirect: // 0x11
		return ParseInstrCallIndirect(opc, ber)

	// Parametric Instructions
	case OpcodeDrop: // 0x1a
		return ParseInstrDrop(opc, ber)
	case OpcodeSelect: // 0x1b
		return ParseInstrSelect(opc, ber)

	// Variable Instructions
	case OpcodeLocalGet: // 0x20
		return ParseInstrLocalGet(opc, ber)
	case OpcodeLocalSet: // 0x21
		return ParseInstrLocalSet(opc, ber)
	case OpcodeLocalTee: // 0x22
		return ParseInstrLocalTee(opc, ber)
	case OpcodeGlobalGet: // 0x23
		return ParseInstrGlobalGet(opc, ber)
	case OpcodeGlobalSet: // 0x24
		return ParseInstrGlobalSet(opc, ber)

	// Memory Instructions
	case OpcodeI32Load: // 0x28
		return ParseInstrI32Load(opc, ber)
	case OpcodeI64Load: // 0x29
		return ParseInstrI64Load(opc, ber)
	case OpcodeI32Load8u: // 0x2d
		return ParseInstrI32Load8u(opc, ber)
	case OpcodeI32Load16s: // 0x2e
		return ParseInstrI32Load16s(opc, ber)
	case OpcodeI32Load16u: // 0x2f
		return ParseInstrI32Load16u(opc, ber)
	case OpcodeI64Load32u: // 0x35
		return ParseInstrI64Load32u(opc, ber)
	case OpcodeI32Store: // 0x36
		return ParseInstrI32Store(opc, ber)
	case OpcodeI64Store: // 0x37
		return ParseInstrI64Store(opc, ber)
	case OpcodeI32Store8: // 0x3a
		return ParseInstrI32Store8(opc, ber)
	case OpcodeI32Store16: // 0x3b
		return ParseInstrI32Store16(opc, ber)

	case OpcodeMemorySize: // 0x3f
	return ParseInstrMemorySize(opc, ber)

	case OpcodeMemoryGrow: // 0x40
		return ParseInstrMemoryGrow(opc, ber)

	// Numeric Instructions
	case OpcodeI32Const: // 0x41
		return ParseInstrI32Const(opc, ber)
	case OpcodeI64Const: // 0x42
		return ParseInstrI64Const(opc, ber)
	case OpcodeI32Eqz: // 0x45
		return ParseInstrI32Eqz(opc, ber)
	case OpcodeI32Eq: // 0x46
		return ParseInstrI32Eq(opc, ber)
	case OpcodeI32Ne: // 0x47
		return ParseInstrI32Ne(opc, ber)
	case OpcodeI32Lts: // 0x48
		return ParseInstrI32Lts(opc, ber)
	case OpcodeI32Ltu: // 0x49
		return ParseInstrI32Ltu(opc, ber)
	case OpcodeI32Gts: // 0x4a
		return ParseInstrI32Gts(opc, ber)
	case OpcodeI32Gtu: // 0x4b
		return ParseInstrI32Gtu(opc, ber)
	case OpcodeI32Les: // 0x4c
		return ParseInstrI32Les(opc, ber)
	case OpcodeI32Leu: // 0x4d
		return ParseInstrI32Leu(opc, ber)
	case OpcodeI32Ges: // 0x4e
		return ParseInstrI32Ges(opc, ber)
	case OpcodeI32Geu: // 0x4f
		return ParseInstrI32Geu(opc, ber)
	case OpcodeI64Ltu: // 0x54
		return ParseInstrI64Ltu(opc, ber)
	case OpcodeI64Gtu: // 0x56
		return ParseInstrI64Gtu(opc, ber)
	case OpcodeI64Leu: // 0x58
		return ParseInstrI64Leu(opc, ber)
	case OpcodeI32Clz: // 0x67
		return ParseInstrI32Clz(opc, ber)
	case OpcodeI32Ctz: // 0x68
		return ParseInstrI32Ctz(opc, ber)
	case OpcodeI32Add: // 0x6a
		return ParseInstrI32Add(opc, ber)
	case OpcodeI32Sub: // 0x6b
		return ParseInstrI32Sub(opc, ber)
	case OpcodeI32Mul: // 0x6c
		return ParseInstrI32Mul(opc, ber)
	case OpcodeI32Divs: // 0x6d
		return ParseInstrI32Divs(opc, ber)
	case OpcodeI32Divu: // 0x6e
		return ParseInstrI32Divu(opc, ber)
	case OpcodeI32And: // 0x71
		return ParseInstrI32And(opc, ber)
	case OpcodeI32Or: // 0x72
		return ParseInstrI32Or(opc, ber)
	case OpcodeI32Xor: // 0x73
		return ParseInstrI32Xor(opc, ber)
	case OpcodeI32Shl: // 0x74
		return ParseInstrI32Shl(opc, ber)
	case OpcodeI32Shru: // 0x76
		return ParseInstrI32Shru(opc, ber)
	case OpcodeI32Rotl: // 0x77
		return ParseInstrI32Rotl(opc, ber)
	case OpcodeI64Sub: // 0x7d
		return ParseInstrI64Sub(opc, ber)
	case OpcodeI64Mul: // 0x7e
		return ParseInstrI64Mul(opc, ber)
	case OpcodeI64Divu: // 0x80
		return ParseInstrI64Divu(opc, ber)
	case OpcodeI64Shru: // 0x88
		return ParseInstrI64Shru(opc, ber)
	case OpcodeI32WrapI64: // 0xa7
		return ParseInstrI32WrapI64(opc, ber)
	case OpcodeI64ExtenduI32: // 0xad
		return ParseInstrI64ExtenduI32(opc, ber)

	default:
		return nil, errors.Errorf("unknown opcode: %#02x", opc)
	}
}

/*
1.  Assert: due to validation, a value of value type t is on the top of the stack.
2.  Pop the value t.const c1 from the stack.
3.  If unop t(c1) is defined, then:
			(a)  Let c be a possible result of computing unop t(c1).
			(b)  Push the value t.const c to the stack.
4.  Else:
			(a)  Trap.
*/
func unop(rt *Runtime, t ValType, f func(v1 *Val) (*Val, error)) error {
	if rt.Stack.CountValuesOnTop() < 1 {
		return errors.New("insufficient values on top of stack")
	}

	c1, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	t1, err := c1.GetType()
	if err != nil || *t1 != t {
		return errors.New("invalid operand type")
	}

	c, err := f(c1)
	if err != nil {
		return err
	}

	return rt.Stack.PushValue(c)
}

/*
1.  Assert: due to validation, two values of value type t are on the top of the stack.
2.  Pop the value t.const c2 from the stack.
3.  Pop the value t.const c1 from the stack.
4.  If binop t(c1,c2) is defined, then:
			(a)  Let c be a possible result of computing binop t(c1,c2).
			(b)  Push the value t.const c to the stack.
5.  Else:
			(a)  Trap.
*/
func binop(rt *Runtime, t ValType, f func(v1, v2 *Val) (*Val, error)) error {
	if rt.Stack.CountValuesOnTop() < 2 {
		return errors.New("insufficient values on top of stack")
	}

	c2, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	t2, err := c2.GetType()
	if err != nil || *t2 != t {
		return errors.New("invalid operand type")
	}

	c1, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	t1, err := c1.GetType()
	if err != nil || *t1 != t {
		return errors.New("invalid operand type")
	}

	c, err := f(c1, c2)
	if err != nil {
		return err
	}

	return rt.Stack.PushValue(c)
}

/*
1.  Assert: due to validation, a value of value type t is on the top of the stack.
2.  Pop the value t.const c1 from the stack.
3.  Let c be the result of computing testop t(c1).
4.  Push the value i32.const c to the stack.
*/
func testop(rt *Runtime, t ValType, f func(v1 *Val) (*Val, error)) error {
	if rt.Stack.CountValuesOnTop() < 1 {
		return errors.New("insufficient values on top of stack")
	}

	c1, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	t1, err := c1.GetType()
	if err != nil || *t1 != t {
		return errors.New("invalid operand type")
	}

	c, err := f(c1)
	if err != nil {
		return err
	}

	ct, err := c.GetType()
	if err != nil {
		return err
	}

	if *ct != ValTypeI32 {
		return errors.New("unexpected type")
	}

	return rt.Stack.PushValue(c)
}

/*
1.  Assert: due to validation, two values of value type t are on the top of the stack.
2.  Pop the value t.const c2 from the stack.
3.  Pop the value t.const c1 from the stack.
4.  Let c be the result of computing relop t(c1,c2).
5.  Push the value i32.const c to the stack.
*/
func relop(rt *Runtime, t ValType, f func(v1, v2 *Val) (*Val, error)) error {
	if rt.Stack.CountValuesOnTop() < 2 {
		return errors.New("insufficient values on top of stack")
	}

	c2, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	t2, err := c2.GetType()
	if err != nil || *t2 != t {
		return errors.New("invalid operand type")
	}

	c1, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	t1, err := c1.GetType()
	if err != nil || *t1 != t {
		return errors.New("invalid operand type")
	}

	c, err := f(c1, c2)
	if err != nil {
		return err
	}

	ct, err := c.GetType()
	if err != nil {
		return err
	}

	if *ct != ValTypeI32 {
		return errors.New("unexpected type")
	}

	return rt.Stack.PushValue(c)
}

/*
	1.  Assert: due to validation, a value of value type t1 is on the top of the stack.
	2.  Pop the value t1.const c1 from the stack.
	3.  If cvtopsx?t1,t2(c1) is defined:
				(a)  Let c2 be a possible result of computing cvtopsx?t1,t2(c1).
				(b)  Push the value t2.const c2 to the stack.
	4.  Else:
				(a)  Trap.
*/
func cvtop(rt *Runtime, t1, t2 ValType, f func(v1 *Val) (*Val, error)) error {
	if rt.Stack.CountValuesOnTop() < 1 {
		return errors.New("insufficient values on top of stack")
	}

	c1, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}

	c1t, err := c1.GetType()
	if err != nil || *c1t != t1 {
		return errors.New("invalid operand type")
	}

	c, err := f(c1)
	if err != nil {
		return err
	}

	ct, err := c.GetType()
	if err != nil {
		return err
	}

	if *ct != t2 {
		return errors.New("unexpected type")
	}

	return rt.Stack.PushValue(c)
}

/*
t.load memarg and t.loadN_sx memarg

1.  Let F be the current frame.
2.  Assert: due to validation, F.module.memaddrs[0] exists.
3.  Let a be the memory address F.module.memaddrs[0].
4.  Assert: due to validation, S.mems[a] exists.
5.  Let mem be the memory instance S.mems[a].
6.  Assert: due to validation, a value of value type i32 is on the top of the stack.
7.  Pop the value i32.const i from the stack.
8.  Let ea be the integer i+memarg.offset.
9.  If N is not part of the instruction, then:
			(a)  Let N be the bit width |t| of value type t.
10.  If ea+N/8 is larger than the length of mem.data, then:
			(a)  Trap.
11.  Let b* be the byte sequence mem.data[ea:N/8].
12.  If N and sx are part of the instruction, then:
			(a)  Let n be the integer for which bytes iN(n) = b*.
			(b)  Let c be the result of computing extend_sx N,|t|(n).
13.  Else:
			(a)  Let c be the constant for which bytes t(c) = b*.
14.  Push the value t.const c to the stack.

t:          target value type
N (largeN): size to read
sx:         signedness. "s" = signed, "u" unsigned
*/
func loadN(rt *Runtime, t ValType, largeN int, sx string, memArg MemArg) error {
	// 1.  Let F be the current frame.
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return errors.New("no frame found")
	}

	// 2.  Assert: due to validation, F.module.memaddrs[0] exists.
	if len(f.Module.MemAddrs) == 0 {
		return errors.New("out of range")
	}

	// 3.  Let a be the memory address F.module.memaddrs[0].
	a := f.Module.MemAddrs[0]

	// 4.  Assert: due to validation, S.mems[a] exists.
	if uint32(a) >= uint32(len(rt.Store.Mems)) {
		return errors.New("out of range")
	}

	// 5.  Let mem be the memory instance S.mems[a].
	mem := rt.Store.Mems[a]

	// 6.  Assert: due to validation, a value of value type i32 is on the top of the stack.
	// 7.  Pop the value i32.const i from the stack.
	i, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}
	it, err := i.GetType()
	if err != nil || *it != ValTypeI32 {
		return errors.New("type mismatch")
	}

	// 8.  Let ea be the integer i+memarg.offset.
	ea := i.MustGetI32() + memArg.Offset

	// 9.  If N is not part of the instruction, then:
	origN := largeN
	if largeN == 0 {
		// (a)  Let N be the bit width |t| of value type t.
		largeN = t.BitCount()
	}

	// 10.  If ea+N/8 is larger than the length of mem.data, then:
	if (ea + uint32(largeN/8)) > uint32(len(mem.Data)) {
		// (a)  Trap.
		return errors.New("out of range")
	}

	// 11.  Let b* be the byte sequence mem.data[ea:N/8].
	b := mem.Data[ea : ea+uint32(largeN/8)]

	// 12.  If N and sx are part of the instruction, then:
	var c *Val
	if origN != 0 && sx != "" {
		// (a)  Let n be the integer for which bytes iN(n) = b*.
		n, err := bytesToInteger(b, largeN)
		if err != nil {
			return err
		}
		// (b)  Let c be the result of computing extend_sx N,|t|(n).
		c, err = extend(sx, largeN, t, n)
		if err != nil {
			return err
		}
	} else {
		// (a)  Let c be the constant for which bytes t(c) = b*.
		c, err = bytesToInteger(b, largeN)
		if err != nil {
			return err
		}
	}

	// 14.  Push the value t.const c to the stack.
	return rt.Stack.PushValue(c)
}

func load(rt *Runtime, t ValType, memArg MemArg) error {
	return loadN(rt, t, 0, "", memArg)
}

/*
t.store memarg and t.storeN memarg

1.  Let F be the current frame.
2.  Assert: due to validation, F.module.memaddrs[0] exists.
3.  Let a be the memory address F.module.memaddrs[0].
4.  Assert: due to validation, S.mems[a] exists.
5.  Let mem be the memory instance S.mems[a].
6.  Assert: due to validation, a value of value type t is on the top of the stack.
7.  Pop the value t.const c from the stack.
8.  Assert: due to validation, a value of value type i32 is on the top of the stack.
9.  Pop the value i32.const i from the stack.
10.  Let ea be the integer i+memarg.offset.
11.  If N is not part of the instruction, then:
			(a)  Let N be the bit width |t| of value type t.
12.  If ea+N/8 is larger than the length of mem.data, then:
			(a)  Trap.
13.  If N is part of the instruction, then:
			(a)  Let n be the result of computing wrap |t|,N(c).
			(b)  Let b* be the byte sequence bytes_iN(n).
14.  Else:
			(a)  Let b* be the byte sequence bytes_t(c).
15.  Replace the bytes mem.data[ea:N/8] with b*.
*/
func storeN(rt *Runtime, t ValType, largeN int, memArg MemArg) error {
	// 1.  Let F be the current frame.
	f := rt.Stack.GetCurrentFrame()
	if f == nil {
		return errors.New("no frame found")
	}

	// 2.  Assert: due to validation, F.module.memaddrs[0] exists.
	if len(f.Module.MemAddrs) == 0 {
		return errors.New("out of range")
	}

	// 3.  Let a be the memory address F.module.memaddrs[0].
	a := f.Module.MemAddrs[0]

	// 4.  Assert: due to validation, S.mems[a] exists.
	if uint32(a) >= uint32(len(rt.Store.Mems)) {
		return errors.New("out of range")
	}

	// 5.  Let mem be the memory instance S.mems[a].
	mem := rt.Store.Mems[a]

	// 6.  Assert: due to validation, a value of value type t is on the top of the stack.
	// 7.  Pop the value t.const c from the stack.
	c, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}
	ct, err := c.GetType()
	if err != nil || *ct != ValTypeI32 {
		return errors.New("type mismatch")
	}

	// 8.  Assert: due to validation, a value of value type i32 is on the top of the stack.
	// 9.  Pop the value i32.const i from the stack.
	i, err := rt.Stack.PopValue()
	if err != nil {
		return err
	}
	it, err := i.GetType()
	if err != nil || *it != ValTypeI32 {
		return errors.New("type mismatch")
	}

	// 10.  Let ea be the integer i+memarg.offset.
	ea := i.MustGetI32() + memArg.Offset

	origN := largeN
	// 11.  If N is not part of the instruction, then:
	if largeN == 0 {
		// (a)  Let N be the bit width |t| of value type t.
		largeN = t.BitCount()
	}

	// 12.  If ea+N/8 is larger than the length of mem.data, then:
	if (ea + uint32(largeN/8)) > uint32(len(mem.Data)) {
		// (a)  Trap.
		return errors.New("out of range")
	}

	var b []byte
	// 13.  If N is part of the instruction, then:
	if origN != 0 {
		// (a)  Let n be the result of computing wrap |t|,N(c).
		n, err := wrap(t, largeN, c)
		if err != nil {
			return err
		}

		// (b)  Let b* be the byte sequence bytes_iN(n).
		b, err = valToBytes(largeN, n)
		if err != nil {
			return err
		}
	} else {
		// 14.  Else:
		// (a)  Let b* be the byte sequence bytes_t(c).
		b, err = valToBytes(t.BitCount(), c)
		if err != nil {
			return err
		}
	}

	// 15.  Replace the bytes mem.data[ea:N/8] with b*.
	err = putBytes(mem.Data, int(ea), b)
	if err != nil {
		return err
	}

	return nil
}

func store(rt *Runtime, t ValType, memArg MemArg) error {
	return storeN(rt, t, 0, memArg)
}

// largeM: bit count of source value n
// t: target type
// n: source value
func extend(sx string, largeM int, t ValType, n *Val) (*Val, error) {
	if sx == "u" {
		nt, err := n.GetType()
		if err != nil {
			return nil, err
		}
		if *nt == t {
			return n, nil
		}

		switch t {
		case ValTypeI32:
			return NewValI32(uint32(n.MustGetI64())), nil
		case ValTypeI64:
			return NewValI64(uint64(n.MustGetI32())), nil
		default:
			return nil, errors.New("invalid type")
		}
	}

	if sx != "s" {
		return nil, errors.New("invalid signedness")
	}

	// Let j be the signed interpretation of i of size M.
	// Return the two’s complement of j relative to size N.
	switch largeM {
	case 8:
		j := int8(n.MustGetI32())
		switch t {
		case ValTypeI32: // 8 -> 32
			return NewValI32(uint32(int32(j))), nil
		case ValTypeI64: // 8 -> 64
			return NewValI64(uint64(int64(j))), nil
		default:
			return nil, errors.New("invalid target type")
		}
	case 16:
		j := int16(n.MustGetI32())
		switch t {
		case ValTypeI32: // 16 -> 32
			return NewValI32(uint32(int32(j))), nil
		case ValTypeI64: // 16 -> 64
			return NewValI64(uint64(int64(j))), nil
		default:
			return nil, errors.New("invalid target type")
		}
	case 32:
		j := int32(n.MustGetI32())
		switch t {
		case ValTypeI64: // 32 -> 64
			return NewValI64(uint64(int64(j))), nil
		default:
			return nil, errors.New("invalid target type")
		}
	default:
		return nil, errors.New("invalid source bit count")
	}
}

func wrap(t ValType, largeN int, i *Val) (*Val, error) {
	switch largeN {
	case 8:
		switch t {
		case ValTypeI32:
			return NewValI32(uint32(uint8(i.MustGetI32()))), nil
		case ValTypeI64:
			return NewValI32(uint32(uint8(i.MustGetI64()))), nil
		default:
			return nil, errors.New("invalid type")
		}
	case 16:
		switch t {
		case ValTypeI32:
			return NewValI32(uint32(uint16(i.MustGetI32()))), nil
		case ValTypeI64:
			return NewValI32(uint32(uint16(i.MustGetI64()))), nil
		default:
			return nil, errors.New("invalid type")
		}
	case 32:
		switch t {
		case ValTypeI64:
			return NewValI32(uint32(i.MustGetI64())), nil
		default:
			return nil, errors.New("invalid type")
		}
	default:
		return nil, errors.New("invalid bit size")
	}
}

func bytesToInteger(b []byte, size int) (*Val, error) {
	switch size {
	case 8:
		if len(b) < 1 {
			return nil, errors.New("insufficient data")
		}
		return NewValI32(uint32(b[0])), nil
	case 16:
		if len(b) < 2 {
			return nil, errors.New("insufficient data")
		}
		return NewValI32(uint32(binary.LittleEndian.Uint16(b))), nil
	case 32:
		if len(b) < 4 {
			return nil, errors.New("insufficient data")
		}
		return NewValI32(binary.LittleEndian.Uint32(b)), nil
	case 64:
		if len(b) < 8 {
			return nil, errors.New("insufficient data")
		}
		return NewValI64(binary.LittleEndian.Uint64(b)), nil
	default:
		return nil, errors.New("invalid size")
	}
}

func valToBytes(size int, v *Val) ([]byte, error) {
	t, err := v.GetType()
	if err != nil {
		return nil, err
	}

	switch size {
	case 8:
		var b byte
		switch *t {
		case ValTypeI32:
			b = uint8(v.MustGetI32())
		case ValTypeI64:
			b = uint8(v.MustGetI64())
		default:
			return nil, errors.New("invalid type")
		}
		return []byte{b}, nil
	case 16:
		b := make([]byte, 2)
		switch *t {
		case ValTypeI32:
			binary.LittleEndian.PutUint16(b, uint16(v.MustGetI32()))
		case ValTypeI64:
			binary.LittleEndian.PutUint16(b, uint16(v.MustGetI64()))
		default:
			return nil, errors.New("invalid type")
		}
		return b, nil
	case 32:
		b := make([]byte, 4)
		switch *t {
		case ValTypeI32:
			binary.LittleEndian.PutUint32(b, uint32(v.MustGetI32()))
		case ValTypeI64:
			binary.LittleEndian.PutUint32(b, uint32(v.MustGetI64()))
		default:
			return nil, errors.New("invalid type")
		}
		return b, nil
	case 64:
		b := make([]byte, 8)
		switch *t {
		case ValTypeI32:
			binary.LittleEndian.PutUint64(b, uint64(v.MustGetI32()))
		case ValTypeI64:
			binary.LittleEndian.PutUint64(b, uint64(v.MustGetI64()))
		default:
			return nil, errors.New("invalid type")
		}
		return b, nil
	default:
		return nil, errors.New("invalid size")
	}
}

func putBytes(dst []byte, from int, src []byte) error {
	if from+len(src) > len(dst) {
		return errors.New("out of range")
	}
	for i, b := range src {
		dst[from+i] = b
	}
	return nil
}
