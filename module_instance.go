package wax

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
func NewModuleInstance(m *Module, s *Store, ev []ExternVal, val []ValType) (*ModuleInstance, error) {
	mi := &ModuleInstance{
		Types: m.GetTypeSection().FuncTypes, // We need this before allocating funcs
	}

	// 2. For each function func_i in module.funcs, do:
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
	//   (a) Let tableaddr_i be the table address resulting from allocating table_i.type.
	
	// 4. For each memory mem_i in module.mems, do:
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
	//
	// 6. Let funcaddr* be the the concatenation of the function addresses funcaddr_i in index order.
	//
	// 7. Let tableaddr* be the the concatenation of the table addresses tableaddr_i in index order.
	//
	// 8. Let memaddr* be the the concatenation of the memory addresses memaddr_i in index order.
	//
	// 9. Let globaladdr* be the the concatenation of the global addresses globaladdr_i in index order.
	//
	// 10. Let funcaddr*_mod be the list of function addresses extracted from externval*_im, concatenated with funcaddr*.
	//
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

	// 16. Let moduleinst be the module instance {types (module.types), funcaddrs funcaddr* mod, tableaddrs tableaddr* mod, memaddrs memaddr* mod, globaladdrs globaladdr* mod, exports exportinst*}.
	mi.FuncAddrs = funcAddrs
	mi.MemAddrs = memAddrs

	// 17. Return moduleinst.
	return mi, nil
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
		Max: mt.Limits.Max,
	}

	// 5.  Append meminst to the mems of 𝑆.
	s.Mems = append(s.Mems, mi)

	// 6.  Return a.
	return a, nil
}

func zeroedMem(n uint32) []byte {
	return make([]byte, n * 64 * 1024)
}