package wax

/*
Export Instances
http://webassembly.github.io/spec/core/exec/runtime.html#export-instances

An export instance is the runtime representation of an export.
It defines the export’s name and the associated external value.

exportinst ::= { name name, value externval }
*/
type ExportInst struct {
	Name  Name
	Value ExternVal
}
