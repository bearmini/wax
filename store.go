package wax

import (
	"github.com/pkg/errors"
)

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

func (s *Store) GetFirstFreeTableAddr() TableAddr {
	return TableAddr(len(s.Tables))
}

func (s *Store) GetFirstFreeMemAddr() MemAddr {
	return MemAddr(len(s.Mems))
}

func (s *Store) GetFirstFreeGlobalAddr() GlobalAddr {
	return GlobalAddr(len(s.Globals))
}

func (s *Store) AssertMemInstExists(a MemAddr) error {
	if uint32(len(s.Mems)) <= uint32(a) {
		return errors.New("invalid memaddr")
	}

	return nil
}

func (s *Store) AssertGlobalInstExists(a GlobalAddr) error {
	if uint32(len(s.Globals)) <= uint32(a) {
		return errors.New("invalid globaladdr")
	}

	return nil
}
