package wax

/*
	global    ::= gt:globaltype e:expr         => {type gt, init e}
*/
type Global struct {
	Type GlobalType
	Init Expr
}

func ParseGlobal(ber *BinaryEncodingReader) (*Global, error) {
	gt, err := ParseGlobalType(ber)
	if err != nil {
		return nil, err
	}
	e, err := ParseExpr(ber)
	if err != nil {
		return nil, err
	}
	return &Global{
		Type: *gt,
		Init: *e,
	}, nil
}
