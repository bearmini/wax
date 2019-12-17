package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bearmini/sexp"
	"github.com/bearmini/wax"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

var opts struct {
	WastDir  string `short:"d" long:"wast-dir" description:"directory for wast files"`
	WastPath string `short:"i" long:"wast-path" description:"path for a wast file"`
	WasmPath string `short:"w" long:"wasm-path" default:"wasm" description:"path to wasm command"`
}

var (
	nModule            int
	currentRuntime     *wax.Runtime
	runtimesWithID     map[string]*wax.Runtime = make(map[string]*wax.Runtime)
	registeredRuntimes map[string]*wax.Runtime = make(map[string]*wax.Runtime)
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("ERROR: %+v\n", err)
	}
}

func run() error {
	_, err := flags.Parse(&opts)
	if err != nil {
		return err
	}

	if opts.WastPath != "" {
		return runTest(opts.WastPath)
	}

	if opts.WastDir != "" {
		return runTests(opts.WastDir)
	}

	return errors.New("either --wast-dir or --wast-path is required")
}

func runTests(dir string) error {
	paths, err := filepath.Glob(filepath.Join(dir, "*.wast"))
	for _, path := range paths {
		err = runTest(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func runTest(path string) error {
	nModule = 0
	progress(path)

	f, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "unable to open file %s", path)
	}
	defer f.Close()

	wr := NewWastReader(f)

	for {
		s, err := wr.NextSexp()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if s == nil {
			break
		}

		err = processSexp(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func processSexp(s *sexp.Sexp) error {
	if s.Children == nil || len(s.Children) == 0 {
		return errors.New("unsupported sexp")
	}

	first := s.Children[0]
	if first.Atom == nil || first.Atom.Type != sexp.TokenTypeSymbol {
		return errors.New("unsupported sexp")
	}

	switch first.Atom.Value {
	case "module":
		nModule++
		fmt.Printf("compiling module: %d\n", nModule)
		m, err := compile(s.String())
		if err != nil {
			return err
		}

		cfg := wax.NewRuntimeConfig().MaxMemorySizeInPage(1024)
		err = prepareImports(m, cfg)
		if err != nil {
			return err
		}

		rt, err := wax.NewRuntime(m, cfg)
		if err != nil {
			return err
		}

		rememberRuntimeIfNeeded(s, rt)
		currentRuntime = rt

	case "register":
		if len(s.Children) < 3 {
			return errors.New("unsupported register")
		}

		if s.Children[1].Atom == nil || s.Children[2].Atom == nil {
			return errors.New("unsupported register")
		}
		second := s.Children[1].Atom
		third := s.Children[2].Atom

		if second.Type != sexp.TokenTypeString || third.Type != sexp.TokenTypeSymbol {
			return errors.New("unsupported register")
		}

		moduleName := strings.Trim(second.Value, `"`)
		moduleID := third.Value

		rt, ok := runtimesWithID[moduleID]
		if !ok {
			return errors.Errorf("module not found: %s", moduleID)
		}
		registeredRuntimes[moduleName] = rt

	case "invoke":
		fmt.Printf("-> %s\n", s.String())
		if len(s.Children) < 2 {
			return errors.Errorf("insufficient arguments: %s", s.String())
		}

		_, err := eval(s)
		if err != nil {
			return err
		}

	case "assert_return":
		fmt.Printf("-> %s\n", s.String())
		if len(s.Children) < 2 {
			return errors.Errorf("insufficient arguments: %s", s.String())
		}
		invokeOrGet := s.Children[1]

		var expectedVal *wax.Val
		if len(s.Children) > 2 {
			expected := s.Children[2]
			ev, err := evalConst(expected)
			if err != nil {
				return err
			}
			expectedVal = ev
		}

		actual, err := eval(invokeOrGet)
		if err != nil {
			return err
		}
		if expectedVal != nil {
			if len(actual) != 1 {
				return errors.Errorf("expected 1 return value, but got %d\n%+v\n", len(actual), actual)
			}
			if !expectedVal.EqualsTo(actual[0]) {
				return errors.Errorf("assertion failed: %s\nExpected: %+v\nActual:   %+v\n", s.String(), expectedVal, actual[0])
			}
		}

	case "assert_trap":
		fmt.Printf("skipping assert_trap\n")
	case "assert_malformed":
		fmt.Printf("skipping assert_malformed\n")
	case "assert_invalid":
		fmt.Printf("skipping assert_invalid\n")
	case "assert_return_canonical_nan":
		fmt.Printf("skipping assert_return_canonical_nan\n")
	case "assert_return_arithmetic_nan":
		fmt.Printf("skipping assert_return_arithmetic_nan\n")
	case "assert_exhaustion":
		fmt.Printf("skipping assert_exhaustion\n")
	case "assert_unlinkable":
		fmt.Printf("skipping assert_unlinkable\n")
	default:
		return errors.Errorf("not implemented: first atom value: %s", first.Atom.Value)
	}

	return nil
}

func rememberRuntimeIfNeeded(s *sexp.Sexp, rt *wax.Runtime) {
	if !doesModuleSexpHaveModuleID(s) {
		return
	}

	moduleID := s.Children[1].Atom.Value

	runtimesWithID[moduleID] = rt
}

func doesModuleSexpHaveModuleID(s *sexp.Sexp) bool {
	if len(s.Children) < 2 {
		return false
	}

	if s.Children[1].Atom == nil {
		return false
	}

	if s.Children[1].Atom.Type != sexp.TokenTypeSymbol {
		return false
	}

	return true
}

func prepareImports(m *wax.Module, cfg *wax.RuntimeConfig) error {
	imports := m.GetImports()
	for _, im := range imports {
		switch im.DescType {
		case wax.ImportDescTypeFunc:
			cfg.AddImportFunc(im.Mod, im.Nm)
		case wax.ImportDescTypeTable:
			e, max, err := provideTable(im.Mod, im.Nm)
			if err != nil {
				return err
			}
			cfg.AddImportTable(im.Mod, im.Nm, e, max)
		case wax.ImportDescTypeMem:
			b, max, err := provideMemory(im.Mod, im.Nm)
			if err != nil {
				return err
			}
			cfg.AddImportMemory(im.Mod, im.Nm, b, max)
		case wax.ImportDescTypeGlobal:
			v, m, err := provideGlobal(im.Mod, im.Nm)
			if err != nil {
				return err
			}
			cfg.AddImportGlobal(im.Mod, im.Nm, *v, m)
		}
	}
	return nil
}

func provideTable(module, name wax.Name) ([]wax.FuncElem, *uint32, error) {
	switch module {
	case "spectest":
		switch name {
		case "table":
			min := 10
			return make([]wax.FuncElem, min), nil, nil
		case "global_i32":
			min := 10
			return make([]wax.FuncElem, min), nil, nil
		default:
			return nil, nil, errors.New("unknown table name")
		}
	default:
		rt, ok := registeredRuntimes[string(module)]
		if !ok {
			return nil, nil, errors.Errorf("unknown table module: %s", module)
		}

		ti, err := rt.GetExportedTable(name)
		if err != nil {
			return nil, nil, err
		}
		return ti.Elem, ti.Max, nil
	}
}

func provideMemory(module, name wax.Name) ([]byte, *uint32, error) {
	switch module {
	case "spectest":
		switch name {
		case "memory":
			return make([]byte, wax.PageSize), nil, nil
		case "global_i32":
			return make([]byte, wax.PageSize), nil, nil
		default:
			return nil, nil, errors.New("unknown memory name")
		}
	default:
		return nil, nil, errors.New("unknown memory module")
	}
}

func provideGlobal(module, name wax.Name) (*wax.Val, wax.Mut, error) {
	switch module {
	case "spectest":
		switch name {
		case "global_i32":
			return wax.NewValI32(0), wax.MutConst, nil
		default:
			return nil, wax.MutConst, errors.New("unknown memory name")
		}
	default:
		return nil, wax.MutConst, errors.New("unknown memory module")
	}
}

func eval(s *sexp.Sexp) ([]*wax.Val, error) {
	if s.Children == nil || len(s.Children) == 0 {
		return nil, errors.New("unexpected expression")
	}

	first := s.Children[0]
	if first.Atom == nil || first.Atom.Type != sexp.TokenTypeSymbol {
		return nil, errors.New("unexpected sexp")
	}

	switch first.Atom.Value {
	case "invoke":
		if len(s.Children) < 2 {
			return nil, errors.Errorf("insufficient arguments: %s", s.String())
		}
		second := s.Children[1]
		if second.Atom == nil {
			return nil, errors.New("unexpected invoke format")
		}

		// we have 2 formats:
		//  (invoke "funcName" (arg1) (arg2)...)
		//  (invoke moduleID "funcName" (arg1) (arg2)...)
		moduleID := ""
		funcName := ""
		args := []*sexp.Sexp{}
		if second.Atom.Type == sexp.TokenTypeString {
			moduleID = ""
			funcName = strings.Trim(second.Atom.Value, `"`)
			args = append(args, s.Children[2:]...)
		}
		if second.Atom.Type == sexp.TokenTypeSymbol && len(s.Children) >= 3 {
			third := s.Children[2]
			if third.Atom == nil || third.Atom.Type != sexp.TokenTypeString {
				return nil, errors.New("unexpected invoke format")
			}
			moduleID = second.Atom.Value
			funcName = strings.Trim(third.Atom.Value, `"`)
			args = append(args, s.Children[3:]...)
		}

		if moduleID == "" && funcName == "" {
			return nil, errors.New("unexpected invoke format")
		}

		rt := currentRuntime
		if moduleID != "" {
			rtWithID, ok := runtimesWithID[moduleID]
			if !ok {
				return nil, errors.Errorf("module not found: %s", moduleID)
			}
			rt = rtWithID
		}

		fa, err := rt.FindFuncAddr(funcName)
		if err != nil {
			return nil, err
		}

		argVals := []*wax.Val{}
		for _, arg := range args {
			v, err := evalConst(arg)
			if err != nil {
				return nil, err
			}
			argVals = append(argVals, v)
		}
		ctx := context.Background()
		return rt.InvokeFunc(ctx, *fa, argVals)

	case "get":
		if len(s.Children) < 2 {
			return nil, errors.Errorf("insufficient arguments: %s", s.String())
		}
		second := s.Children[1]
		if second.Atom == nil {
			return nil, errors.New("unexpected get format")
		}

		// we have 2 formats:
		//  (get "globalName")
		//  (get moduleID "globalName")
		moduleID := ""
		globalName := ""
		if second.Atom.Type == sexp.TokenTypeString {
			globalName = strings.Trim(second.Atom.Value, `"`)
		}
		if second.Atom.Type == sexp.TokenTypeSymbol {
			third := s.Children[2]
			if third.Atom == nil || third.Atom.Type != sexp.TokenTypeString {
				return nil, errors.New("unexpected invoke format")
			}
			moduleID = second.Atom.Value
			globalName = strings.Trim(third.Atom.Value, `"`)
		}

		if moduleID == "" && globalName == "" {
			return nil, errors.New("unexpected get format")
		}

		rt := currentRuntime
		if moduleID != "" {
			rtWithID, ok := runtimesWithID[moduleID]
			if !ok {
				return nil, errors.Errorf("module not found: %s", moduleID)
			}
			rt = rtWithID
		}
		v, err := rt.GetExportedGlobal(wax.Name(globalName))
		if err != nil {
			return nil, err
		}

		return []*wax.Val{&v.Value}, nil

	default:
		return nil, errors.New("not implemented")
	}
}

func reverse(vals []*wax.Val) {
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
}

func evalConst(s *sexp.Sexp) (*wax.Val, error) {
	if s.Children == nil || len(s.Children) < 2 {
		return nil, errors.Errorf("invalid const expression: %s", s.String())
	}

	first := s.Children[0]
	if first.Atom == nil || first.Atom.Type != sexp.TokenTypeSymbol {
		return nil, errors.Errorf("invalid const expression: %s", s.String())
	}

	switch first.Atom.Value {
	case "i32.const":
		second := s.Children[1]
		if second.Atom == nil || second.Atom.Type != sexp.TokenTypeNumber {
			return nil, errors.Errorf("invalid const expression: %s", s.String())
		}
		v, err := parseI32(second.Atom.Value)
		if err != nil {
			return nil, err
		}
		return wax.NewValI32(v), nil
	case "i64.const":
		second := s.Children[1]
		if second.Atom == nil || second.Atom.Type != sexp.TokenTypeNumber {
			return nil, errors.Errorf("invalid const expression: %s", s.String())
		}
		v, err := parseI64(second.Atom.Value)
		if err != nil {
			return nil, err
		}
		return wax.NewValI64(v), nil
	case "f32.const":
		second := s.Children[1]
		if second.Atom == nil {
			return nil, errors.Errorf("invalid const expression: %s", s.String())
		}

		sat := second.Atom.Type
		sav := second.Atom.Value
		switch sat {
		case sexp.TokenTypeNumber:
			if strings.HasPrefix(sav, "-nan") {
				nan, err := wax.ParseNaN32WithPayload(sav)
				if err != nil {
					return nil, err
				}
				return wax.NewValF32(nan), nil
			}
			if strings.HasPrefix(sav, "-inf") {
				return wax.NewValF32(float32(math.Inf(-1))), nil
			}
			buf := bytes.NewBuffer([]byte{})
			bew := wax.NewBinaryEncodingWriter(buf)
			err := bew.WriteU8(uint8(wax.OpcodeF32Const))
			if err != nil {
				return nil, err
			}
			v, err := strconv.ParseFloat(sav, 32)
			if err != nil {
				if _, ok := err.(*strconv.NumError); !ok {
					return nil, err
				}
				v32, err := wax.ParseHexfloat32(sav)
				if err != nil {
					return nil, err
				}
				v = float64(v32)
			}
			err = bew.WriteU32LE(math.Float32bits(float32(v)))
			if err != nil {
				return nil, err
			}
			val := wax.Val(buf.Bytes())
			return &val, nil

		case sexp.TokenTypeSymbol:
			if strings.HasPrefix(sav, "nan") {
				nan, err := wax.ParseNaN32WithPayload(sav)
				if err != nil {
					return nil, err
				}
				return wax.NewValF32(nan), nil
			}
			if strings.HasPrefix(sav, "inf") {
				return wax.NewValF32(float32(math.Inf(1))), nil
			}
			return nil, errors.Errorf("invalid const expression: %s", s.String())

		default:
			return nil, errors.Errorf("invalid const expression: %s", s.String())
		}

	case "f64.const":
		second := s.Children[1]
		if second.Atom == nil {
			return nil, errors.Errorf("invalid const expression: %s", s.String())
		}

		sat := second.Atom.Type
		sav := second.Atom.Value
		switch sat {
		case sexp.TokenTypeNumber:
			if strings.HasPrefix(sav, "-nan") {
				nan, err := wax.ParseNaN64WithPayload(sav)
				if err != nil {
					return nil, err
				}
				return wax.NewValF64(nan), nil
			}
			if strings.HasPrefix(sav, "-inf") {
				return wax.NewValF64(math.Inf(-1)), nil
			}
			buf := bytes.NewBuffer([]byte{})
			bew := wax.NewBinaryEncodingWriter(buf)
			err := bew.WriteU8(uint8(wax.OpcodeF64Const))
			if err != nil {
				return nil, err
			}
			v, err := strconv.ParseFloat(sav, 64)
			if err != nil {
				if _, ok := err.(*strconv.NumError); !ok {
					return nil, err
				}
				v, err = wax.ParseHexfloat64(sav)
				if err != nil {
					return nil, err
				}
			}
			err = bew.WriteU64LE(math.Float64bits(float64(v)))
			if err != nil {
				return nil, err
			}
			val := wax.Val(buf.Bytes())
			return &val, nil
		case sexp.TokenTypeSymbol:
			if strings.HasPrefix(sav, "nan") {
				nan, err := wax.ParseNaN64WithPayload(sav)
				if err != nil {
					return nil, err
				}
				return wax.NewValF64(nan), nil
			}
			if strings.HasPrefix(sav, "inf") {
				return wax.NewValF64(math.Inf(1)), nil
			}
			return nil, errors.Errorf("invalid const expression: %s", s.String())

		default:
			return nil, errors.Errorf("invalid const expression: %s", s.String())
		}
	}
	return nil, errors.New("not implemented")
}

func parseI32(s string) (uint32, error) {
	u, err := strconv.ParseUint(s, 0, 32)
	if err == nil {
		return uint32(u), nil
	}

	ne, ok := err.(*strconv.NumError)
	if !ok {
		return 0, err
	}
	if ne.Err != strconv.ErrSyntax { // ParseUint() fails for negative numbers such as "-123" with ErrSyntax. So we will try ParseInt() if the error is ErrSyntax
		return 0, err
	}

	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}

	return uint32(i), nil
}

func parseI64(s string) (uint64, error) {
	u, err := strconv.ParseUint(s, 0, 64)
	if err == nil {
		return uint64(u), nil
	}

	ne, ok := err.(*strconv.NumError)
	if !ok {
		return 0, err
	}
	if ne.Err != strconv.ErrSyntax { // ParseUint() fails for negative numbers such as "-123" with ErrSyntax. So we will try ParseInt() if the error is ErrSyntax
		return 0, err
	}

	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}

	return uint64(i), nil
}

func compile(s string) (*wax.Module, error) {
	srcPath, err := writeToTempFile(s, "wast")
	if err != nil {
		return nil, err
	}
	defer os.Remove(srcPath)

	dstPath, err := writeToTempFile("", "wasm")
	if err != nil {
		return nil, err
	}
	defer os.Remove(dstPath)

	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})
	cmd := exec.Command(opts.WasmPath, "-d", srcPath, "-o", dstPath)
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf(stdout.String())
		fmt.Fprintf(os.Stderr, stderr.String())
		return nil, err
	}

	bin, err := os.Open(dstPath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open temporary destination file: %s", dstPath)
	}
	defer bin.Close()

	return wax.ParseBinaryModule(bin)
}

func writeToTempFile(s, ext string) (string, error) {
	f, err := ioutil.TempFile("", fmt.Sprintf("webassembly-standard-test-*.%s", ext))
	if err != nil {
		return "", err
	}
	fname := f.Name()

	_, err = f.Write([]byte(s))
	if err != nil {
		os.Remove(fname)
		return "", err
	}

	err = f.Close()
	if err != nil {
		os.Remove(fname)
		return "", err
	}

	return fname, nil
}

func progress(path string) {
	fmt.Printf("processing ... %s\n", path)
}
