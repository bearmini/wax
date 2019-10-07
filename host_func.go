package wax

type HostFunc struct {
	Module Name
	Name   Name
}

func NewHostFunc(mod, nm Name) HostFunc {
	return HostFunc{
		Module: mod,
		Name:   nm,
	}
}
