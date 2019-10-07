package wax

import (
	"fmt"
)

type Label struct {
	Arity uint32
	Instr []Instr
}

func NewLabel(arity uint32, instr []Instr) *Label {
	return &Label{
		Arity: arity,
		Instr: instr,
	}
}
func (l *Label) String() string {
	return fmt.Sprintf("Arity: %d, Instr: (not implemented)", l.Arity)
}
