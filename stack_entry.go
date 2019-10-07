package wax

import "fmt"

type StackEntry struct {
	Value *Val
	Label *Label
	Frame *Frame
}

func (e *StackEntry) String() string {
	if e.Value != nil {
		return fmt.Sprintf("(value) %s", e.Value.String())
	}
	if e.Label != nil {
		return fmt.Sprintf("(label) %s", e.Label.String())
	}
	return fmt.Sprintf("(frame) %s", e.Frame.String())
}
