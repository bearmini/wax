package wax

import "io"

/*
Expressions¶
https://webassembly.github.io/multi-value/core/syntax/instructions.html#syntax-expr

Function bodies, initialization values for globals, and offsets of element or data segments are given as expressions,
which are sequences of instructions terminated by an 'end' marker.

expr ::= instr* end

In some places, validation restricts expressions to be constant, which limits the set of allowable instructions.


Expressions
http://webassembly.github.io/spec/core/binary/instructions.html#expressions

Expressions are encoded by their instruction sequence terminated with an explicit 0x0B opcode for 'end'

expr ::= (in:instr)* 0x0B => in* end

*/
type Expr []Instr

func ParseExpr(ber *BinaryEncodingReader) (*Expr, error) {
	result := make(Expr, 0)

	for {
		instr, err := ParseInstr(ber)
		if err != nil {
			if err == io.EOF {
				return &result, nil
			}
			return nil, err
		}

		result = append(result, instr)
		if instr.Opcode() == OpcodeEnd {
			return &result, nil
		}
	}
}
