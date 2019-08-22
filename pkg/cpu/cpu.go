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

/*

7	N	ネガティブ	Aの7ビット目と同じになります。負数の判定用。
6	V	オーバーフロー	演算がオーバーフローを起こした場合セットされます。
5	R	予約済み	使用できません。常にセットされています。
4	B	ブレークモード	BRK発生時はセットされ、IRQ発生時はクリアされます。
3	D	デシマルモード	セットすると、BCDモードで動作します。(ファミコンでは未実装)
2	I	IRQ禁止	クリアするとIRQが許可され、セットするとIRQが禁止になります。
1	Z	ゼロ	演算結果が0になった場合セットされます。ロード命令でも変化します。
0	C	キャリー	キャリー発生時セットされます。
*/

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
