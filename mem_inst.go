package wax

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

func NewMemInstances(m *Module) []MemInst {
	// TODO: implement
	return nil
}
