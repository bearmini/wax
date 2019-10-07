package wax

/*
externval ::= func   funcaddr
            | table  tableaddr
            | mem    memaddr
            | global globaladdr
*/
type ExternVal struct {
	Func   FuncAddr
	Table  TableAddr
	Mem    MemAddr
	Global GlobalAddr
}
