package bus

import (
	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/memory"
)

type Bus struct {
	rom  *memory.ROM
	wram memory.RAM
	// ppu
}

func (b *Bus) ReadROM(addr uint16) byte {
	return b.rom.Read(addr)
}

func (b *Bus) ROMSize() int {
	return b.rom.Size()
}

func (b *Bus) ReadWRAM(addr uint16) byte {
	return b.wram.Read(addr)
}

func (b *Bus) Write(addr uint16, val byte) {
	b.wram.Write(addr, val)
}

func New(c *cassette.Cassette) *Bus {
	return &Bus{
		rom:  c.ProgramROM(),
		wram: memory.RAM{},
	}
}
