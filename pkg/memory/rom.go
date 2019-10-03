package memory

type ROM struct {
	source []byte
}

func NewROM(source []byte) *ROM {
	return &ROM{
		source: source,
	}
}

func (r *ROM) Read(addr uint16) byte {
	return r.source[addr]
}

func (r *ROM) Size() int {
	return len(r.source)
}
