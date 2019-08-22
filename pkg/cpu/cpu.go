package cpu

import "fmt"

type CPU struct {
	register *Register
	source   []byte
}

type Executor interface {
	Exec() error
}

func (c *CPU) Pos() int {
	return int(c.register.PC)
}

func (c *CPU) Exec() error {
	fmt.Println("len:::", len(c.source))
	for c.Next() {
		fmt.Println(c.Current())
		c.operate(c.Current())

		fmt.Printf("%x\n", c.source[c.Pos()])
	}
	return nil
}

func (c *CPU) operate(op *Operator) {
	switch op.order {
	case lda:
		c.LDA()
		c.Next()
	case ldx:
		c.LDX()
		c.Next()
	case ldy:

	case sta:
		c.STA()
		c.Next()
	case stx:
	case sty:
	case tax:
	case tay:
	case tsx:
	case txa:
	case txs:
		c.TXS()
	case tya:
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
	case cld:
	case cli:
	case clv:
	case sec:
	case sed:
	case sei:
		c.SEI()
	case brk:
	case nop:

	}
}

func (c *CPU) Current() *Operator {
	return Lookup(c.source[c.register.PC])
}

func (c *CPU) Next() bool {
	c.register.PC++
	return len(c.source) > c.Pos()
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
	c.source[val] = c.register.A
}

func (c *CPU) TXS() {
	c.register.S = c.register.X
}

func (c *CPU) JSR() {
}

func (c *CPU) SEI() {
	c.register.P.I = true
}

func (c *CPU) addressingValue() uint16 {
	op := c.Current()
	switch op.addressing {
	case implied:
	case accumulator:
		return uint16(c.register.A)
	case immediate:
		return uint16(c.source[c.Pos()+1])
	case zeropage:
		return uint16(c.source[c.source[c.Pos()+1]])
	case zeropageX:
		return uint16(c.source[uint8(c.source[c.Pos()+1])+c.register.X])
	case zeropageY:
		return uint16(c.source[uint8(c.source[c.Pos()+1])+c.register.Y])
	case relative:
		return uint16(c.source[int(c.source[c.Pos()+1])+c.Pos()])
	case absolute:
		addr := c.convertValue(c.source[c.Pos()+2], c.source[c.Pos()+1])
		return uint16(c.source[c.source[addr]])
	case absoluteX:
		addr := c.convertValue(c.source[c.Pos()+2], c.source[c.Pos()+1]) + uint16(c.register.X)
		return uint16(c.source[c.source[addr]])
	case absoluteY:
		addr := c.convertValue(c.source[c.Pos()+2], c.source[c.Pos()+1]) + uint16(c.register.Y)
		return uint16(c.source[c.source[addr]])
	case indirect:
	case indirectX:
	case indirectY:
	}

	return 0
}

func (c *CPU) convertValue(upper, lower byte) uint16 {
	return uint16(upper)<<8 | uint16(lower)
}

func New(source []byte) Executor {
	return &CPU{
		register: &Register{
			PC: -1,
			P:  &StatusRegister{},
		},
		source: source,
	}
}
