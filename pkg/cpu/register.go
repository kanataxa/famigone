package cpu

import (
	"github.com/kanataxa/famigone/pkg/bus"
)

type Register struct {
	A  uint8
	X  uint8
	Y  uint8
	S  uint8
	P  *StatusRegister
	PC uint16
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

func NewRegister(bus bus.Bus) *Register {
	head := uint16(bus.Read(0xFFFD))<<8 | uint16(bus.Read(uint16(0xFFFC)))
	return &Register{
		PC: head,
		P:  &StatusRegister{},
	}
}

func (r *Register) branch(addr uint16) {
	r.jump(addr)
}

func (r *Register) jump(addr uint16) {
	//fmt.Printf("JMP:  [%04x]\n", addr)
	r.PC = addr
}
