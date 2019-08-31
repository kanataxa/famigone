package bus

import (
	"github.com/kanataxa/famigone/pkg/cassettie"
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

func (b *Bus) ReadWRAM(addr uint16) byte {
	return b.wram.Read(addr)
}

func (b *Bus) Write(addr uint16, val byte) {
	b.wram.Write(addr, val)
}

func New(c *cassettie.Cassettie) *Bus {
	return &Bus{
		rom:  c.ProgramROM(),
		wram: memory.RAM{},
	}
}
