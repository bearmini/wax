package wax

/*
Table Instances
https://webassembly.github.io/multi-value/core/exec/runtime.html#syntax-funcinst

A table instance is the runtime representation of a table.
It holds a vector of function elements and an optional maximum size,
if one was specified in the table type at the table’s definition site.

Each function element is either empty, representing an uninitialized table entry,
or a function address.
Function elements can be mutated through the execution of an element segment
or by external means provided by the embedder.

tableinst ::= {elem vec(funcelem), max u32?}
funcelem  ::= funcaddr?

It is an invariant of the semantics that the length of the element vector never exceeds the maximum size, if present.

Note:
Other table elements may be added in future versions of WebAssembly.
*/
type TableInst struct {
	Elem []FuncElem
	Max  *uint32
}

type FuncElem *FuncAddr

func NewTableInstances(m *Module) []TableInst {
	// TODO: implement
	return nil
}
