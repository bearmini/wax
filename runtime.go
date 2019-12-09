package wax

import (
	"context"

	"github.com/pkg/errors"
)

type Runtime struct {
	ip             uint32 // instruction pointer
	Store          *Store
	Stack          *Stack
	Module         *Module
	ModuleInstance *ModuleInstance
}

func NewRuntime(m *Module) (*Runtime, error) {
	//s := NewStore(m)
	s := NewEmptyStore()
	e, err := createExternValsForImportSection(m)
	if err != nil {
		return nil, err
	}

	mi, err := NewModuleInstance(m, s, e)
	if err != nil {
		return nil, err
	}

	return &Runtime{
		Store:          s,
		Stack:          NewStack(),
		Module:         m,
		ModuleInstance: mi,
	}, nil
}

func createExternValsForImportSection(m *Module) ([]ExternVal, error) {
	result := []ExternVal{}

	imports := m.GetImports()
	nFunctions := 0
	for _, im := range imports {
		switch im.DescType {
		case ImportDescTypeFunc:
			fa := FuncAddr(nFunctions)
			result = append(result, ExternVal{Func: &fa})
		default:
			return nil, errors.New("unsupported import desc type")
		}
	}
	return result, nil
}

func (rt *Runtime) Dump() string {
	s := ""
	if rt.Stack != nil {
		s += "\nStack:\n" + rt.Stack.Dump()
	}
	return s
}

func (rt *Runtime) FindFuncAddr(fname string) (*FuncAddr, error) {
	fi := rt.findFuncIdxByName(fname)
	if fi == nil {
		return nil, errors.Errorf("func not found: %s", fname)
	}

	nif := uint32(0) // number of imported functions
	is := rt.Module.GetImportSection()
	if is != nil {
		nif = is.GetFuncImportsCount()
	}

	if uint32(*fi) >= uint32(len(rt.Store.Funcs))+nif {
		return nil, errors.Errorf("out of range")
	}

	var fa FuncAddr
	fa = FuncAddr(uint32(*fi) - nif)

	return &fa, nil
}

func (rt *Runtime) findFuncIdxByName(fname string) *FuncIdx {
	es := rt.Module.GetExportSection()
	for _, e := range es.Exports {
		fi, ok := e.Desc.(*FuncIdx)
		if !ok {
			continue
		}

		if e.Nm == Name(fname) {
			return fi
		}
	}
	return nil
}

/*
Invocation
https://webassembly.github.io/multi-value/core/exec/modules.html#invocation

Once a module has been instantiated, any exported function can be invoked externally via its function address funcaddr
in the store S and an appropriate list val* of argument values.

Invocation may fail with an error if the arguments do not fit the function type.
Invocation can also result in a trap. It is up to the embedder to define how such conditions are reported.

Note:
If the embedder API performs type checks itself, either statically or dynamically, before performing an invocation, then no failure other than traps can occur.

The following steps are performed:
1. Assert: S.funcs[funcaddr] exists.
2. Let funcinst be the function instance S.funcs[funcaddr].
3. Let [tn_1]→[tm_2] be the function type funcinst.type.
4. If the length |val*| of the provided argument values is different from the number n of expected arguments, then:
  a. Fail.
5. For each value type ti in tn_1 and corresponding value val_i in val*, do:
  a. If val_i is not ti.const ci for some ci, then:
    i. Fail.
6. Push the values val* to the stack.
7. Invoke the function instance at address funcaddr.

Once the function has returned, the following steps are executed:
1.  Assert: due tovalidation, m values are on the top of the stack.
2.  Pop val_res^m from the stack.
The values val_res^m are returned as the results of the invocation.
*/
func (rt *Runtime) InvokeFunc(ctx context.Context, fa FuncAddr, val []*Val) ([]*Val, error) {
	// 1. Assert: S.funcs[funcaddr] exists.
	if uint32(fa) >= uint32(len(rt.Store.Funcs)) {
		return nil, errors.Errorf("no func found with funcaddr %#08x", fa)
	}

	// 2. Let funcinst be the function instance S.funcs[funcaddr].
	fi := rt.Store.Funcs[fa]

	// 3. Let [tn1]→[tm2] be the function type funcinst.type.
	ft := fi.Type

	// 4. If the length |val*| of the provided argument values is different from the number n of expected arguments, then:
	//    a. Fail.
	if len(val) != len(ft.ParamTypes) {
		return nil, errors.Errorf("number of expected arguments is %d, but provided %d", len(ft.ParamTypes), len(val))
	}

	// 5. For each value type ti in tn_1 and corresponding value val_i in val*, do:
	//    a. If val_i is not ti.const ci for some ci, then:
	//       i. Fail.
	for i := range ft.ParamTypes {
		v := val[i]
		t, err := v.GetType()
		if err != nil {
			return nil, err
		}
		if *t != ft.ParamTypes[i] {
			return nil, errors.New("argument type does not match")
		}
	}

	// 6. Push the values val* to the stack.
	for i := len(val) - 1; i >= 0; i-- {
		err := rt.Stack.PushValue(val[i])
		if err != nil {
			return nil, err
		}
	}

	// 7. Invoke the function instance at address funcaddr.
	err := rt.InvokeFuncAddr(ctx, fa)
	if err != nil {
		return nil, err
	}

	// Once the function has returned, the following steps are executed:
	// 1.  Assert: due tovalidation, m values are on the top of the stack.
	if rt.Stack.CountValuesOnTop() != len(ft.ReturnTypes) {
		return nil, errors.New("return type mismatch")
	}

	// 2.  Pop val_res^m from the stack.
	// The values val_res^m are returned as the results of the invocation.
	return rt.Stack.PopValues(len(ft.ReturnTypes))
}

/*
Invocation of function address a
https://webassembly.github.io/multi-value/core/exec/instructions.html#invocation-of-function-address

1. Assert: due to validation, S.funcs[a] exists.
2. Let f be the function instance, S.funcs[a].
3. Let [tn_1]→[tm_2] be the function type f.type.
4. Let t* be the list of value types f.code.locals.
5. Let instr* be the expression f.code.body.
6. Assert: due to validation, n values are on top of the stack.
7. Pop the values valn from the stack.
8. Let val*_0 be the list of zero values of types t*.
9. Let F be the frame { module f.module, locals valn val*_0 }.
10. Push the activation of F with arity m to the stack.
11. Let L be the label whose arity is m and whose continuation is the end of the function.
12. Enter the instruction sequnce instr* with label L.

Returning from a function
When the end of a function is reached without a jump (i.e.,return) or trap aborting it,
then the following steps are performed.
1.  Let f be the current frame.
2.  Let n be the arity of the activation of F.
3.  Assert: due to validation, there are n values on the top of the stack.
4.  Pop the results val^n from the stack.
5.  Assert: due to validation, the frame F is now on the top of the stack.
6.  Pop the frame from the stack.
7.  Push val^n back to the stack.
8.  Jump to the instruction after the original call.
*/
func (rt *Runtime) InvokeFuncAddr(ctx context.Context, a FuncAddr) error {
	// 1. Assert: due to validation, S.funcs[a] exists.
	if uint32(a) >= uint32(len(rt.Store.Funcs)) {
		return errors.Errorf("no func found with funcaddr %#08x", a)
	}

	// 2. Let f be the function instance, S.funcs[a].
	f := rt.Store.Funcs[a]

	// 3. Let [tn_1]→[tm_2] be the function type f.type.
	ft := f.Type
	n := len(ft.ParamTypes)
	m := uint32(len(ft.ReturnTypes))

	// 4. Let t* be the list of value types f.code.locals.
	t := f.Code.Locals

	// 5. Let instr* be the expression f.code.body.
	instr := f.Code.Body

	// 6. Assert: due to validation, n values are on top of the stack.
	if rt.Stack.Count() < n {
		return errors.New("values on the stack are not enough")
	}

	// 7. Pop the values val^n from the stack.
	vals, err := rt.Stack.PopValues(n)
	if err != nil {
		return err
	}

	// 8. Let val*_0 be the list of zero values of types t*.
	for _, vt := range t {
		v, err := NewZeroVal(vt)
		if err != nil {
			return err
		}
		vals = append(vals, v)
	}

	// 9. Let F be the frame { module f.module, locals val^n val*_0 }.
	fr := NewFrame(m, f.Module, vals)

	// 10. Push the activation of F with arity m to the stack.
	err = rt.Stack.PushFrame(fr)
	if err != nil {
		return err
	}

	// 11. Let L be the label whose arity is m and whose continuation is the end of the function.
	l := NewLabel(m, f.Code.Body[len(f.Code.Body)-1:])

	// 12. Enter the instruction sequnce instr* with label L.
	err = rt.enterInstructionsWithLabel(ctx, l, instr)
	if err != nil {
		if _, ok := err.(*EndWithReturn); ok {
			return nil
		}
		if ewj, ok := err.(*EndWithJump); ok {
			ewj.LabelsExited--
			if ewj.LabelsExited != 0 {
				return err
			}
			// if we get ewj.LabelsExited == 0, we want to continue exit process
		} else {
			return err
		}
	}

	err = rt.exitInstructionsWithLabel()
	if err != nil {
		return err
	}

	// 1.  Let f be the current frame.
	fr = rt.Stack.GetCurrentFrame()

	// 2.  Let n be the arity of the activation of F.
	n2 := fr.Arity

	// 3.  Assert: due to validation, there are n values on the top of the stack.
	if uint32(rt.Stack.CountValuesOnTop()) != n2 {
		return errors.New("number of returns unmatch")
	}

	// 4.  Pop the results val^n from the stack.
	retvals, err := rt.Stack.PopValues(int(n2))
	if err != nil {
		return err
	}

	// 5.  Assert: due to validation, the frame F is now on the top of the stack.
	if rt.Stack.Top() == nil || rt.Stack.Top().Frame == nil {
		return errors.New("frame must be on top of stack")
	}

	// 6.  Pop the frame from the stack.
	_, err = rt.Stack.PopFrame()
	if err != nil {
		return err
	}

	// 7.  Push val^n back to the stack.
	err = rt.Stack.PushValuesBack(retvals)
	if err != nil {
		return err
	}

	// 8.  Jump to the instruction after the original call.
	return nil
}

func zeroValue(vt ValType) interface{} {
	switch vt {
	case ValTypeI32:
		return uint32(0)
	case ValTypeI64:
		return uint64(0)
	case ValTypeF32:
		return float32(0)
	case ValTypeF64:
		return float64(0)
	default:
		return nil
	}
}

/*
Entering instr* with label L
https://webassembly.github.io/multi-value/core/exec/instructions.html#exec-instr-seq-enter

	1. Push L to the stack.
	2. Jump to the start of the instruction sequence instr*.

Note:
No formal reduction rule is needed for entering an instruction sequence, because the label L
is embedded in the administrative instruction that structured control instructions reduce to directly.
*/
func (rt *Runtime) enterInstructionsWithLabel(ctx context.Context, l *Label, instructions []Instr) error {
	err := rt.Stack.PushLabel(l)
	if err != nil {
		return err
	}

	ip := 0
	for {
		if ip >= len(instructions) {
			break
		}

		instr := instructions[ip]
		continuation, err := instr.Perform(ctx, rt)
		if err != nil {
			ewj, ok := err.(*EndWithJump)
			if !ok {
				return err
			}

			if ewj.LabelsExited > 0 {
				ewj.LabelsExited--
				return ewj
			}
		}

		ip++

		if continuation != nil {
			err = rt.Stack.PushLabel(continuation)
			instructions = continuation.Instr
			ip = 0
		}

		err = checkDeadline(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
Exiting instr* with label L

When the end of a block is reached without a jump or trap aborting it, then the following steps are performed.
1.  Let m be the number of values on the top of the stack.
2.  Pop the valuesval𝑚from the stack.
3.  Assert: due tovalidation, the label L is now on the top of the stack.
4.  Pop the label from the stack.
5.  Push val^m back to the stack.
6.  Jump to the position after the end of the structured control instruction associated with the label L.
*/
func (rt *Runtime) exitInstructionsWithLabel() error {
	// 1.  Let m be the number of values on the top of the stack.
	m := rt.Stack.CountValuesOnTop()

	// 2.  Pop the values val^m from the stack.
	vals, err := rt.Stack.PopValues(m)
	if err != nil {
		return err
	}

	// 3.  Assert: due tovalidation, the label L is now on the top of the stack.
	if rt.Stack.Top() == nil || rt.Stack.Top().Label == nil {
		return errors.Errorf("label must be on top of the stack:\n%s", rt.Stack.Dump())
	}

	// 4.  Pop the label from the stack.
	_, err = rt.Stack.PopLabel()
	if err != nil {
		return err
	}

	// 5.  Push val^m back to the stack.
	err = rt.Stack.PushValuesBack(vals)
	if err != nil {
		return err
	}

	// 6.  Jump to the position after the end of the structured control instruction associated with the label L.
	return nil
}

func checkDeadline(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
