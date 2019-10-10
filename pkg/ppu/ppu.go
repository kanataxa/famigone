package ppu

import (
	"fmt"
	"time"

	"github.com/kanataxa/famigone/pkg/memory"
)

// PPUCTRL 8bit
// VPHB SINN
// |||| ||||
// |||| ||++- Base nametable address
// |||| ||    (0 = $2000; 1 = $2400; 2 = $2800; 3 = $2C00)
// |||| |+--- VRAM address increment per CPU read/write of PPUDATA
// |||| |     (0: add 1, going across; 1: add 32, going down)
// |||| +---- Sprite pattern table address for 8x8 sprites
// ||||       (0: $0000; 1: $1000; ignored in 8x16 mode)
// |||+------ Background pattern table address (0: $0000; 1: $1000)
// ||+------- Sprite size (0: 8x8 pixels; 1: 8x16 pixels)
// |+-------- PPU master/slave select
// |          (0: read backdrop from EXT pins; 1: output color on EXT pins)
// +--------- Generate an NMI at the start of the
type ctrl struct {
	nameTableAddress byte
	vramAddress      byte
	spriteAddress    byte
	bgAddress        byte
	spriteSize       byte
	masterSlave      byte
	nmi              byte
}

func (c *ctrl) write(val byte) {
	c.nameTableAddress = val & 3
	c.vramAddress = (val >> 2) & 1
	c.spriteAddress = (val >> 3) & 1
	c.bgAddress = (val >> 4) & 1
	c.spriteSize = (val >> 5) & 1
	c.masterSlave = (val >> 6) & 1
	c.nmi = (val >> 7) & 1
}

type PPU struct {
	ram memory.RAM

	*ctrl
	mask uint8 // BGRs bMmG
}

func (p *PPU) Write(addr uint16, val byte) {
	if addr == 0x2000 {
		p.ctrl.write(val)
	} else if addr == 0x2001 {
		p.mask = val
	} else if addr == 0x2002 {
		// read only
		fmt.Println("not write")
	} else if addr == 0x2003 {

	}
	fmt.Printf("PPU**** 0x%04x, 0x%04x 0b%08b \n", addr, val, val)
	time.Sleep(time.Second * 3)

}

func (p *PPU) Read(addr uint16) byte { return 0 }

func (p *PPU) Run() error {
	return nil
}

func New() *PPU {
	return &PPU{
		ctrl: &ctrl{},
	}
}
