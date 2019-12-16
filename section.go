package wax

import (
	"io"

	"github.com/pkg/errors"
)

/*
http://webassembly.github.io/spec/core/binary/modules.html#sections

Sections

Each section consists of

    - a one-byte section id,
    - the u32 size of the contents, in bytes,
    - the actual contents, whose structure is depended on the section id.

Every section is optional; an omitted section is equivalent to the section being present with empty contents.

The following parameterized grammar rule defines the generic structure of a section with id N
and contents described by the grammar 𝙱.

sectionN(B) ::= N:byte size:u32 cont:B => cont (if size = ||B||)
							| ε                      => ε

For most sections, the contents B encodes a vector. In these cases, the empty result ε is interpreted as the empty vector.

Note
Other than for unknown custom sections, the size is not required for decoding, but can be used to skip sections
when navigating through a binary. The module is malformed if the size does not match the length of the binary contents B.
*/
type Section interface {
	GetID() SectionID
	Encode(w *BinaryEncodingWriter) error
}

func ParseSections(ber *BinaryEncodingReader) ([]Section, error) {
	result := make([]Section, 0)

	for {
		idu8, err := ber.ReadU8()
		if err == io.EOF {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
		id := SectionID(idu8)

		var s Section
		switch id {
		case CustomSectionID:
			s, err = ParseCustomSection(ber, id)
		case TypeSectionID:
			s, err = ParseTypeSection(ber, id)
		case ImportSectionID:
			s, err = ParseImportSection(ber, id)
		case FunctionSectionID:
			s, err = ParseFunctionSection(ber, id)
		case TableSectionID:
			s, err = ParseTableSection(ber, id)
		case MemorySectionID:
			s, err = ParseMemorySection(ber, id)
		case GlobalSectionID:
			s, err = ParseGlobalSection(ber, id)
		case ExportSectionID:
			s, err = ParseExportSection(ber, id)
		case StartSectionID:
			s, err = ParseStartSection(ber, id)
		case ElementSectionID:
			s, err = ParseElementSection(ber, id)
		case CodeSectionID:
			s, err = ParseCodeSection(ber, id)
		case DataSectionID:
			s, err = ParseDataSection(ber, id)
		default:
			return nil, errors.Errorf("unknown section id: %#02x (%d)", id, id)
		}
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
}
