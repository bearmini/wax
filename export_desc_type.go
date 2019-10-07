package wax

type ExportDescType uint8

const (
	ExportDescTypeFunc ExportDescType = iota
	ExportDescTypeTable
	ExportDescTypeMem
	ExportDescTypeGlobal
)