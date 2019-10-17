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

1. If module is not valid, then:
  (a) Fail.
2. Assert: module is valid with external types externtype^m_im classifying its imports.
3. If the number m of imports is not equal to the number n of provided external values, then:
  (a) Fail.
4. For each external value externval_i in externval^n and external type externtype'_i in externtype^n_im, do:
  (a) If externval_i is not valid with an external type externtype_i in store S, then:
    i. Fail.
  (b) If externtype_i does not match externtype'_i, then:
    i. Fail.
5. Let val* be the vector of global initialization values determined by module and externval^n. These may be
calculated as follows.
  (a) Let moduleinst_im be the auxiliary module instance {globaladdrs globals(externval^n)} that only consists of the imported globals.
  (b) Let F_im be the auxiliary frame {module moduleinst_im, locals 𝜖}.
  (c) Push the frame 𝐹_im to the stack.
  (d) For each global global_i in module.globals, do:
    i. Let val_i be the result of evaluating the initializer expression global_i.init.
  (e) Assert: due to validation, the frame F_im is now on the top of the stack.
  (f) Pop the frame F_im from the stack.
6. Let moduleinst be a new module instance allocated from module in store S with imports externval^n and
global initializer values val*, and let S' be the extended store produced by module allocation.
7. Let F be the frame {module moduleinst, locals 𝜖}.
8. Push the frame F to the stack.
9. For each element segment elem_i in module.elem, do:
  (a) Let eoval_i be the result of evaluating the expression elem_i.offset.
  (b) Assert: due to validation, eoval_i is of the form i32.const eo_i.
  (c) Let tableidx_i be the table index elem_i.table.
  (d) Assert: due to validation, moduleinst.tableaddrs[tableidx_i] exists.
  (e) Let tableaddr_i be the table address moduleinst.tableaddrs[tableidx_i].
  (f) Assert: due to validation, S'.tables[tableaddr_i] exists.
  (g) Let tableinst_i be the table instance S'.tables[tableaddr_i].
  (h) Let eend_i be eo_i plus the length of elem_i.init.
  (i) If eend_i is larger than the length of tableinst_i.elem, then:
    i. Fail.
10. For each data segment data_i in module.data, do:
  (a) Let doval_i be the result of evaluating the expression data_i.offset.
  (b) Assert: due to validation, doval_i is of the form i32.const do_i.
  (c) Let memidx_𝑖 be the memory index data_i.data.
  (d) Assert: due to validation, moduleinst.memaddrs[memidx_i] exists.
  (e) Let memaddr_i be the memory address moduleinst.memaddrs[memidx_i].
  (f) Assert: due to validation, S'.mems[memaddr_i] exists.
  (g) Let meminst_i be the memory instance S'.mems[memaddr_𝑖].
  (h) Let dend_i be do_i plus the length of data_i.init.
  (i) If dend_i is larger than the length of meminst_i.data, then:
    i. Fail.
11. Assert: due to validation, the frame F is now on the top of the stack.
12. Pop the frame from the stack.
13. For each element segment elem_i in module.elem, do:
  (a) For each function index funcidx_ij in elem_i.init (starting with 𝑗 = 0), do:
    i. Assert: due to validation, moduleinst.funcaddrs[funcidx_ij] exists.
    ii. Let funcaddr_ij be the function address moduleinst.funcaddrs[funcidx_ij].
    iii. Replace tableinst_i.elem[eo_i + j] with funcaddr_ij.
14. For each data segment data_i in module.data, do:
  (a) For each byte b_ij in data_i.init (starting with j = 0), do:
    i. Replace meminst_i.data[do_i + j] with b_ij.
15. If the start function module.start is not empty, then:
  (a) Assert: due to validation, moduleinst.funcaddrs[module.start.func] exists.
  (b) Let funcaddr be the function address moduleinst.funcaddrs[module.start.func].
  (c) Invoke the function instance at funcaddr.
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
	//   (a) Let moduleinst_im be the auxiliary module instance {globaladdrs globals(externval^n)} that only consists of the imported globals.
	//   (b) Let F_im be the auxiliary frame {module moduleinst_im, locals 𝜖}.
	//   (c) Push the frame 𝐹_im to the stack.
	//   (d) For each global global_i in module.globals, do:
	//     i. Let val_i be the result of evaluating the initializer expression global_i.init.
	//   (e) Assert: due to validation, the frame F_im is now on the top of the stack.
	//   (f) Pop the frame F_im from the stack.
	vals := []ValType{}

	// 6. Let moduleinst be a new module instance allocated from module in store S with imports externval^n and
	// global initializer values val*, and let S' be the extended store produced by module allocation.
	mi, err := allocModuleInstance(m, s, evs, vals)
	if err != nil {
		return nil, err
	}

	// 7. Let F be the frame {module moduleinst, locals 𝜖}.

	// 8. Push the frame F to the stack.

	// 9. For each element segment elem_i in module.elem, do:
	//   (a) Let eoval_i be the result of evaluating the expression elem_i.offset.
	//   (b) Assert: due to validation, eoval_i is of the form i32.const eo_i.
	//   (c) Let tableidx_i be the table index elem_i.table.
	//   (d) Assert: due to validation, moduleinst.tableaddrs[tableidx_i] exists.
	//   (e) Let tableaddr_i be the table address moduleinst.tableaddrs[tableidx_i].
	//   (f) Assert: due to validation, S'.tables[tableaddr_i] exists.
	//   (g) Let tableinst_i be the table instance S'.tables[tableaddr_i].
	//   (h) Let eend_i be eo_i plus the length of elem_i.init.
	//   (i) If eend_i is larger than the length of tableinst_i.elem, then:
	//     i. Fail.

	// 10. For each data segment data_i in module.data, do:
	//   (a) Let doval_i be the result of evaluating the expression data_i.offset.
	//   (b) Assert: due to validation, doval_i is of the form i32.const do_i.
	//   (c) Let memidx_𝑖 be the memory index data_i.data.
	//   (d) Assert: due to validation, moduleinst.memaddrs[memidx_i] exists.
	//   (e) Let memaddr_i be the memory address moduleinst.memaddrs[memidx_i].
	//   (f) Assert: due to validation, S'.mems[memaddr_i] exists.
	//   (g) Let meminst_i be the memory instance S'.mems[memaddr_𝑖].
	//   (h) Let dend_i be do_i plus the length of data_i.init.
	//   (i) If dend_i is larger than the length of meminst_i.data, then:
	//     i. Fail.

	// 11. Assert: due to validation, the frame F is now on the top of the stack.

	// 12. Pop the frame from the stack.

	// 13. For each element segment elem_i in module.elem, do:
	//   (a) For each function index funcidx_ij in elem_i.init (starting with 𝑗 = 0), do:
	//     i. Assert: due to validation, moduleinst.funcaddrs[funcidx_ij] exists.
	//     ii. Let funcaddr_ij be the function address moduleinst.funcaddrs[funcidx_ij].
	//     iii. Replace tableinst_i.elem[eo_i + j] with funcaddr_ij.

	// 14. For each data segment data_i in module.data, do:
	//   (a) For each byte b_ij in data_i.init (starting with j = 0), do:
	//     i. Replace meminst_i.data[do_i + j] with b_ij.

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

1. Let module be the module to allocate and externval^*_im the vector of external values providing the module's
imports, and val* the initialization values of the module's globals.

2. For each function func_i in module.funcs, do:
  (a) Let funcaddr_𝑖 be the function address resulting from allocating func_i for the module instance
      moduleinst defined below.

3. For each table table_i in module.tables, do:
  (a) Let tableaddr_i be the table address resulting from allocating table_i.type.

4. For each memory mem_i in module.mems, do:
  (a) Let memaddr_i be the memory address resulting from allocating mem_i.type.

5. For each global global_i in module.globals, do:
  (a) Let globaladdr_i be the global address resulting from allocating global_i.type with initializer value val*[i].

6. Let funcaddr* be the the concatenation of the function addresses funcaddr_i in index order.

7. Let tableaddr* be the the concatenation of the table addresses tableaddr_i in index order.

8. Let memaddr* be the the concatenation of the memory addresses memaddr_i in index order.

9. Let globaladdr* be the the concatenation of the global addresses globaladdr_i in index order.

10. Let funcaddr*_mod be the list of function addresses extracted from externval*_im, concatenated with funcaddr*.

11. Let tableaddr*_mod be the list of table addresses extracted from externval*_im, concatenated with tableaddr*.

12. Let memaddr*_mod be the list of memory addresses extracted from externval*_im, concatenated with memaddr*.

13. Let globaladdr*_mod be the list of global addresses extracted from externval*_im, concatenated with globaladdr*.

14. For each export export_i in module.exports, do:
  (a) If export_i is a function export for function index x, then let externval_i be the external value func(funcaddr* mod[x]).
  (b) Else, if export_i is a table export for table index 𝑥, then let externval_i be the external value table(tableaddr* mod[x]).
  (c) Else, if export_i is a memory export for memory index x, then let externval_i be the external value mem(memaddr* mod[x]).
  (d) Else, if export_i is a global export for global index x, then let externval_i be the external value global(globaladdr* mod[x]).
  (e) Let exportinst_i be the export instance {name (export_i.name), value externval_i}.

15. Let exportinst* be the the concatenation of the export instances exportinst_i in index order.

16. Let moduleinst be the module instance {types (module.types), funcaddrs funcaddr* mod, tableaddrs tableaddr* mod, memaddrs memaddr* mod, globaladdrs globaladdr* mod, exports exportinst*}.

17. Return moduleinst.

*/
func allocModuleInstance(m *Module, s *Store, evs []ExternVal, val []ValType) (*ModuleInstance, error) {
	mi := &ModuleInstance{
		Types: m.GetTypeSection().FuncTypes, // We need this before allocating funcs
	}

	// 2. For each function func_i in module.funcs, do:
	// 6. Let funcaddr* be the the concatenation of the function addresses funcaddr_i in index order.
	funcs := m.GetFuncs()
	funcAddrs := make([]FuncAddr, len(funcs))
	for i, f := range funcs {
		// (a) Let funcaddr_𝑖 be the function address resulting from allocating func_i for the module instance moduleinst defined below.
		fa, err := allocFunc(s, mi, f)
		if err != nil {
			return nil, err
		}
		funcAddrs[i] = fa
	}

	// 3. For each table table_i in module.tables, do:


	
	// 4. For each memory mem_i in module.mems, do:
	// 8. Let memaddr* be the the concatenation of the memory addresses memaddr_i in index order.
	mems := m.GetMems()
	memAddrs := make([]MemAddr, len(mems))
	for i, mem := range mems {
		// (a) Let memaddr_i be the memory address resulting from allocating mem_i.type.
		ma, err := allocMem(s, *mem)
		if err != nil {
			return nil, err
		}
		memAddrs[i] = ma
	}

	// 5. For each global global_i in module.globals, do:
	//   (a) Let globaladdr_i be the global address resulting from allocating global_i.type with initializer value val*[i].
	// 9. Let globaladdr* be the the concatenation of the global addresses globaladdr_i in index order.

	// 10. Let funcaddr*_mod be the list of function addresses extracted from externval*_im, concatenated with funcaddr*.
	funcAddrMods := append(extractFuncAddrs(evs), funcAddrs...)

	// 11. Let tableaddr*_mod be the list of table addresses extracted from externval*_im, concatenated with tableaddr*.
	//
	// 12. Let memaddr*_mod be the list of memory addresses extracted from externval*_im, concatenated with memaddr*.
	//
	// 13. Let globaladdr*_mod be the list of global addresses extracted from externval*_im, concatenated with globaladdr*.
	//
	// 14. For each export export_i in module.exports, do:
	//   (a) If export_i is a function export for function index x, then let externval_i be the external value func(funcaddr* mod[x]).
	//   (b) Else, if export_i is a table export for table index 𝑥, then let externval_i be the external value table(tableaddr* mod[x]).
	//   (c) Else, if export_i is a memory export for memory index x, then let externval_i be the external value mem(memaddr* mod[x]).
	//   (d) Else, if export_i is a global export for global index x, then let externval_i be the external value global(globaladdr* mod[x]).
	//   (e) Let exportinst_i be the export instance {name (export_i.name), value externval_i}.
	//
	// 15. Let exportinst* be the the concatenation of the export instances exportinst_i in index order.

	// 16. Let moduleinst be the module instance {types (module.types), funcaddrs funcaddr*_mod, tableaddrs tableaddr* mod, memaddrs memaddr* mod, globaladdrs globaladdr* mod, exports exportinst*}.
	mi.FuncAddrs = funcAddrMods
	mi.MemAddrs = memAddrs

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

func zeroedMem(n uint32) []byte {
	return make([]byte, n*64*1024)
}
