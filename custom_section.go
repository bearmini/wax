package wax

import (
	"bytes"

	"github.com/pkg/errors"
)

/*
http://webassembly.github.io/spec/core/binary/modules.html#custom-section

Custom Section

Custom sections have the id 0. They are intended to be used for debugging information or third-party extensions,
and are ignored by the WebAssembly semantics. Their contents consist of a name further identifying the custom section,
followed by an uninterpreted sequence of bytes for custom use.

customsec ::= section0(custom)
custom    ::= name byte*

Note
If an implementation interprets the contents of a custom section, then errors in that contents, or the placement of the section,
must not invalidate the module.
*/
type CustomSection struct {
	SectionBase
	Name Name
}

func (s *CustomSection) Encode(w *BinaryEncodingWriter) error {
	originalContent := s.Content

	// re-generate content
	bb := bytes.NewBuffer([]byte{})
	bew := NewBinaryEncodingWriter(bb)

	err := bew.WriteVaruint(uint64(len(s.Name)))
	if err != nil {
		return err
	}

	n, err := bew.Write([]byte(s.Name))
	if err != nil {
		return err
	}
	if n != len(s.Name) {
		return errors.New("could not write all data")
	}

	n, err = bew.Write(originalContent)
	if err != nil {
		return err
	}
	if n != len(originalContent) {
		return errors.New("could not write all data")
	}

	s.Content = bb.Bytes()
	s.Size = uint32(len(s.Content))

	return s.SectionBase.Encode(w)
}

func ParseCustomSection(ber *BinaryEncodingReader, id SectionID) (*CustomSection, error) {
	sb, err := ParseSectionBase(ber, id)
	if err != nil {
		return nil, err
	}

	cr := NewBinaryEncodingReader(bytes.NewReader(sb.Content))

	name, c, err := ParseName(cr)
	if err != nil {
		return nil, err
	}

	dataLen := sb.Size - uint32(len(c))

	data := make([]byte, dataLen)
	n, err := cr.Read(data)
	if err != nil {
		return nil, err
	}
	if uint32(n) != dataLen {
		return nil, errors.New("insufficient data")
	}

	sb.Size = dataLen
	sb.Content = data

	return &CustomSection{
		SectionBase: *sb,
		Name:        *name,
	}, nil
}
