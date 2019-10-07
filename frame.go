package wax

import (
	"fmt"
	"strings"
)

type Frame struct {
	Arity  uint32
	Module *ModuleInstance
	Locals []*Val
}

func NewFrame(arity uint32, m *ModuleInstance, locals []*Val) *Frame {
	return &Frame{
		Arity:  arity,
		Module: m,
		Locals: locals,
	}
}

func (f *Frame) String() string {
	locals := []string{}
	for i, l := range f.Locals {
		locals = append(locals, fmt.Sprintf("%d:%s", i, l.String()))
	}
	return fmt.Sprintf("Arity: %d, Locals: %+v, Module: %+v", f.Arity, strings.Join(locals, ","), f.Module)
}
