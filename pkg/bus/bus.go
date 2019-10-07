package bus

import (
	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/memory"
)

type Bus struct {
	rom  *memory.ROM
	wram memory.RAM
	ppu  memory.RAM
	apu  memory.RAM
}

// See: https://wiki.nesdev.com/w/index.php/Mirroring
func (b *Bus) Read(addr uint16) byte {
	if addr < 0x0800 {
		return b.wram.Read(addr)
	} else if addr < 0x2000 {
		// mirror address
		return b.wram.Read(addr - 0x0800)
	} else if addr < 0x4000 {
		// mirror address
		return b.ppu.Read((addr - 0x2000) % 8)
	} else if addr < 0x4020 {
		// apu and i/o register
		// TODO: implements
		return 0
	} else if addr < 0x6000 {
		// extends ROM
		return 0
	} else if addr < 0x8000 {
		// extends RAM
		return 0
	} else if addr < 0xC000 {
		return b.rom.Read(addr - 0x8000)
	} else {
		// smaller than 16kb, start 0xC000
		if b.rom.Size() <= 0x4000 {
			return b.rom.Read(addr - 0xC000)
		}
		return b.rom.Read(addr - 0x8000)
	}
}

func (b *Bus) ROM() *memory.ROM {
	return b.rom
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
