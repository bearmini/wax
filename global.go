package wax

/*
	global    ::= gt:globaltype e:expr         => {type gt, init e}
*/
type Global struct {
	Type GlobalType
	Init InitExpr
}

func ParseGlobal(ber *BinaryEncodingReader) (*Global, error) {
	t, err := ParseGlobalType(ber)
	if err != nil {
		return nil, err
	}
	i, err := ParseInitExpr(ber)
	if err != nil {
		return nil, err
	}
	return &Global{
		Type: *t,
		Init: i,
	}, nil
}
