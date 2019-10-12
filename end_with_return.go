package wax

type EndWithReturn struct {
}

func NewEndWithReturn() *EndWithReturn {
	return &EndWithReturn{}
}

func (e *EndWithReturn) Error() string {
	return "end with return"
}
