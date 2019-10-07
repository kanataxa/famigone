package cpu

import "github.com/kanataxa/famigone/pkg/bus"

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

func NewRegister(bus *bus.Bus) *Register {
	return &Register{
		PC: uint16(bus.Read(uint16(bus.ROMSize()-3)))<<8 | uint16(bus.Read(uint16(bus.ROMSize()-4))),
		P:  &StatusRegister{},
	}
}
