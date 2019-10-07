package wax

func Disassemble(code Code) (string, error) {
	lines := make(disasmTree, 0)
	for _, instr := range code.Code.Body {
		d, err := instr.Disassemble()
		if err != nil {
			return "", err
		}
		lines = append(lines, d)
	}

	return lines.Join("\n"), nil
}
