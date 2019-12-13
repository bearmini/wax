package wax

import (
	"github.com/pkg/errors"
)

/*
Memory Instances
https://webassembly.github.io/multi-value/core/exec/runtime.html#memory-instances

A memory instance is the runtime representation of a linear memory.
It holds a vector of bytes and an optional maximum size, if one was specified at the definition site of the memory.

meminst ::= {data vec(byte), max u32?}

The length of the vector always is a multiple of the WebAssembly page size, which is defined to be the constant 65536
– abbreviated 64Ki. Like in a memory type, the maximum size in a memory instance is given in units of this page size.

The bytes can be mutated through memory instructions, the execution of a data segment, or by external means provided by the embedder.

It is an invariant of the semantics that the length of the byte vector, divided by page size, never exceeds the maximum size, if present.
*/
type MemInst struct {
	Data []byte
	Max  *uint32
}

/*
Growing memories
1. Let meminst be the memory instance to grow and 𝑛 the number of pages by which to grow it.
2. Assert: The length of meminst.data is divisible by the page size 64 Ki.
3. Let len be 𝑛 added to the length of meminst.data divided by the page size 64 Ki.
4. If len is larger than 2
16, then fail.
5. If meminst.max is not empty and its value is smaller than len, then fail.
6. Append 𝑛 times 64 Ki bytes with value 0x00 to meminst.data.
*/
func (mi *MemInst) TryGrowing(n, hardMax uint32) error {
	// 1. Let meminst be the memory instance to grow and 𝑛 the number of pages by which to grow it.

	// 2. Assert: The length of meminst.data is divisible by the page size 64 Ki.
	if len(mi.Data)%PageSize != 0 {
		return errors.New("invalid mem size")
	}
	cur := uint32(len(mi.Data) / PageSize)

	// 3. Let len be 𝑛 added to the length of meminst.data divided by the page size 64 Ki.
	l := n + cur

	// 4. If len is larger than 2^16, then fail.
	if l > 65536 {
		return errors.New("out of memory")
	}

	// 5. If meminst.max is not empty and its value is smaller than len, then fail.
	if mi.Max != nil && *mi.Max < l {
		return errors.New("out of memory")
	}
	if hardMax < l {
		return errors.New("out of memory")
	}

	// 6. Append 𝑛 times 64 Ki bytes with value 0x00 to meminst.data.
	mi.Data = append(mi.Data, zeroedMem(n)...)

	return nil
}
