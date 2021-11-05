package cpu

import (
	"fmt"

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

func (r *StatusRegister) String() string {
	return fmt.Sprintf("%02x", r.Value())
}

func (r *StatusRegister) Value() uint8 {
	var bit uint8
	for idx, f := range []bool{
		r.C, r.Z, r.I, r.D, r.B, r.R, r.V, r.N,
	} {
		if f {
			bit |= 1 << uint8(idx)
		}
	}
	return bit
}

func (r *StatusRegister) Load(val uint8) {
	r.C = (val>>0)&1 == 1
	r.Z = (val>>1)&1 == 1
	r.I = (val>>2)&1 == 1
	r.D = (val>>3)&1 == 1
	r.B = (val>>4)&1 == 1
	//r.R = (val>>5)&1 == 1
	r.V = (val>>6)&1 == 1
	r.N = (val>>7)&1 == 1
}

func NewRegister(bus bus.Bus) *Register {
	head := uint16(bus.Read(0xFFFD))<<8 | uint16(bus.Read(uint16(0xFFFC)))
	return &Register{
		PC: head,
		P:  &StatusRegister{R: true, I: true},
		S:  0xFD,
	}
}

func (r *Register) branch(addr uint16) {
	r.jump(addr)
}

func (r *Register) jump(addr uint16) {
	r.PC = addr
}

func (r *Register) SetZ(val uint16) {
	r.P.Z = val == 0
}

func (r *Register) SetN(val uint16) {
	r.P.N = (val & 0x80) != 0
}
