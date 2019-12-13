package wax

/*
2.3.7 Global Types
Global types classify global variables, which hold a value and can either be mutable or immutable.

globaltype ::= mut valtype
mut ::= const | var


5.3.7 Global Types
Global types are encoded by their value type and a flag for their mutability.

globaltype ::= t:valtype m:mut ⇒ m t
mut ::= 0x00 ⇒ const
      | 0x01 ⇒ var
*/
type GlobalType struct {
	Mut     Mut
	ValType ValType
}

func ParseGlobalType(ber *BinaryEncodingReader) (*GlobalType, error) {
	vt, _, err := ParseValType(ber)
	if err != nil {
		return nil, err
	}
	m, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}

	return &GlobalType{
		Mut:     Mut(m),
		ValType: *vt,
	}, nil
}
