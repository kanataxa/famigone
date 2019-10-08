package ppu

import "github.com/kanataxa/famigone/pkg/memory"

type PPU struct {
	ram memory.RAM
}

func (p *PPU) Write(addr, val uint16) {

}

func (p *PPU) Read(addr uint16) byte { return 0 }

func (p *PPU) Run() error {
	return nil
}
