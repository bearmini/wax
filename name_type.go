package wax

type NameType uint8

const (
	NameTypeModule NameType = iota
	NameTypeFunction
	NameTypeLocal
)
