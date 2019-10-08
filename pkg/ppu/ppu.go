package ppu

import (
	"fmt"

	"github.com/kanataxa/famigone/pkg/memory"
)

type PPU struct {
	ram memory.RAM
}

func (p *PPU) Write(addr uint16, val byte) {
	fmt.Println("PPU****", addr, val)

}

func (p *PPU) Read(addr uint16) byte { return 0 }

func (p *PPU) Run() error {
	return nil
}
