package wax

/*
Function Instances
https://webassembly.github.io/multi-value/core/exec/runtime.html#syntax-funcinst

A function instance is the runtime representation of a function.
It effectively is a closure of the original function over the runtime module instance of its originating module.
The module instance is used to resolve references to other definitions during execution of the function.

funcinst ::= {type functype, module moduleinst, code func}
					 | {type functype, hostcode hostfunc}

A host function is a function expressed outside WebAssembly but passed to a module as an import.
The definition and behavior of host functions are outside the scope of this specification.
For the purpose of this specification, it is assumed that when invoked, a host function behaves non-deterministically,
but within certain constraints that ensure the integrity of the runtime.

Note:
Function instances are immutable, and their identity is not observable by WebAssembly code.
However, the embedder might provide implicit or explicit means for distinguishing their addresses.
*/
type FuncInst struct {
	TypeIdx  TypeIdx
	Type     FuncType
	Module   *ModuleInstance
	Code     Func
	HostCode HostFunc
}

func NewFuncInstances(m *Module, mi *ModuleInstance) []FuncInst {
	result := make([]FuncInst, 0)
	is := m.GetImportSection()
	ts := m.GetTypeSection()
	cs := m.GetCodeSection()
	fs := m.GetFunctionSection()

	for _, im := range is.Imports {
		if im.DescType != ImportDescTypeFunc {
			continue
		}

		ti := im.Desc.(TypeIdx)
		funcType := ts.FuncTypes[ti]
		result = append(result, FuncInst{
			TypeIdx:  ti,
			Type:     funcType,
			Module:   mi,
			HostCode: NewHostFunc(im.Mod, im.Nm),
		})
	}

	for i, typeIdx := range fs.Types {
		funcType := ts.FuncTypes[typeIdx]
		code := cs.Code[i]
		result = append(result, FuncInst{
			TypeIdx: typeIdx,
			Type:    funcType,
			Module:  mi,
			Code:    code.Code,
		})
	}

	return result
}
