package wax

import (
	"github.com/pkg/errors"
)

/*
Module Instances
http://webassembly.github.io/spec/core/exec/runtime.html#module-instances

A module instance is the runtime representation of a module.
It is created by instantiating a module, and collects runtime representations of all entities that are imported, defined, or exported by the module.

moduleinst ::= { types       functype*,
                 funcaddrs   funcaddr*,
                 tableaddrs  tableaddr*,
                 memaddrs    memaddr*,
                 globaladdrs globaladdr*,
                 exports     exportinst* }

Each component references runtime instances corresponding to respective declarations from the original module
– whether imported or defined – in the order of their static indices.
Function instances, table instances, memory instances, and global instances are referenced with an indirection through their respective addresses in the store.

It is an invariant of the semantics that all export instances in a given module instance have different names.
*/
type ModuleInstance struct {
	Types       []FuncType
	FuncAddrs   []FuncAddr
	TableAddrs  []TableAddr
	MemAddrs    []MemAddr
	GlobalAddrs []GlobalAddr
	Exports     []ExportInst
}

/*
4.5.4 Instantiation

Given a store S, a module module is instantiated with a list of external values externval^n supplying the required
imports as follows.

Instantiation checks that the module is valid and the provided imports match the declared types, and may fail with
an error otherwise. Instantiation can also result in a trap from executing the start function. It is up to the embedder
to define how such conditions are reported.
*/
func NewModuleInstance(m *Module, s *Store, evs []ExternVal) (*ModuleInstance, error) {
	// 1. If module is not valid, then:
	if !m.IsValid() {
		//   (a) Fail.
		return nil, errors.New("module is not valid")
	}

	// 2. Assert: module is valid with external types externtype^m_im classifying its imports.

	// 3. If the number m of imports is not equal to the number n of provided external values, then:
	//   (a) Fail.
	imports := m.GetImports()
	if len(evs) != len(imports) {
		return nil, errors.New("number of extern values and imports are not matched")
	}
	for i, im := range m.GetImports() {
		switch im.DescType {
		case ImportDescTypeFunc:
			if evs[i].Func == nil {
				return nil, errors.Errorf("an extern value and an import are not matched @%d", i)
			}
		case ImportDescTypeTable:
			if evs[i].Table == nil {
				return nil, errors.Errorf("an extern value and an import are not matched @%d", i)
			}
		case ImportDescTypeMem:
			if evs[i].Mem == nil {
				return nil, errors.Errorf("an extern value and an import are not matched @%d", i)
			}
		case ImportDescTypeGlobal:
			if evs[i].Global == nil {
				return nil, errors.Errorf("an extern value and an import are not matched @%d", i)
			}
		default:
			return nil, errors.Errorf("unknown import desc type: %d", im.DescType)
		}
	}

	// 4. For each external value externval_i in externval^n and external type externtype'_i in externtype^n_im, do:
	//   (a) If externval_i is not valid with an external type externtype_i in store S, then:
	//     i. Fail.
	//   (b) If externtype_i does not match externtype'_i, then:
	//     i. Fail.

	// 5. Let val* be the vector of global initialization values determined by module and externval^n. These may be
	// calculated as follows.
	vals := []Val{}

	//   (a) Let moduleinst_im be the auxiliary module instance {globaladdrs globals(externval^n)} that only consists of the imported globals.
	miim := &ModuleInstance{GlobalAddrs: extractGlobalAddrs(evs)}

	//   (b) Let F_im be the auxiliary frame {module moduleinst_im, locals ε}.
	f := &Frame{Module: miim}

	//   (c) Push the frame F_im to the stack.
	rt := &Runtime{
		Stack: NewStack(),
	}
	rt.Stack.PushFrame(f)

	//   (d) For each global global_i in module.globals, do:
	globals := m.GetGlobals()
	for _, global := range globals {
		//     i. Let val_i be the result of evaluating the initializer expression global_i.init.
		val, err := rt.evaluateInitializerExpression(global.Init)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}

	//   (e) Assert: due to validation, the frame F_im is now on the top of the stack.
	if !rt.Stack.IsTopFrame() {
		return nil, errors.New("invalid global initializer executed")
	}

	//   (f) Pop the frame F_im from the stack.
	_, err := rt.Stack.Pop()
	if err != nil {
		return nil, err
	}

	// 6. Let moduleinst be a new module instance allocated from module in store S with imports externval^n and
	// global initializer values val*, and let S' be the extended store produced by module allocation.
	mi, err := allocModuleInstance(m, s, evs, vals)
	if err != nil {
		return nil, err
	}

	// 7. Let F be the frame {module moduleinst, locals ε}.

	// 8. Push the frame F to the stack.

	// 9. For each element segment elem_i in module.elem, do:
	es := m.GetElementSection()
	if es != nil {
		for _, e := range es.Elem {
			//   (a) Let eoval_i be the result of evaluating the expression elem_i.offset.
			if len(e.Offset) < 1 {
				return nil, errors.New("unsupported elem offset expression")
			}
			eoVal := e.Offset[0]

			//   (b) Assert: due to validation, eoval_i is of the form i32.const eo_i.
			if eoVal.Opcode() != OpcodeI32Const {
				return nil, errors.New("unsupported elem offset expression")
			}

			eo, ok := eoVal.(*InstrI32Const)
			if !ok {
				return nil, errors.New("unsupported elem offset expression")
			}

			//   (c) Let tableidx_i be the table index elem_i.table.
			tableIdx := e.Table

			//   (d) Assert: due to validation, moduleinst.tableaddrs[tableidx_i] exists.
			if len(mi.TableAddrs) <= int(tableIdx) {
				return nil, errors.New("invalid tableidx")
			}

			//   (e) Let tableaddr_i be the table address moduleinst.tableaddrs[tableidx_i].
			tableaddr := mi.TableAddrs[tableIdx]

			//   (f) Assert: due to validation, S'.tables[tableaddr_i] exists.
			if len(s.Tables) <= int(tableaddr) {
				return nil, errors.New("invalid tableaddr")
			}

			//   (g) Let tableinst_i be the table instance S'.tables[tableaddr_i].
			tableinst := s.Tables[tableaddr]

			//   (h) Let eend_i be eo_i plus the length of elem_i.init.
			eend := eo.N + uint32(len(e.Init))

			//   (i) If eend_i is larger than the length of tableinst_i.elem, then:
			if eend > uint32(len(tableinst.Elem)) {
				//     i. Fail.
				return nil, errors.New("invalid elem size")
			}
		}
	}

	// 10. For each data segment data_i in module.data, do:
	ds := m.GetDataSection()
	if ds != nil {
		for _, d := range ds.Data {
			//   (a) Let doval_i be the result of evaluating the expression data_i.offset.
			if len(d.Offset) < 1 {
				return nil, errors.New("unsupported data offset expression")
			}
			doVal := d.Offset[0]

			//   (b) Assert: due to validation, doval_i is of the form i32.const do_i.
			if doVal.Opcode() != OpcodeI32Const {
				return nil, errors.New("unsupported data offset expression")
			}

			do, ok := doVal.(*InstrI32Const)
			if !ok {
				return nil, errors.New("unsupported data offset expression")
			}

			//   (c) Let memidx_i be the memory index data_i.data.
			memIdx := d.Data

			//   (d) Assert: due to validation, moduleinst.memaddrs[memidx_i] exists.
			if len(mi.MemAddrs) <= int(memIdx) {
				return nil, errors.New("invalid memidx")
			}

			//   (e) Let memaddr_i be the memory address moduleinst.memaddrs[memidx_i].
			memaddr := mi.MemAddrs[memIdx]

			//   (f) Assert: due to validation, S'.mems[memaddr_i] exists.
			if len(s.Mems) <= int(memaddr) {
				return nil, errors.New("invalid memaddr")
			}

			//   (g) Let meminst_i be the memory instance S'.mems[memaddr_𝑖].
			meminst := s.Mems[memaddr]

			//   (h) Let dend_i be do_i plus the length of data_i.init.
			dend := do.N + uint32(len(d.Init))

			//   (i) If dend_i is larger than the length of meminst_i.data, then:
			if dend > uint32(len(meminst.Data)) {
				//     i. Fail.
				return nil, errors.New("invalid data size")
			}
		}
	}

	// 11. Assert: due to validation, the frame F is now on the top of the stack.

	// 12. Pop the frame from the stack.

	// 13. For each element segment elem_i in module.elem, do:
	if es != nil {
		for _, e := range es.Elem {
			//   (a) For each function index funcidx_ij in elem_i.init (starting with 𝑗 = 0), do:
			for j, fi := range e.Init {
				//     i. Assert: due to validation, moduleinst.funcaddrs[funcidx_ij] exists.
				if uint32(len(mi.FuncAddrs)) <= uint32(fi) {
					return nil, errors.New("invalid func index")
				}

				//     ii. Let funcaddr_ij be the function address moduleinst.funcaddrs[funcidx_ij].
				fa := mi.FuncAddrs[fi]

				//     iii. Replace tableinst_i.elem[eo_i + j] with funcaddr_ij.
				eoVal := e.Offset[0]
				eo := eoVal.(*InstrI32Const)
				tableIdx := e.Table
				tableaddr := mi.TableAddrs[tableIdx]
				tableinst := s.Tables[tableaddr]
				tableinst.Elem[eo.N+uint32(j)] = &fa
			}
		}
	}

	// 14. For each data segment data_i in module.data, do:
	if ds != nil {
		for _, d := range ds.Data {
			//   (a) For each byte b_ij in data_i.init (starting with j = 0), do:
			//     i. Replace meminst_i.data[do_i + j] with b_ij.
			memIdx := d.Data
			ma := mi.MemAddrs[memIdx]
			meminst := s.Mems[ma]
			doVal := d.Offset[0]
			do := doVal.(*InstrI32Const)

			copy(meminst.Data[do.N:], d.Init)
		}
	}

	// 15. If the start function module.start is not empty, then:
	//   (a) Assert: due to validation, moduleinst.funcaddrs[module.start.func] exists.
	//   (b) Let funcaddr be the function address moduleinst.funcaddrs[module.start.func].
	//   (c) Invoke the function instance at funcaddr.

	return mi, nil
}

/*
Modules
The allocation function for modules requires a suitable list of external values that are assumed to match the import
vector of the module, and a list of initialization values for the module's globals.
*/
func allocModuleInstance(m *Module, s *Store, evs []ExternVal, vals []Val) (*ModuleInstance, error) {
	// 1. Let module be the module to allocate and externval^*_im the vector of external values providing the module's imports, and val* the initialization values of the module's globals.
	mi := &ModuleInstance{
		Types: m.GetFuncTypes(), // We need this before allocating funcs
	}

	// 2. For each function func_i in module.funcs, do:
	funcs := m.GetFuncs()
	funcAddrs := make([]FuncAddr, len(funcs))
	for i, f := range funcs {
		// (a) Let funcaddr_i be the function address resulting from allocating func_i for the module instance moduleinst defined below.
		fa, err := allocFunc(s, mi, f)
		if err != nil {
			return nil, err
		}
		funcAddrs[i] = fa
	}
	// 6. Let funcaddr* be the the concatenation of the function addresses funcaddr_i in index order.

	// 3. For each table table_i in module.tables, do:
	tables := m.GetTables()
	tableAddrs := make([]TableAddr, len(tables))
	for i, t := range tables {
		// (a) Let tableaddr_i be the table address resulting from allocating table_i.type.
		ta, err := allocTable(s, mi, t)
		if err != nil {
			return nil, err
		}
		tableAddrs[i] = ta
	}
	// 7. Let tableaddr* be the the concatenation of the table addresses tableaddr_i in index order.

	// 4. For each memory mem_i in module.mems, do:
	mems := m.GetMems()
	memAddrs := make([]MemAddr, len(mems))
	for i, mem := range mems {
		// (a) Let memaddr_i be the memory address resulting from allocating mem_i.type.
		ma, err := allocMem(s, mem)
		if err != nil {
			return nil, err
		}
		memAddrs[i] = ma
	}
	// 8. Let memaddr* be the the concatenation of the memory addresses memaddr_i in index order.

	// 5. For each global global_i in module.globals, do:
	globals := m.GetGlobals()
	globalAddrs := make([]GlobalAddr, len(globals))
	for i, g := range globals {
		//   (a) Let globaladdr_i be the global address resulting from allocating global_i.type with initializer value val*[i].
		ga, err := allocGlobal(s, g.Type, vals[i])
		if err != nil {
			return nil, err
		}
		globalAddrs[i] = ga
	}
	// 9. Let globaladdr* be the the concatenation of the global addresses globaladdr_i in index order.

	// 10. Let funcaddr*_mod be the list of function addresses extracted from externval*_im, concatenated with funcaddr*.
	funcAddrMods := append(extractFuncAddrs(evs), funcAddrs...)

	// 11. Let tableaddr*_mod be the list of table addresses extracted from externval*_im, concatenated with tableaddr*.
	tableAddrMods := append(extractTableAddrs(evs), tableAddrs...)

	// 12. Let memaddr*_mod be the list of memory addresses extracted from externval*_im, concatenated with memaddr*.
	memAddrMods := append(extractMemAddrs(evs), memAddrs...)

	// 13. Let globaladdr*_mod be the list of global addresses extracted from externval*_im, concatenated with globaladdr*.
	globalAddrMods := append(extractGlobalAddrs(evs), globalAddrs...)

	// 14. For each export export_i in module.exports, do:
	exportInsts := []ExportInst{}

	//   (a) If export_i is a function export for function index x, then let externval_i be the external value func(funcaddr* mod[x]).
	//   (b) Else, if export_i is a table export for table index 𝑥, then let externval_i be the external value table(tableaddr* mod[x]).
	//   (c) Else, if export_i is a memory export for memory index x, then let externval_i be the external value mem(memaddr* mod[x]).
	//   (d) Else, if export_i is a global export for global index x, then let externval_i be the external value global(globaladdr* mod[x]).
	//   (e) Let exportinst_i be the export instance {name (export_i.name), value externval_i}.

	// 15. Let exportinst* be the the concatenation of the export instances exportinst_i in index order.

	// 16. Let moduleinst be the module instance {types (module.types), funcaddrs funcaddr*_mod, tableaddrs tableaddr* mod, memaddrs memaddr* mod, globaladdrs globaladdr* mod, exports exportinst*}.
	mi.FuncAddrs = funcAddrMods
	mi.TableAddrs = tableAddrMods
	mi.MemAddrs = memAddrMods
	mi.GlobalAddrs = globalAddrMods
	mi.Exports = exportInsts

	// 17. Return moduleinst.
	return mi, nil
}

func extractFuncAddrs(evs []ExternVal) []FuncAddr {
	result := []FuncAddr{}
	for _, ev := range evs {
		if ev.Func != nil {
			result = append(result, *ev.Func)
		}
	}
	return result
}

func extractTableAddrs(evs []ExternVal) []TableAddr {
	result := []TableAddr{}
	for _, ev := range evs {
		if ev.Table != nil {
			result = append(result, *ev.Table)
		}
	}
	return result
}

func extractMemAddrs(evs []ExternVal) []MemAddr {
	result := []MemAddr{}
	for _, ev := range evs {
		if ev.Mem != nil {
			result = append(result, *ev.Mem)
		}
	}
	return result
}

func extractGlobalAddrs(evs []ExternVal) []GlobalAddr {
	result := []GlobalAddr{}
	for _, ev := range evs {
		if ev.Global != nil {
			result = append(result, *ev.Global)
		}
	}
	return result
}

/*
4.5.3 Allocation

Function

1. Let `func` be the function to allocate and moduleinst its module instance.
2. Let `a` be the first free function address in `S`.
3. Let `functype` be the function type `moduleinst.types[func.type]``
4. Let `funcinst` be the function instance {type functype, module𝖾 moduleinst, code func}.
5. Append funcinst to the `funcs` of S.
6. Return `a`.

 allocfunc(S,func,moduleinst) = S',funcaddr
 where:
    funcaddr = |S.funcs|
    functype = moduleinst.types[func.type]
    funcinst = {type functype, module moduleinst, code func}
         S′ = S′ ⊕ {funcs funcinst}
*/
func allocFunc(s *Store, mi *ModuleInstance, f Func) (FuncAddr, error) {
	// 1. Let `func` be the function to allocate and moduleinst its module instance.

	// 2. Let `a` be the first free function address in `S`.
	a := s.GetFirstFreeFuncAddr()

	// 3. Let `functype` be the function type `moduleinst.types[func.type]``
	ft := mi.Types[f.Type]

	// 4. Let `funcinst` be the function instance {type functype, module moduleinst, code func}.
	fi := FuncInst{
		Type:   ft,
		Module: mi,
		Code:   f,
	}

	// 5. Append funcinst to the `funcs` of S.
	s.Funcs = append(s.Funcs, fi)

	// 6. Return `a`.
	return a, nil
}

/*
Tables
1. Let tabletype be the table type to allocate.
2. Let ({min n, max m?} elemtype) be the structure of table type tabletype.
3. Let a be the first free table address in S.
4. Let tableinst be the table instance {elem (ε)^n, max m?} with n empty elements.
5. Append tableinst to the tables of S.
6. Return a.
*/
func allocTable(s *Store, mi *ModuleInstance, tt TableType) (TableAddr, error) {
	// 1. Let tabletype be the table type to allocate.
	// 2. Let ({min n, max m?} elemtype) be the structure of table type tabletype.

	// 3. Let a be the first free table address in S.
	a := s.GetFirstFreeTableAddr()

	// 4. Let tableinst be the table instance {elem (ε)^n, max m?} with n empty elements.
	ti := TableInst{
		Elem: make([]FuncElem, tt.Limits.Min),
		Max:  tt.Limits.Max,
	}

	// 5. Append tableinst to the tables of S.
	s.Tables = append(s.Tables, ti)

	// 6. Return a.
	return a, nil
}

/*
Memories

1.  Let memtype be the memory type to allocate.
2.  Let {min n, max m?} be the structure of memory type memtype.
3.  Let a be the first free memory address in S.
4.  Let meminst be the memory instance {data (0x00)^n・64Ki, max m^?} that contains n pages of zeroed bytes.
5.  Append meminst to the mems of 𝑆.
6.  Return a.
*/
func allocMem(s *Store, mt MemType) (MemAddr, error) {
	// 1.  Let memtype be the memory type to allocate.
	// 2.  Let {min n, max m?} be the structure of memory type memtype.

	// 3.  Let a be the first free memory address in S.
	a := s.GetFirstFreeMemAddr()

	// 4.  Let meminst be the memory instance {data (0x00)^n・64Ki, max m^?} that contains n pages of zeroed bytes.
	mi := MemInst{
		Data: zeroedMem(mt.Limits.Min),
		Max:  mt.Limits.Max,
	}

	// 5.  Append meminst to the mems of 𝑆.
	s.Mems = append(s.Mems, mi)

	// 6.  Return a.
	return a, nil
}

/*
Globals
1. Let globaltype be the global type to allocate and val the value to initialize the global with.
2. Let mut t be the structure of global type globaltype.
3. Let a be the first free global address in S.
4. Let globalinst be the global instance {value val, mut mut}.
5. Append globalinst to the globals of S.
6. Return a.
*/
func allocGlobal(s *Store, gt GlobalType, val Val) (GlobalAddr, error) {
	// 1. Let globaltype be the global type to allocate and val the value to initialize the global with.
	// 2. Let mut t be the structure of global type globaltype.
	// 3. Let a be the first free global address in S.
	a := s.GetFirstFreeGlobalAddr()

	// 4. Let globalinst be the global instance {value val, mut mut}.
	gi := GlobalInst{
		Value: val,
		Mut:   gt.Mut,
	}

	// 5. Append globalinst to the globals of S.
	s.Globals = append(s.Globals, gi)

	// 6. Return a.
	return a, nil
}

func zeroedMem(n uint32) []byte {
	return make([]byte, n*PageSize)
}

func (mi *ModuleInstance) AssertMemAddrExists(idx MemIdx) error {
	if uint32(len(mi.MemAddrs)) <= uint32(idx) {
		return errors.New("memaddr not found")
	}

	return nil
}

func (mi *ModuleInstance) AssertGlobalAddrExists(idx GlobalIdx) error {
	if uint32(len(mi.GlobalAddrs)) <= uint32(idx) {
		return errors.New("globaladdr not found")
	}

	return nil
}
