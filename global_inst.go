package wax

/*
Global Instances
https://webassembly.github.io/multi-value/core/exec/runtime.html#memory-instances

A global instance is the runtime representation of a global variable.
It holds an individual value and a flag indicating whether it is mutable.

globalinst ::= {value val, mut mut}

The value of mutable globals can be mutated through variable instructions or by external means provided by the embedder.
*/
type GlobalInst struct {
	Value Val
	Mut   Mut
}

func NewGlobalInstances(m *Module) []GlobalInst {
	// TODO: implement
	return nil
}
