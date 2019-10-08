package bus

import (
	"fmt"

	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/memory"
	"github.com/kanataxa/famigone/pkg/ppu"
)

type CPUBus struct {
	*cassette.Cassette
	wram memory.RAM
	ppu  *ppu.PPU
	// apu  memory.RAM
}

// See: https://wiki.nesdev.com/w/index.php/Mirroring
// TODO: implements mapper
func (b *CPUBus) Read(addr uint16) byte {
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
		return b.ProgramROM().Read(addr - 0x8000)
	} else {
		// smaller than 16kb, start 0xC000
		if b.ProgramROM().Size() <= 0x4000 {
			return b.ProgramROM().Read(addr - 0xC000)
		}
		return b.ProgramROM().Read(addr - 0x8000)
	}
}

func (b *CPUBus) Write(addr uint16, val byte) {
	fmt.Printf("WRITE **** %04x %d\n", addr, val)

	if addr < 0x0800 {
		b.wram.Write(addr, val)
	} else if addr < 0x2000 {
		// mirror address
		b.wram.Write(addr-0x0800, val)
	} else if addr < 0x4000 {
		// mirror address
		b.ppu.Write((addr-0x2000)%8, val)
	} else if addr < 0x4020 {
		// apu and i/o register
		// TODO: implements
		return
	} else if addr < 0x6000 {
		// extends ROM
		return
	} else if addr < 0x8000 {
		// extends RAM
		return
	} else if addr < 0xC000 {
		return
	} else {
		// smaller than 16kb, start 0xC000
		if b.ProgramROM().Size() <= 0x4000 {
			return
		}
		return
	}
}

func NewCPUBus(c *cassette.Cassette) Bus {
	return &CPUBus{
		Cassette: c,
		wram:     memory.RAM{},
	}
}
