package wax

import "bytes"

/*
Code Section
http://webassembly.github.io/spec/core/binary/modules.html#code-section

The code section has the id 10.
It decodes into a vector of code entries that are pairs of value type vectors and expressions.
They represent the locals and body field of the functions in the funcs component of a module.
The type fields of the respective functions are encoded separately in the function section.

The encoding of each code entry consists of

  - the u32 size of the function code in bytes,
  - the actual function code, which in turn consists of
      - the declaration of locals,
      - the function body as an expression.

Local declarations are compressed into a vector whose entries consist of

  - a u32 count,
  - a value type,

denoting count locals of the same value type.

	codesec ::= code*: section10(vec(code)) => code*
	code    ::= size:u32 code:func          => code               (if size = ||func||)
	func    ::= (t*)*:vec(locals) e:expr    => concat((t*)*),e*   (if |concat((t*)*)| < 2^32)
  locals  ::= n:u32 t:valtype             => tn


Here, code ranges over pairs (valtype*,expr).
The meta function concat((t*)*) concatenates all sequences ti* in (t*)*.

Any code for which the length of the resulting sequence is out of bounds of the maximum size of a vector is malformed.

Note

Like with sections, the code size is not needed for decoding, but can be used to skip functions when navigating through a binary. The module is malformed if a size does not match the length of the respective function code.
*/
type CodeSection struct {
	SectionBase
	Code []Code
}

func ParseCodeSection(ber *BinaryEncodingReader, id SectionID) (*CodeSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	// Read count of vector
	count64, _, err := cr.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	count := uint32(count64)

	codes := make([]Code, 0, count)
	for i := uint32(0); i < count; i++ {
		c, err := ParseCode(cr)
		if err != nil {
			return nil, err
		}
		codes = append(codes, *c)
	}

	return &CodeSection{
		SectionBase: *sb,
		Code:        codes,
	}, nil
}
