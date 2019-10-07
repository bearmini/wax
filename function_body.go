package wax

/*
Function bodies consist of a sequence of local variable declarations followed by bytecode instructions.
Instructions are encoded as an opcode followed by zero or more immediates as defined by the tables below.
Each function body must end with the end opcode.

| Field         | Type           | Description
| body_size     | varuint32      | size of function body to follow, in bytes
| local_count   | varuint32      | number of local entries
| locals        | local_entry*   | local variables
| code          | byte*          | bytecode of the function
| end           | byte           | 0x0b, indicating the end of the body
*/
/*
type FunctionBody struct {
	BodySize   uint32
	LocalCount uint32
	Locals     []*LocalEntry
	CodeBytes  CodeBytes
	//Code       Code
	End byte
}

func ParseFunctionBody(ber *BinaryEncodingReader) (*FunctionBody, error) {
	bs64, _, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	bs := uint32(bs64)

	bodyBytes := make([]byte, bs)
	n, err := ber.Read(bodyBytes)
	if err != nil {
		return nil, err
	}
	if uint32(n) != bs {
		return nil, errors.New("insufficient data")
	}

	consumedBytes := []byte{}
	br := NewBinaryEncodingReader(bytes.NewReader(bodyBytes))
	lc, c, err := br.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	consumedBytes = append(consumedBytes, c...)

	var locals []*LocalEntry
	if lc > 0 {
		locals = make([]*LocalEntry, 0, lc)
		for i := uint64(0); i < lc; i++ {
			l, c, err := ParseLocalEntry(br)
			if err != nil {
				return nil, err
			}
			locals = append(locals, l)
			consumedBytes = append(consumedBytes, c...)
		}
	}

	codeBytes := bodyBytes[len(consumedBytes):]
	end := codeBytes[len(codeBytes)-1:][0]
	code := codeBytes[0 : len(codeBytes)-1]

	return &FunctionBody{
		BodySize:   bs,
		LocalCount: uint32(lc),
		Locals:     locals,
		CodeBytes:  code,
		//Code:       code,
		End: end,
	}, nil
}
*/
