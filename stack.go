package wax

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Stack struct {
	entries []*StackEntry
}

func NewStack() *Stack {
	return &Stack{
		entries: make([]*StackEntry, 0),
	}
}

func (s *Stack) Count() int {
	return len(s.entries)
}

func (s *Stack) CountValuesOnTop() int {
	n := 0
	for i := len(s.entries) - 1; i >= 0; i-- {
		if s.entries[i].Value == nil {
			break
		}
		n++
	}
	return n
}

func (s *Stack) CountLabels() int {
	n := 0
	for i := len(s.entries) - 1; i >= 0; i-- {
		if s.entries[i].Label != nil {
			n++
		}
	}
	return n
}

func (s *Stack) CountFrames() int {
	n := 0
	for i := len(s.entries) - 1; i >= 0; i-- {
		if s.entries[i].Frame != nil {
			n++
		}
	}
	return n
}

func (s *Stack) Top() *StackEntry {
	if len(s.entries) == 0 {
		return nil
	}

	return s.entries[len(s.entries)-1]
}

func (s *Stack) IsTopValue() bool {
	top := s.Top()
	if top == nil {
		return false
	}

	if top.Value == nil {
		return false
	}

	return true
}

func (s *Stack) AssertTopIsValueI32() error {
	if !s.IsTopValue() {
		return errors.New("stack top is not a value")
	}

	t, err := s.Top().Value.GetType()
	if err != nil {
		return err
	}

	if *t != ValTypeI32 {
		return errors.New("stack top is not i32")
	}

	return nil
}

func (s *Stack) GetLabelAt(labelIdx LabelIdx) (*Label, error) {
	n := 0
	for i := len(s.entries) - 1; i >= 0; i-- {
		if s.entries[i].Label != nil {
			if uint32(n) == uint32(labelIdx) {
				return s.entries[i].Label, nil
			}
			n++
		}
	}
	return nil, errors.Errorf("no label found at: %d", labelIdx)
}

func (s *Stack) PushValue(v *Val) error {
	s.entries = append(s.entries, &StackEntry{Value: v})
	return nil
}

func (s *Stack) PushValuesBack(vals []*Val) error {
	for i := len(vals) - 1; i >= 0; i-- {
		err := s.PushValue(vals[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Stack) Pop() (*StackEntry, error) {
	if len(s.entries) == 0 {
		return nil, errors.New("no items in the stack")
	}

	last := len(s.entries) - 1
	e := s.entries[last]
	s.entries = s.entries[0:last]

	return e, nil
}

func (s *Stack) PopValue() (*Val, error) {
	e, err := s.Pop()
	if err != nil {
		return nil, err
	}

	if e.Value == nil {
		return nil, errors.New("value expected")
	}

	return e.Value, nil
}

func (s *Stack) PopValues(n int) ([]*Val, error) {
	if n < 0 {
		return nil, errors.New("invalid argument")
	}
	if len(s.entries) < n {
		return nil, errors.New("not enough items in the stack")
	}

	vals := make([]*Val, n)
	for i := n - 1; i >= 0; i-- {
		v, err := s.PopValue()
		if err != nil {
			return nil, err
		}
		vals[i] = v
	}
	return vals, nil
}

func (s *Stack) Dump() string {
	lines := []string{}
	for i := len(s.entries) - 1; i >= 0; i-- {
		lines = append(lines, fmt.Sprintf("%d:\t%s", i, s.entries[i].String()))
	}
	return strings.Join(lines, "\n")
}

func (s *Stack) PushFrame(f *Frame) error {
	s.entries = append(s.entries, &StackEntry{Frame: f})
	return nil
}

func (s *Stack) PopFrame() (*Frame, error) {
	if len(s.entries) == 0 {
		return nil, errors.New("no items in the stack")
	}

	last := len(s.entries) - 1
	e := s.entries[last]
	s.entries = s.entries[0:last]

	if e.Frame == nil {
		return nil, errors.New("frame expected")
	}

	return e.Frame, nil
}

func (s *Stack) GetCurrentFrame() *Frame {
	for i := len(s.entries) - 1; i >= 0; i-- {
		if s.entries[i].Frame != nil {
			return s.entries[i].Frame
		}
	}
	return nil
}

func (s *Stack) PushLabel(l *Label) error {
	s.entries = append(s.entries, &StackEntry{Label: l})
	return nil
}

func (s *Stack) PopLabel() (*Label, error) {
	if len(s.entries) == 0 {
		return nil, errors.New("no items in the stack")
	}

	last := len(s.entries) - 1
	e := s.entries[last]
	s.entries = s.entries[0:last]

	if e.Label == nil {
		return nil, errors.New("label expected")
	}

	return e.Label, nil
}
