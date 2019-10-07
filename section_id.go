package wax

type SectionID uint8

const (
	CustomSectionID SectionID = iota
	TypeSectionID
	ImportSectionID
	FunctionSectionID
	TableSectionID
	MemorySectionID
	GlobalSectionID
	ExportSectionID
	StartSectionID
	ElementSectionID
	CodeSectionID
	DataSectionID
)
