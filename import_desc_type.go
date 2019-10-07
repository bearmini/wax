package wax

type ImportDescType uint8

const (
	ImportDescTypeFunc ImportDescType = iota
	ImportDescTypeTable
	ImportDescTypeMem
	ImportDescTypeGlobal
)
