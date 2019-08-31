package memory

type RAM [2 * 1024]byte

func (r RAM) Read(addr uint16) byte {
	return r[addr]
}

func (r RAM) Write(addr uint16, val byte) {
	r[addr] = val
}
