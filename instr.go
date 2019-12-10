package wax

import (
	"context"
	"encoding/binary"
	"math"

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
	Opcode() Opcode
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
	case OpcodeF32Load: // 0x2a
		return ParseInstrF32Load(opc, ber)
	case OpcodeF64Load: // 0x2b
		return ParseInstrF64Load(opc, ber)
	case OpcodeI32Load8s: // 0x2c
		return ParseInstrI32Load8s(opc, ber)
	case OpcodeI32Load8u: // 0x2d
		return ParseInstrI32Load8u(opc, ber)
	case OpcodeI32Load16s: // 0x2e
		return ParseInstrI32Load16s(opc, ber)
	case OpcodeI32Load16u: // 0x2f
		return ParseInstrI32Load16u(opc, ber)
	case OpcodeI64Load8s: // 0x30
		return ParseInstrI64Load8s(opc, ber)
	case OpcodeI64Load8u: // 0x31
		return ParseInstrI64Load8u(opc, ber)
	case OpcodeI64Load16s: // 0x32
		return ParseInstrI64Load16s(opc, ber)
	case OpcodeI64Load16u: // 0x33
		return ParseInstrI64Load16u(opc, ber)
	case OpcodeI64Load32s: // 0x34
		return ParseInstrI64Load32s(opc, ber)
	case OpcodeI64Load32u: // 0x35
		return ParseInstrI64Load32u(opc, ber)
	case OpcodeI32Store: // 0x36
		return ParseInstrI32Store(opc, ber)
	case OpcodeI64Store: // 0x37
		return ParseInstrI64Store(opc, ber)
	case OpcodeF32Store: // 0x38
		return ParseInstrF32Store(opc, ber)
	case OpcodeF64Store: // 0x39
		return ParseInstrF64Store(opc, ber)
	case OpcodeI32Store8: // 0x3a
		return ParseInstrI32Store8(opc, ber)
	case OpcodeI32Store16: // 0x3b
		return ParseInstrI32Store16(opc, ber)
	case OpcodeI64Store8: // 0x3c
		return ParseInstrI64Store8(opc, ber)
	case OpcodeI64Store16: // 0x3d
		return ParseInstrI64Store16(opc, ber)
	case OpcodeI64Store32: // 0x3e
		return ParseInstrI64Store32(opc, ber)
	case OpcodeMemorySize: // 0x3f
		return ParseInstrMemorySize(opc, ber)
	case OpcodeMemoryGrow: // 0x40
		return ParseInstrMemoryGrow(opc, ber)

	// Numeric Instructions
	case OpcodeI32Const: // 0x41
		return ParseInstrI32Const(opc, ber)
	case OpcodeI64Const: // 0x42
		return ParseInstrI64Const(opc, ber)
	case OpcodeF32Const: // 0x43
		return ParseInstrF32Const(opc, ber)
	case OpcodeF64Const: // 0x44
		return ParseInstrF64Const(opc, ber)
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

	case OpcodeI64Eqz: // 0x50
		return ParseInstrI64Eqz(opc, ber)
	case OpcodeI64Eq: // 0x51
		return ParseInstrI64Eq(opc, ber)
	case OpcodeI64Ne: // 0x52
		return ParseInstrI64Ne(opc, ber)
	case OpcodeI64Lts: // 0x53
		return ParseInstrI64Lts(opc, ber)
	case OpcodeI64Ltu: // 0x54
		return ParseInstrI64Ltu(opc, ber)
	case OpcodeI64Gts: // 0x55
		return ParseInstrI64Gts(opc, ber)
	case OpcodeI64Gtu: // 0x56
		return ParseInstrI64Gtu(opc, ber)
	case OpcodeI64Les: // 0x57
		return ParseInstrI64Les(opc, ber)
	case OpcodeI64Leu: // 0x58
		return ParseInstrI64Leu(opc, ber)
	case OpcodeI64Ges: // 0x59
		return ParseInstrI64Ges(opc, ber)
	case OpcodeI64Geu: // 0x5a
		return ParseInstrI64Geu(opc, ber)

	case OpcodeF32Eq: // 0x5b
		return ParseInstrF32Eq(opc, ber)
	case OpcodeF32Ne: // 0x5c
		return ParseInstrF32Ne(opc, ber)
	case OpcodeF32Lt: // 0x5d
		return ParseInstrF32Lt(opc, ber)
	case OpcodeF32Gt: // 0x5e
		return ParseInstrF32Gt(opc, ber)
	case OpcodeF32Le: // 0x5f
		return ParseInstrF32Le(opc, ber)
	case OpcodeF32Ge: // 0x60
		return ParseInstrF32Ge(opc, ber)

	case OpcodeF64Eq: // 0x61
		return ParseInstrF64Eq(opc, ber)
	case OpcodeF64Ne: // 0x62
		return ParseInstrF64Ne(opc, ber)
	case OpcodeF64Lt: // 0x63
		return ParseInstrF64Lt(opc, ber)
	case OpcodeF64Gt: // 0x64
		return ParseInstrF64Gt(opc, ber)
	case OpcodeF64Le: // 0x65
		return ParseInstrF64Le(opc, ber)
	case OpcodeF64Ge: // 0x66
		return ParseInstrF64Ge(opc, ber)

	case OpcodeI32Clz: // 0x67
		return ParseInstrI32Clz(opc, ber)
	case OpcodeI32Ctz: // 0x68
		return ParseInstrI32Ctz(opc, ber)
	case OpcodeI32Popcnt: // 0x69
		return ParseInstrI32Popcnt(opc, ber)
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
	case OpcodeI32Rems: // 0x6f
		return ParseInstrI32Rems(opc, ber)
	case OpcodeI32Remu: // 0x70
		return ParseInstrI32Remu(opc, ber)
	case OpcodeI32And: // 0x71
		return ParseInstrI32And(opc, ber)
	case OpcodeI32Or: // 0x72
		return ParseInstrI32Or(opc, ber)
	case OpcodeI32Xor: // 0x73
		return ParseInstrI32Xor(opc, ber)
	case OpcodeI32Shl: // 0x74
		return ParseInstrI32Shl(opc, ber)
	case OpcodeI32Shrs: // 0x75
		return ParseInstrI32Shrs(opc, ber)
	case OpcodeI32Shru: // 0x76
		return ParseInstrI32Shru(opc, ber)
	case OpcodeI32Rotl: // 0x77
		return ParseInstrI32Rotl(opc, ber)
	case OpcodeI32Rotr: // 0x78
		return ParseInstrI32Rotr(opc, ber)

	case OpcodeI64Clz: // 0x79
		return ParseInstrI64Clz(opc, ber)
	case OpcodeI64Ctz: // 0x7a
		return ParseInstrI64Ctz(opc, ber)
	case OpcodeI64Popcnt: // 0x7b
		return ParseInstrI64Popcnt(opc, ber)
	case OpcodeI64Add: // 0x7c
		return ParseInstrI64Add(opc, ber)
	case OpcodeI64Sub: // 0x7d
		return ParseInstrI64Sub(opc, ber)
	case OpcodeI64Mul: // 0x7e
		return ParseInstrI64Mul(opc, ber)
	case OpcodeI64Divs: // 0x7f
		return ParseInstrI64Divs(opc, ber)
	case OpcodeI64Divu: // 0x80
		return ParseInstrI64Divu(opc, ber)
	case OpcodeI64Rems: // 0x81
		return ParseInstrI64Rems(opc, ber)
	case OpcodeI64Remu: // 0x82
		return ParseInstrI64Remu(opc, ber)
	case OpcodeI64And: // 0x83
		return ParseInstrI64And(opc, ber)
	case OpcodeI64Or: // 0x84
		return ParseInstrI64Or(opc, ber)
	case OpcodeI64Xor: // 0x85
		return ParseInstrI64Xor(opc, ber)
	case OpcodeI64Shl: // 0x86
		return ParseInstrI64Shl(opc, ber)
	case OpcodeI64Shrs: // 0x87
		return ParseInstrI64Shrs(opc, ber)
	case OpcodeI64Shru: // 0x88
		return ParseInstrI64Shru(opc, ber)
	case OpcodeI64Rotl: // 0x89
		return ParseInstrI64Rotl(opc, ber)
	case OpcodeI64Rotr: // 0x8a
		return ParseInstrI64Rotr(opc, ber)

	case OpcodeF32Abs: // 0x8b
		return ParseInstrF32Abs(opc, ber)
	case OpcodeF32Neg: // 0x8c
		return ParseInstrF32Neg(opc, ber)
	case OpcodeF32Ceil: // 0x8d
		return ParseInstrF32Ceil(opc, ber)
	case OpcodeF32Floor: // 0x8e
		return ParseInstrF32Floor(opc, ber)
	case OpcodeF32Trunc: // 0x8f
		return ParseInstrF32Trunc(opc, ber)
	case OpcodeF32Nearest: // 0x90
		return ParseInstrF32Nearest(opc, ber)
	case OpcodeF32Sqrt: // 0x91
		return ParseInstrF32Sqrt(opc, ber)
	case OpcodeF32Add: // 0x92
		return ParseInstrF32Add(opc, ber)
	case OpcodeF32Sub: // 0x93
		return ParseInstrF32Sub(opc, ber)
	case OpcodeF32Mul: // 0x94
		return ParseInstrF32Mul(opc, ber)
	case OpcodeF32Div: // 0x95
		return ParseInstrF32Div(opc, ber)
	case OpcodeF32Min: // 0x96
		return ParseInstrF32Min(opc, ber)
	case OpcodeF32Max: // 0x97
		return ParseInstrF32Max(opc, ber)
	case OpcodeF32CopySign: // 0x98
		return ParseInstrF32CopySign(opc, ber)

	case OpcodeF64Abs: // 0x99
		return ParseInstrF64Abs(opc, ber)
	case OpcodeF64Neg: // 0x9a
		return ParseInstrF64Neg(opc, ber)
	case OpcodeF64Ceil: // 0x9b
		return ParseInstrF64Ceil(opc, ber)
	case OpcodeF64Floor: // 0x9c
		return ParseInstrF64Floor(opc, ber)
	case OpcodeF64Trunc: // 0x9d
		return ParseInstrF64Trunc(opc, ber)
	case OpcodeF64Nearest: // 0x9e
		return ParseInstrF64Nearest(opc, ber)
	case OpcodeF64Sqrt: // 0x9f
		return ParseInstrF64Sqrt(opc, ber)
	case OpcodeF64Add: // 0xa0
		return ParseInstrF64Add(opc, ber)
	case OpcodeF64Sub: // 0xa1
		return ParseInstrF64Sub(opc, ber)
	case OpcodeF64Mul: // 0xa2
		return ParseInstrF64Mul(opc, ber)
	case OpcodeF64Div: // 0xa3
		return ParseInstrF64Div(opc, ber)
	case OpcodeF64Min: // 0xa4
		return ParseInstrF64Min(opc, ber)
	case OpcodeF64Max: // 0xa5
		return ParseInstrF64Max(opc, ber)
	case OpcodeF64CopySign: // 0xa6
		return ParseInstrF64CopySign(opc, ber)

	case OpcodeI32WrapI64: // 0xa7
		return ParseInstrI32WrapI64(opc, ber)
	case OpcodeI32TruncF32s: // 0xa8
		return ParseInstrI32TruncF32s(opc, ber)
	case OpcodeI32TruncF32u: // 0xa9
		return ParseInstrI32TruncF32u(opc, ber)
	case OpcodeI32TruncF64s: // 0xaa
		return ParseInstrI32TruncF64s(opc, ber)
	case OpcodeI32TruncF64u: // 0xab
		return ParseInstrI32TruncF64u(opc, ber)
	case OpcodeI64ExtendI32s: // 0xac
		return ParseInstrI64ExtendI32s(opc, ber)
	case OpcodeI64ExtendI32u: // 0xad
		return ParseInstrI64ExtendI32u(opc, ber)
	case OpcodeI64TruncF32s: // 0xae
		return ParseInstrI64TruncF32s(opc, ber)
	case OpcodeI64TruncF32u: // 0xaf
		return ParseInstrI64TruncF32u(opc, ber)
	case OpcodeI64TruncF64s: // 0xb0
		return ParseInstrI64TruncF64s(opc, ber)
	case OpcodeI64TruncF64u: // 0xb1
		return ParseInstrI64TruncF64u(opc, ber)
	case OpcodeF32ConvertI32s: // 0xb2
		return ParseInstrF32ConvertI32s(opc, ber)
	case OpcodeF32ConvertI32u: // 0xb3
		return ParseInstrF32ConvertI32u(opc, ber)
	case OpcodeF32ConvertI64s: // 0xb4
		return ParseInstrF32ConvertI64s(opc, ber)
	case OpcodeF32ConvertI64u: // 0xb5
		return ParseInstrF32ConvertI64u(opc, ber)
	case OpcodeF32DemoteF64: // 0xb6
		return ParseInstrF32DemoteF64(opc, ber)
	case OpcodeF64ConvertI32s: // 0xb7
		return ParseInstrF64ConvertI32s(opc, ber)
	case OpcodeF64ConvertI32u: // 0xb8
		return ParseInstrF64ConvertI32u(opc, ber)
	case OpcodeF64ConvertI64s: // 0xb9
		return ParseInstrF64ConvertI64s(opc, ber)
	case OpcodeF64ConvertI64u: // 0xba
		return ParseInstrF64ConvertI64u(opc, ber)
	case OpcodeF64PromoteF32: // 0xbb
		return ParseInstrF64PromoteF32(opc, ber)
	case OpcodeI32ReinterpretF32: // 0xbc
		return ParseInstrI32ReinterpretF32(opc, ber)
	case OpcodeI64ReinterpretF64: // 0xbd
		return ParseInstrI64ReinterpretF64(opc, ber)
	case OpcodeF32ReinterpretI32: // 0xbe
		return ParseInstrF32ReinterpretI32(opc, ber)
	case OpcodeF64ReinterpretI64: // 0xbf
		return ParseInstrF64ReinterpretI64(opc, ber)

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
		n, err := bytesToVal(t, b, largeN)
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
		c, err = bytesToVal(t, b, largeN)
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
	if err != nil || *ct != t {
		return errors.Errorf("type mismatch: expected %s, actual: %s", t, ct)
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

func bytesToVal(t ValType, b []byte, size int) (*Val, error) {
	switch t {
	case ValTypeI32, ValTypeI64:
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
	case ValTypeF32, ValTypeF64:
		switch size {
		case 32:
			if len(b) < 4 {
				return nil, errors.New("insufficient data")
			}
			return NewValF32(math.Float32frombits(binary.LittleEndian.Uint32(b))), nil
		case 64:
			if len(b) < 8 {
				return nil, errors.New("insufficient data")
			}
			return NewValF64(math.Float64frombits(binary.LittleEndian.Uint64(b))), nil
		default:
			return nil, errors.New("invalid size")
		}
	default:
		return nil, errors.New("invalid val type")
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
		case ValTypeF32:
			binary.LittleEndian.PutUint32(b, math.Float32bits(v.MustGetF32()))
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
		case ValTypeF64:
			binary.LittleEndian.PutUint64(b, math.Float64bits(v.MustGetF64()))
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
