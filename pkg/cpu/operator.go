package cpu

import "fmt"

type addressingType int

func (a addressingType) String() string {
	switch a {
	case implied:
		return "Implied"
	case accumulator:
		return "Accumulator"
	case immediate:
		return "Immediate"
	case zeropage:
		return "Zeropage"
	case zeropageX:
		return "Zeropage, X"
	case zeropageY:
		return "Zeropage, Y"
	case relative:
		return "Relative"
	case absolute:
		return "Absolute"
	case absoluteX:
		return "Absolute, X"
	case absoluteY:
		return "Absolute, Y"
	case indirect:
		return "(Indirect)"
	case indirectX:
		return "(Indirect, X)"
	case indirectY:
		return "(Indirect, Y)"
	}

	return "Unknown"
}

type orderType int

func (o orderType) String() string {
	switch o {
	case lda:
		return "LDA"
	case ldx:
		return "LDX"
	case ldy:
		return "LDY"
	case sta:
		return "STA"
	case stx:
		return "STX"
	case sty:
		return "STY"
	case tax:
		return "TAX"
	case tay:
		return "TAY"
	case tsx:
		return "TSX"
	case txa:
		return "TXA"
	case txs:
		return "TXS"
	case tya:
		return "TYA"
	case adc:
		return "ADC"
	case and:
		return "AND"
	case asl:
		return "ASL"
	case bit:
		return "BIT"
	case cmp:
		return "CMP"
	case cpx:
		return "CPX"
	case cpy:
		return "CPY"
	case dec:
		return "DEC"
	case dex:
		return "DEX"
	case dey:
		return "DEY"
	case eor:
		return "EOR"
	case inc:
		return "INC"
	case inx:
		return "INX"
	case iny:
		return "INY"
	case lsr:
		return "LSR"
	case ora:
		return "ORA"
	case rol:
		return "ROL"
	case ror:
		return "ROR"
	case sbc:
		return "SBC"
	case pha:
		return "PHA"
	case php:
		return "PHP"
	case pla:
		return "PLA"
	case plp:
		return "PLP"
	case jmp:
		return "JMP"
	case jsr:
		return "JSR"
	case rts:
		return "RTS"
	case rti:
		return "RTI"
	case bcc:
		return "BCC"
	case bcs:
		return "BCS"
	case beq:
		return "BEQ"
	case bmi:
		return "BMI"
	case bne:
		return "BNE"
	case bpl:
		return "BPL"
	case bvc:
		return "BVC"
	case bvs:
		return "BVS"
	case clc:
		return "CLC"
	case cld:
		return "CLD"
	case cli:
		return "CLI"
	case clv:
		return "CLV"
	case sec:
		return "SEC"
	case sed:
		return "SED"
	case sei:
		return "SEI"
	case brk:
		return "BRK"
	case nop:
		return "NOP"
	}
	return "Unknown"
}

const (
	implied addressingType = iota
	accumulator
	immediate
	zeropage
	zeropageX
	zeropageY
	relative
	absolute
	absoluteX
	absoluteY
	indirect
	indirectX
	indirectY
)

const (
	lda orderType = iota
	ldx
	ldy
	sta
	stx
	sty
	tax
	tay
	tsx
	txa
	txs
	tya
	adc
	and
	asl
	bit
	cmp
	cpx
	cpy
	dec
	dex
	dey
	eor
	inc
	inx
	iny
	lsr
	ora
	rol
	ror
	sbc
	pha
	php
	pla
	plp
	jmp
	jsr
	rts
	rti
	bcc
	bcs
	beq
	bmi
	bne
	bpl
	bvc
	bvs
	clc
	cld
	cli
	clv
	sec
	sed
	sei
	brk
	nop
)

var operators map[int8]*Operator

type Operator struct {
	code       int8
	order      orderType
	addressing addressingType
}

func (o *Operator) String() string {
	return fmt.Sprintf("[%s] [%s]", o.order, o.addressing)
}
func register(operator *Operator) {
	operators[operator.code] = operator
}

func init() {
	operators = make(map[int8]*Operator)
	register(&Operator{code: 0xA9, order: lda, addressing: immediate})
	register(&Operator{code: 0xA5, order: lda, addressing: zeropage})
	register(&Operator{code: 0xB5, order: lda, addressing: zeropageX})
	register(&Operator{code: 0xAD, order: lda, addressing: absolute})
	register(&Operator{code: 0xBD, order: lda, addressing: absoluteX})
	register(&Operator{code: 0xB9, order: lda, addressing: absoluteY})
	register(&Operator{code: 0xA1, order: lda, addressing: indirectX})
	register(&Operator{code: 0xB1, order: lda, addressing: indirectY})
}
