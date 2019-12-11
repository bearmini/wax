package wax

import (
	"github.com/pkg/errors"
)

type SectionBase struct {
	ID      SectionID
	Size    uint32
	Content []byte
}

func (s SectionBase) GetID() SectionID {
	return s.ID
}

func (s SectionBase) Encode(bew *BinaryEncodingWriter) error {
	err := bew.WriteVaruint(uint64(s.ID))
	if err != nil {
		return err
	}

	err = bew.WriteVaruint(uint64(s.Size))
	if err != nil {
		return err
	}

	n, err := bew.Write([]byte(s.Content))
	if err != nil {
		return err
	}
	if uint32(n) != s.Size {
		return errors.New("could not write all data")
	}

	return nil
}

func ParseSectionBase(ber *BinaryEncodingReader, id SectionID) (*SectionBase, error) {
	size64, _, err := ber.ReadVaruint()
	if err != nil {
		return nil, err
	}
	size := uint32(size64)

	// Read Payload
	content := make([]byte, size)
	n, err := ber.Read(content)
	if err != nil {
		return nil, err
	}
	if uint32(n) != size {
		return nil, errors.New("insufficient data")
	}

	return &SectionBase{
		ID:      id,
		Size:    size,
		Content: content,
	}, nil
}
