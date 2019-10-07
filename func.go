package wax

/*
Functions
https://webassembly.github.io/multi-value/core/syntax/modules.html#syntax-func

The 'funcs' component of a module defines a vector of functions with the following structure:

func ::= {type typeidx, locals vec(valtype), body expr}


The 'type' of a function declares its signature by reference to a type defined in the 'module'.
The parameters of the function are referenced through 0-based local indices in the function’s body; they are mutable.

The 'locals' declare a vector of mutable local variables and their types.
These variables are referenced through local indices in the function’s body.
The index of the first local is the smallest index not referencing a parameter.

The 'body' is an instruction sequence that upon termination must produce a stack matching the function type’s result type.

Functions are referenced through function indices, starting with the smallest index not referencing a function import.
*/
type Func struct {
	Type   TypeIdx
	Locals []ValType
	Body   Expr
}
