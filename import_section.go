package wax

import (
	"bytes"
)

/*
Import Section
https://webassembly.github.io/multi-value/core/binary/modules.html#import-section

The import section has the id 2.
It decodes into a vector of imports that represent the imports component of a module.

importsec  ::= im*:section_2(vec(import))     => im*
import     ::= mod:name nm:name d:importdesc  => {module mod, name nm, desc d}
importdesc ::= 0x00 x:typeidx                 => func x
						 | 0x01 tt:tabletype              => table tt
						 | 0x02 mt:memtype                => mem mt
						 | 0x03 gt:globaltype             => global gt
*/
type ImportSection struct {
	SectionBase
	Imports []*Import
}

func ParseImportSection(ber *BinaryEncodingReader, id SectionID) (*ImportSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	br := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	// Read Count
	count64, _, err := br.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}
	count := uint32(count64)

	imports := make([]*Import, 0, count)
	for i := uint32(0); i < count; i++ {
		im, err := ParseImport(br)
		if err != nil {
			return nil, err
		}
		imports = append(imports, im)
	}

	return &ImportSection{
		SectionBase: *sb,
		Imports:     imports,
	}, nil
}

func (is *ImportSection) GetFuncImportsCount() uint32 {
	num := uint32(0)
	if is == nil {
		return 0
	}
	for _, im := range is.Imports {
		if im.DescType == ImportDescTypeFunc {
			num++
		}
	}
	return num
}
