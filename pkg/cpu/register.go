package cpu

import (
	"github.com/kanataxa/famigone/pkg/bus"
)

type Register struct {
	A    uint8
	X    uint8
	Y    uint8
	S    uint8
	P    *StatusRegister
	PC   uint16
	head uint16
}

type StatusRegister struct {
	N bool
	V bool
	R bool
	B bool
	D bool
	I bool
	Z bool
	C bool
}

func NewRegister(bus *bus.Bus) *Register {
	rom := bus.ROM()
	head := uint16(rom.Read(uint16(rom.Size()-3)))<<8 | uint16(rom.Read(uint16(rom.Size()-4)))
	return &Register{
		head: head,
		PC:   head,
		P:    &StatusRegister{},
	}
}

func (r *Register) branch(addr uint16) {
	r.PC = addr
}

func (r *Register) jump(addr uint16) {
	r.PC = r.head + addr
}
