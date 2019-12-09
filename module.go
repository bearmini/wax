package wax

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Module struct {
	valid    bool
	Preamble Preamble
	Sections []Section
}

func (m *Module) IsValid() bool {
	return m.valid
}

func (m *Module) GetTypeSection() *TypeSection {
	for _, s := range m.Sections {
		if s.GetID() != TypeSectionID {
			continue
		}

		ts, ok := s.(*TypeSection)
		if !ok {
			return nil
		}

		return ts
	}
	return nil
}

func (m *Module) GetImportSection() *ImportSection {
	for _, s := range m.Sections {
		if s.GetID() != ImportSectionID {
			continue
		}

		is, ok := s.(*ImportSection)
		if !ok {
			return nil
		}

		return is
	}
	return nil
}

func (m *Module) GetFunctionSection() *FunctionSection {
	for _, s := range m.Sections {
		if s.GetID() != FunctionSectionID {
			continue
		}

		fs, ok := s.(*FunctionSection)
		if !ok {
			return nil
		}

		return fs
	}
	return nil
}

func (m *Module) GetTableSection() *TableSection {
	for _, s := range m.Sections {
		if s.GetID() != TableSectionID {
			continue
		}

		ts, ok := s.(*TableSection)
		if !ok {
			return nil
		}

		return ts
	}
	return nil
}

func (m *Module) GetMemorySection() *MemorySection {
	for _, s := range m.Sections {
		if s.GetID() != MemorySectionID {
			continue
		}

		ms, ok := s.(*MemorySection)
		if !ok {
			return nil
		}

		return ms
	}
	return nil
}

func (m *Module) GetGlobalSection() *GlobalSection {
	for _, s := range m.Sections {
		if s.GetID() != GlobalSectionID {
			continue
		}

		gs, ok := s.(*GlobalSection)
		if !ok {
			return nil
		}

		return gs
	}
	return nil
}

func (m *Module) GetExportSection() *ExportSection {
	for _, s := range m.Sections {
		if s.GetID() != ExportSectionID {
			continue
		}

		es, ok := s.(*ExportSection)
		if !ok {
			return nil
		}

		return es
	}
	return nil
}

func (m *Module) GetStartSection() *StartSection {
	for _, s := range m.Sections {
		if s.GetID() != StartSectionID {
			continue
		}

		ss, ok := s.(*StartSection)
		if !ok {
			return nil
		}

		return ss
	}
	return nil
}

func (m *Module) GetElementSection() *ElementSection {
	for _, s := range m.Sections {
		if s.GetID() != ElementSectionID {
			continue
		}

		es, ok := s.(*ElementSection)
		if !ok {
			return nil
		}

		return es
	}
	return nil
}

func (m *Module) GetCodeSection() *CodeSection {
	for _, s := range m.Sections {
		if s.GetID() != CodeSectionID {
			continue
		}

		cs, ok := s.(*CodeSection)
		if !ok {
			return nil
		}

		return cs
	}
	return nil
}

func (m *Module) GetDataSection() *DataSection {
	for _, s := range m.Sections {
		if s.GetID() != DataSectionID {
			continue
		}

		ds, ok := s.(*DataSection)
		if !ok {
			return nil
		}

		return ds
	}
	return nil
}

func (m *Module) FindCustomSectionWithName(name string) *CustomSection {
	for _, s := range m.Sections {
		if s.GetID() != CustomSectionID {
			continue
		}

		cs, ok := s.(*CustomSection)
		if !ok {
			return nil
		}

		if string(cs.Name) != name {
			continue
		}

		return cs
	}

	return nil
}

func (m *Module) GetImports() []*Import {
	is := m.GetImportSection()
	if is == nil {
		return []*Import{}
	}

	return is.Imports
}

func (m *Module) GetFuncs() []Func {
	fs := m.GetFunctionSection()
	cs := m.GetCodeSection()

	result := make([]Func, 0, len(cs.Code))
	for i, c := range cs.Code {
		result = append(result, Func{
			Type:   fs.Types[i],
			Locals: c.Code.Locals,
			Body:   c.Code.Body,
		})
	}

	return result
}

func (m *Module) GetExportedFuncName(fa FuncAddr) Name {
	is := m.GetImportSection()
	nfim := is.GetFuncImportsCount()
	es := m.GetExportSection()

	for _, e := range es.Exports {
		if e.DescType != ExportDescTypeFunc {
			continue
		}

		fi := e.Desc.(*FuncIdx)
		if uint32(*fi)-nfim == uint32(fa) {
			return e.Nm
		}
	}

	return Name("(unknown)")
}

func (m *Module) GetMems() []*MemType {
	ms := m.GetMemorySection()
	if ms == nil {
		return []*MemType{}
	}
	return ms.MemTypes
}

func (m *Module) ToJSON() string {
	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("%+v", err)
	}
	return string(b)
}

func (m *Module) EncodeToBinaryModule(w io.Writer) error {
	err := m.Preamble.Encode(w)
	if err != nil {
		return err
	}

	for _, s := range m.Sections {
		err = s.Encode(NewBinaryEncodingWriter(w))
		if err != nil {
			return err
		}
	}

	return nil
}

func ParseBinaryModule(r io.Reader) (*Module, error) {
	ber := NewBinaryEncodingReader(r)
	p, err := ParsePreamble(ber)
	if err != nil {
		return nil, err
	}

	s, err := ParseSections(ber)
	if err != nil {
		return nil, err
	}

	err = validateSections(s)
	if err != nil {
		return nil, err
	}

	return &Module{
		valid:    true,
		Preamble: *p,
		Sections: s,
	}, nil
}

func validateSections(sections []Section) error {
	// Each known section is optional and may appear at most once. Custom sections all have the same id (0), and can be named non-uniquely (all bytes composing their names may be identical).

	m := make(map[SectionID]bool)
	for _, s := range sections {
		if s.GetID() == CustomSectionID {
			continue
		}
		if m[s.GetID()] {
			return errors.Errorf("duplicated sections. section id: %d", s.GetID())
		}
		m[s.GetID()] = true
	}
	return nil
}
