package cpu

import (
	"fmt"

	"github.com/kanataxa/famigone/pkg/bus"
)

type CPU struct {
	register *Register
	bus      *bus.Bus
}

type Executor interface {
	Exec() error
}

func (c *CPU) Pos() uint16 {
	return c.register.PC
}

func (c *CPU) Exec() error {
	for c.HasNext() {
		fmt.Println(Lookup(c.Current()))
		c.operate(Lookup(c.Current()))

		fmt.Printf("%x\n", c.bus.ReadROM(c.Pos()))
	}
	return nil
}

func (c *CPU) operate(op *Operator) {
	switch op.order {
	case lda:
		c.LDA()
	case ldx:
		c.LDX()
	case ldy:

	case sta:
		c.STA()
	case stx:
	case sty:
	case tax:
		c.TAX()
	case tay:
		c.TAY()
	case tsx:
		c.TSX()
	case txa:
		c.TXA()
	case txs:
		c.TXS()
	case tya:
		c.TYA()
	case adc:
	case and:
	case asl:
	case bit:
	case cmp:
	case cpx:
	case cpy:
	case dec:
	case dex:
	case dey:
	case eor:
	case inc:
	case inx:
	case iny:
	case lsr:
	case ora:
	case rol:
	case ror:
	case sbc:
	case pha:
	case php:
	case pla:
	case plp:
	case jmp:
	case jsr:
	//c.operate(Lookup(c.source[]))
	case rts:
	case rti:
	case bcc:
	case bcs:
	case beq:
	case bmi:
	case bne:
	case bpl:
	case bvc:
	case bvs:
	case clc:
		c.CLC()
	case cld:
		c.CLD()
	case cli:
		c.CLI()
	case clv:
		c.CLV()
	case sec:
		c.SEC()
	case sed:
		c.SED()
	case sei:
		c.SEI()
	case brk:
	case nop:

	}
	c.Next()
}

func (c *CPU) Current() byte {
	return c.bus.ReadROM(c.register.PC)
}

func (c *CPU) Next() byte {
	c.register.PC++
	return c.bus.ReadROM(c.register.PC)
}

func (c *CPU) HasNext() bool {
	return c.bus.ROMSize() > int(c.Pos())+1
}

func (c *CPU) LDA() {
	val := c.addressingValue()
	fmt.Println("value", val)
	c.register.A = uint8(val)
}

func (c *CPU) LDX() {
	val := c.addressingValue()
	fmt.Println("value", val)
	c.register.X = uint8(val)
}

func (c *CPU) STA() {
	val := c.addressingValue()
	fmt.Println("value", val)
	c.bus.Write(val, c.register.A)
}

func (c *CPU) TAX() {
	c.register.X = c.register.A
}

func (c *CPU) TAY() {
	c.register.Y = c.register.A
}

func (c *CPU) TSX() {
	c.register.X = c.register.S
}

func (c *CPU) TXA() {
	c.register.A = c.register.X
}

func (c *CPU) TXS() {
	c.register.S = c.register.X
}

func (c *CPU) TYA() {
	c.register.A = c.register.Y
}

func (c *CPU) JSR() {
}

func (c *CPU) CLC() {
	c.register.P.C = false
}

func (c *CPU) CLD() {
	// not implements in NES
}

func (c *CPU) CLI() {
	c.register.P.I = false
}

func (c *CPU) CLV() {
	c.register.P.V = false
}

func (c *CPU) SEC() {
	c.register.P.C = true
}

func (c *CPU) SED() {
	// not implements in NES
}

func (c *CPU) SEI() {
	c.register.P.I = true
}

func (c *CPU) addressingValue() uint16 {
	op := Lookup(c.Current())
	switch op.addressing {
	case implied:
	case accumulator:
		return uint16(c.register.A)
	case immediate:
		return uint16(c.Next())
	case zeropage:
		return uint16(c.Next())
	case zeropageX:
		return uint16(c.Next())
	case zeropageY:
		return uint16(c.bus.ReadROM(uint16(c.Next()) + uint16(c.register.Y)))
	case relative:
		return uint16(c.bus.ReadROM(uint16(c.Next()) + c.Pos()))
	case absolute:
		lower := c.Next()
		upper := c.Next()
		addr := c.convertValue(upper, lower)
		return uint16(c.bus.ReadROM(uint16(c.bus.ReadROM(addr))))
	case absoluteX:
		lower := c.Next()
		upper := c.Next()
		addr := c.convertValue(upper, lower) + uint16(c.register.X)
		return uint16(c.bus.ReadROM(uint16(c.bus.ReadROM(addr))))
	case absoluteY:
		lower := c.Next()
		upper := c.Next()
		addr := c.convertValue(upper, lower) + uint16(c.register.Y)
		return uint16(c.bus.ReadROM(uint16(c.bus.ReadROM(addr))))
	case indirect:
	case indirectX:
	case indirectY:
	}

	return 0
}

func (c *CPU) convertValue(upper, lower byte) uint16 {
	return uint16(upper)<<8 | uint16(lower)
}

func New(bus *bus.Bus) Executor {
	return &CPU{
		register: NewRegister(bus),
		bus:      bus,
	}
}
