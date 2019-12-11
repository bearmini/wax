package wax

/*
/*
Global Types
https://webassembly.github.io/multi-value/core/syntax/types.html#syntax-mut

Global types classify global variables, which hold a value and can either be mutable or immutable.

globaltype ::= mut valtype
mut        ::= const | var


Global Types
http://webassembly.github.io/spec/core/binary/types.html

Global types are encoded by their value type and a flag for their mutability.

globaltype ::= r:valtype m:mut => m t
mut        ::= 0x00 => const
             | 0x01 => va
*/
type GlobalType struct {
	R ValType
	M Mut
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
		R: *vt,
		M: Mut(m),
	}, nil
}
