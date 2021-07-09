package memory

type RAM struct {
	data []byte
}

func NewRAM() *RAM {
	return &RAM{
		data: make([]byte, 2*1024),
	}
}

func (r *RAM) Read(addr uint16) byte {
	return r.data[addr]
}

func (r *RAM) Write(addr uint16, val byte) {
	r.data[addr] = val
}
