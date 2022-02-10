package cpu

import (
	"fmt"

	"github.com/kanataxa/famigone/pkg/bus"
)

type CPU struct {
	register *Register
	bus      bus.Bus
	stack    uint16
}

func (c *CPU) Pos() uint16 {
	return c.register.PC
}

func (c *CPU) Run() error {
	//	fmt.Println("----RUN----")
	//fmt.Printf("%04x  %s\n", c.register.PC, Lookup(c.Current()))
	c.operate(Lookup(c.Current()))

	//	fmt.Printf("TEST RESULT: [%x]\n", c.bus.Read(0xc002))
	//	fmt.Printf("TEST RESULT: [%x]\n", c.bus.Read(0xc003))

	return nil
}

func (c *CPU) operate(op *Operator) {
	switch op.order {
	case lda:
		c.LDA()
	case ldx:
		c.LDX()
	case ldy:
		c.LDY()
	case sta:
		c.STA()
	case stx:
		c.STX()
	case sty:
		c.STY()
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
		c.ADC()
	case and:
		c.AND()
	case asl:
		c.ASL()
	case bit:
		c.BIT()
	case cmp:
		c.CMP()
	case cpx:
		c.CPX()
	case cpy:
		c.CPY()
	case dec:
		c.DEC()
	case dex:
		c.DEX()
	case dey:
		c.DEY()
	case eor:
		c.EOR()
	case inc:
		c.INC()
	case inx:
		c.INX()
	case iny:
		c.INY()
	case lsr:
		c.LSR()
	case ora:
		c.ORA()
	case rol:
		c.ROL()
	case ror:
		c.ROR()
	case sbc:
		c.SBC()
	case pha:
		c.PHA()
	case php:
		c.PHP()
	case pla:
		c.PLA()
	case plp:
		c.PLP()
	case jmp:
		c.JMP()
		// don't call Next()
		return
	case jsr:
		c.JSR()
		// don't call Next()
		return
	case rts:
		c.RTS()
	case rti:
		c.RTI()
		// don't call Next()
		return
	case bcc:
		if c.BCC() {
			return
		}
	case bcs:
		if c.BCS() {
			return
		}
	case beq:
		if c.BEQ() {
			return
		}
	case bmi:
		if c.BMI() {
			return
		}
	case bne:
		if c.BNE() {
			return
		}
	case bpl:
		if c.BPL() {
			return
		}
	case bvc:
		if c.BVC() {
			return
		}
	case bvs:
		if c.BVS() {
			return
		}
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
		if c.BRK() {
			return
		}
	case nop:
		c.NOP()
	}
	c.Next()
}

func (c *CPU) Current() byte {
	return c.bus.Read(c.register.PC)
}

func (c *CPU) Next() byte {
	c.register.PC++
	return c.bus.Read(c.register.PC)
}

func (c *CPU) LDA() {
	pos := c.register.PC
	op := Lookup(c.Current())
	address := c.addressingValue()
	val := uint16(c.bus.Read(address))
	fmt.Printf("LDA %04x %s %04x %04x\n", pos, op.addressing, address, val)
	c.register.A = uint8(val)
	c.register.SetZ(val)
	c.register.SetN(val)
}

func (c *CPU) LDX() {
	val := uint16(c.bus.Read(c.addressingValue()))
	c.register.X = uint8(val)
	c.register.SetZ(val)
	c.register.SetN(val)
}

func (c *CPU) LDY() {
	val := uint16(c.bus.Read(c.addressingValue()))
	c.register.Y = uint8(val)
	c.register.SetZ(val)
	c.register.SetN(val)
}

func (c *CPU) STA() {
	val := c.addressingValue()
	fmt.Printf("STA %04x %04x \n", val, c.register.A)
	c.bus.Write(val, c.register.A)
}

func (c *CPU) STX() {
	val := c.addressingValue()
	c.bus.Write(val, c.register.X)
}

func (c *CPU) STY() {
	val := c.addressingValue()
	c.bus.Write(val, c.register.Y)
}

func (c *CPU) TAX() {
	c.register.X = c.register.A
	c.register.SetN(uint16(c.register.X))
	c.register.SetZ(uint16(c.register.X))
}

func (c *CPU) TAY() {
	c.register.Y = c.register.A
	c.register.SetN(uint16(c.register.Y))
	c.register.SetZ(uint16(c.register.Y))
}

func (c *CPU) TSX() {
	c.register.X = c.register.S
	c.register.SetN(uint16(c.register.X))
	c.register.SetZ(uint16(c.register.X))
}

func (c *CPU) TXA() {
	c.register.A = c.register.X
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))
}

func (c *CPU) TXS() {
	c.register.S = c.register.X
}

func (c *CPU) TYA() {
	c.register.A = c.register.Y
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))
}

func (c *CPU) ADC() {
	val := c.bus.Read(c.addressingValue())
	var rc uint8
	if c.register.P.C {
		rc = 1
	}
	a := c.register.A
	c.register.A = a + rc + val
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))
	c.register.P.C = uint16(a)+uint16(rc)+uint16(val) > 0xFF
	c.register.P.V = (a^val)&0x80 == 0 && (a^c.register.A)&0x80 != 0
}

func (c *CPU) AND() {
	val := c.bus.Read(c.addressingValue())
	c.register.A &= val
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))
}

func (c *CPU) ASL() {
	if Lookup(c.Current()).addressing == accumulator {
		c.register.P.C = c.register.A&0x80 == 0x80
		c.register.A <<= 1
		c.register.SetN(uint16(c.register.A))
		c.register.SetZ(uint16(c.register.A))
	} else {
		address := c.addressingValue()
		val := uint16(c.bus.Read(address))
		c.register.P.C = val&0x80 == 0x80
		val <<= 1
		c.bus.Write(address, uint8(val))
		c.register.SetN(val)
		c.register.SetZ(val)
	}
}

func (c *CPU) BIT() {
	val := c.bus.Read(c.addressingValue())
	a := c.register.A
	c.register.SetZ(uint16(a & val))
	c.register.SetN(uint16(val))
	c.register.P.V = (val & 0x40) != 0
}

func (c *CPU) CMP() {
	c.cp(c.register.A, c.bus.Read(c.addressingValue()))
}

func (c *CPU) CPX() {
	c.cp(c.register.X, c.bus.Read(c.addressingValue()))
}

func (c *CPU) CPY() {
	c.cp(c.register.Y, c.bus.Read(c.addressingValue()))
}

func (c *CPU) cp(a, b uint8) {
	result := uint16(a - b)
	c.register.SetN(result)
	c.register.SetZ(result)
	c.register.P.C = a >= b
}

func (c *CPU) EOR() {
	val := c.bus.Read(c.addressingValue())
	c.register.A ^= val
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))
}

func (c *CPU) DEC() {
	address := c.addressingValue()
	val := c.bus.Read(address) - 1
	c.bus.Write(address, val)
	c.register.SetN(uint16(val))
	c.register.SetZ(uint16(val))
}

func (c *CPU) DEX() {
	c.register.X--
	c.register.SetN(uint16(c.register.X))
	c.register.SetZ(uint16(c.register.X))
}

func (c *CPU) DEY() {
	c.register.Y--
	c.register.SetN(uint16(c.register.Y))
	c.register.SetZ(uint16(c.register.Y))
}

func (c *CPU) INC() {
	address := c.addressingValue()
	val := c.bus.Read(address) + 1
	c.bus.Write(address, val)
	c.register.SetN(uint16(val))
	c.register.SetZ(uint16(val))
}

func (c *CPU) INX() {
	c.register.X++
	c.register.SetN(uint16(c.register.X))
	c.register.SetZ(uint16(c.register.X))
}

func (c *CPU) INY() {
	c.register.Y++
	c.register.SetN(uint16(c.register.Y))
	c.register.SetZ(uint16(c.register.Y))
}

func (c *CPU) LSR() {
	if Lookup(c.Current()).addressing == accumulator {
		c.register.P.C = c.register.A&1 == 1
		c.register.A >>= 1
		c.register.SetN(uint16(c.register.A))
		c.register.SetZ(uint16(c.register.A))
	} else {
		address := c.addressingValue()
		val := uint16(c.bus.Read(address))
		c.register.P.C = val&1 == 1
		val >>= 1
		c.bus.Write(address, uint8(val))
		c.register.SetN(val)
		c.register.SetZ(val)
	}
}

func (c *CPU) ORA() {
	val := c.bus.Read(c.addressingValue())
	c.register.A |= val
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))
}

func (c *CPU) ROL() {
	var cv uint8
	if c.register.P.C {
		cv = 1
	}
	if Lookup(c.Current()).addressing == accumulator {
		c.register.P.C = c.register.A&0x80 == 0x80
		c.register.A = (c.register.A << 1) | cv
		c.register.SetN(uint16(c.register.A))
		c.register.SetZ(uint16(c.register.A))
	} else {
		address := c.addressingValue()
		val := c.bus.Read(address)
		c.register.P.C = val&0x80 == 0x80
		val = (val << 1) | cv
		c.bus.Write(address, val)
		c.register.SetN(uint16(val))
		c.register.SetZ(uint16(val))
	}
}

func (c *CPU) ROR() {
	var cv uint8
	if c.register.P.C {
		cv = 1
	}
	if Lookup(c.Current()).addressing == accumulator {
		c.register.P.C = c.register.A&1 == 1
		c.register.A = (c.register.A >> 1) | (cv << 7)
		c.register.SetN(uint16(c.register.A))
		c.register.SetZ(uint16(c.register.A))
	} else {
		address := c.addressingValue()
		val := c.bus.Read(address)
		c.register.P.C = val&1 == 1
		val = (val >> 1) | (cv << 7)
		c.bus.Write(address, val)
		c.register.SetN(uint16(val))
		c.register.SetZ(uint16(val))
	}
}

func (c *CPU) SBC() {
	val := c.bus.Read(c.addressingValue())
	a := c.register.A
	var carry uint8
	if !c.register.P.C {
		carry = 1
	}
	c.register.A = a - val - carry
	c.register.SetN(uint16(c.register.A))
	c.register.SetZ(uint16(c.register.A))

	// 繰り上がりなし(繰り下がり発生): C=1, 繰り上がり発生(繰り下がりなし): C=0
	c.register.P.C = int(a)-int(val)-int(carry) >= 0
	// オーバフローは同符号の時に発生する
	c.register.P.V = (a^val)&0x80 != 0 && (a^c.register.A)&0x80 != 0
}

func (c *CPU) PHA() {
	c.pushStack(c.register.A)
}

func (c *CPU) PHP() {
	c.pushStack(c.register.P.Value())
}

func (c *CPU) PLA() {
	c.register.A = c.popStack()
	c.register.SetZ(uint16(c.register.A))
	c.register.SetN(uint16(c.register.A))
}

func (c *CPU) PLP() {
	c.register.P.Load(c.popStack())
}

func (c *CPU) JMP() {
	v := c.addressingValue()
	c.register.jump(v)
}

func (c *CPU) JSR() {
	v := c.addressingValue()
	c.pushStack(uint8(c.register.PC >> 8))
	c.pushStack(uint8(c.register.PC))
	c.register.jump(v)
}

func (c *CPU) RTS() {
	lower := c.popStack()
	upper := c.popStack()
	c.register.jump(c.convertValue(upper, lower))
}

func (c *CPU) RTI() {
	c.register.P.Load(c.popStack())
	lower := c.popStack()
	upper := c.popStack()
	c.register.jump(c.convertValue(upper, lower))
}

func (c *CPU) BCC() bool {
	v := c.addressingValue()
	if c.register.P.C {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BCS() bool {
	v := c.addressingValue()
	if !c.register.P.C {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BEQ() bool {
	v := c.addressingValue()
	if !c.register.P.Z {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BMI() bool {
	v := c.addressingValue()
	if !c.register.P.N {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BNE() bool {
	v := c.addressingValue()
	if c.register.P.Z {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BPL() bool {
	v := c.addressingValue()
	if c.register.P.N {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BVC() bool {
	v := c.addressingValue()
	if c.register.P.V {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) BVS() bool {
	v := c.addressingValue()
	if !c.register.P.V {
		return false
	}
	c.register.branch(v)
	return true
}

func (c *CPU) CLC() {
	c.register.P.C = false
}

func (c *CPU) CLD() {
	c.register.P.D = false
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
	c.register.P.D = true
}

func (c *CPU) SEI() {
	c.register.P.I = true
}

func (c *CPU) BRK() bool {
	if c.register.P.I {
		return false
	}
	c.register.P.B = true
	c.Next()
	c.pushStack(uint8(c.register.PC) >> 4)
	c.pushStack(uint8(c.register.PC & 0x0F))
	c.pushStack(c.register.P.Value())
	c.register.P.I = true
	lower := c.bus.Read(0xFFFE)
	upper := c.bus.Read(0xFFFF)
	c.register.jump(c.convertValue(upper, lower))
	return true
}

func (c *CPU) NOP() {
}

func (c *CPU) addressingValue() uint16 {
	op := Lookup(c.Current())
	switch op.addressing {
	case implied:
	case accumulator:
		return uint16(c.register.A)
	case immediate:
		c.Next()
		//fmt.Printf("AD IMMD %04x %04x \n", c.register.PC, c.bus.Read(c.register.PC))
		return c.register.PC
	case zeropage:
		return uint16(c.Next())
	case zeropageX:
		lower := c.Next() + c.register.X
		return c.convertValue(0x00, lower)
	case zeropageY:
		lower := c.Next() + c.register.Y
		return c.convertValue(0x00, lower)
	case relative:
		offset := uint16(c.Next())
		if offset < 0x80 {
			return c.Pos() + offset + 1
		}
		return c.Pos() + offset + 1 - 0x100
	case absolute:
		lower := c.Next()
		upper := c.Next()
		return c.convertValue(upper, lower)
	case absoluteX:
		lower := c.Next()
		upper := c.Next()
		return c.convertValue(upper, lower) + uint16(c.register.X)
	case absoluteY:
		lower := c.Next()
		upper := c.Next()
		return c.convertValue(upper, lower) + uint16(c.register.Y)
	case indirect:
		lower := c.Next()
		upper := c.Next()
		mem := c.convertValue(upper, lower)
		return c.convertValue(c.bus.Read(c.nextIndirectAddr(mem)), c.bus.Read(mem))
	case indirectX:
		address := uint16(c.Next() + c.register.X)
		lower := c.bus.Read(address)
		upper := c.bus.Read(c.nextIndirectAddr(address))
		//fmt.Printf("AD INDIRECT %04x %04x %04x \n", address, upper, lower)
		return c.convertValue(upper, lower)
	case indirectY:
		address := uint16(c.Next())
		lower := c.bus.Read(address)
		upper := c.bus.Read(c.nextIndirectAddr(address))
		//fmt.Printf("AD INDIRECTY %04x %04x %04x %04x %04x\n", address, upper, lower, c.register.Y, c.convertValue(upper, lower)+uint16(c.register.Y))
		return c.convertValue(upper, lower) + uint16(c.register.Y)
	}
	return 0
}

// indirect時に上位はaddress+1となるが、0x02FFを0x0300ではなく0x0200に循環させる
func (c *CPU) nextIndirectAddr(addr uint16) uint16 {
	return (addr & 0xFF00) | uint16(uint8(addr)+1)
}

func (c *CPU) convertValue(upper, lower byte) uint16 {
	return uint16(upper)<<8 | uint16(lower)
}

func (c *CPU) pushStack(val byte) {
	c.bus.Write(0x0100|uint16(c.register.S), val)
	c.register.S--
}

func (c *CPU) popStack() byte {
	c.register.S++
	return c.bus.Read(0x0100 | uint16(c.register.S))
}

func New(bus bus.Bus) *CPU {
	return &CPU{
		register: NewRegister(bus),
		bus:      bus,
	}
}
