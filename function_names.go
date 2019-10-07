package wax

type FunctionNames NameMap

func (fn *FunctionNames) GetNameType() NameType {
	return NameTypeFunction
}

func (fn *FunctionNames) FindByName(name string) *Naming {
	return (*NameMap)(fn).FindByName(name)
}

func ParseFunctionNames(ber *BinaryEncodingReader) (*FunctionNames, error) {
	nm, err := ParseNameMap(ber)
	if err != nil {
		return nil, err
	}

	fn := FunctionNames(*nm)
	return &fn, nil
}
