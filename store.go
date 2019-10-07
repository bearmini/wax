package wax

/*
Store

The store represents all global state that can be manipulated by WebAssembly programs.
It consists of the runtime representation of all instances of functions, tables, memories, and globals that have been allocated during the life time of the abstract machine.

Syntactically, the store is defined as a record listing the existing instances of each category:
store ::= {
	funcs   funcinst*
	tables  tableinst*
	mems    meminst*
	globals globalinst* }
*/
type Store struct {
	Funcs   []FuncInst
	Tables  []TableInst
	Mems    []MemInst
	Globals []GlobalInst
}

func NewEmptyStore() *Store {
	return &Store{
		Funcs:   []FuncInst{},
		Tables:  []TableInst{},
		Mems:    []MemInst{},
		Globals: []GlobalInst{},
	}
}

func (s *Store) GetFirstFreeFuncAddr() FuncAddr {
	return FuncAddr(len(s.Funcs))
}

func (s *Store) GetFirstFreeMemAddr() MemAddr {
	return MemAddr(len(s.Mems))
}
