package wax

/*
resizable_limits

A packed tuple that describes the limits of a table or memory:
| Field    | Type        | Description
| flags    | varuint1    | 1 if the maximum field is present, 0 otherwise
| initial  | varuint32   | initial length (in units of table elements or wasm pages)
| maximum  | varuint32?  | only present if specified by flags

Note: In the future :unicorn:, the “flags” field may be changed to varuint32, e.g., to include a flag for sharing between threads.
*/
type ResizableLimits struct {
	Flags   uint32
	Initial uint32
	Maximum uint32
}

func ParseResizableLimits(ber *BinaryEncodingReader) (*ResizableLimits, error) {
	f, _, err := ber.ReadVaruintN(1)
	if err != nil {
		return nil, err
	}

	i, _, err := ber.ReadVaruintN(32)
	if err != nil {
		return nil, err
	}

	var m uint32
	if f == 1 {
		m64, _, err := ber.ReadVaruintN(32)
		if err != nil {
			return nil, err
		}
		m = uint32(m64)
	}

	return &ResizableLimits{
		Flags:   uint32(f),
		Initial: uint32(i),
		Maximum: m,
	}, nil
}
