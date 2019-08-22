package cpu

import (
	"fmt"
	"log"
)

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

var operators map[byte]*Operator

type Operator struct {
	code       uint8
	order      orderType
	addressing addressingType
}

func (o *Operator) String() string {
	return fmt.Sprintf("[%s] [%s]", o.order, o.addressing)
}

func Lookup(b byte) *Operator {
	return operators[b]
}

func register(operator *Operator) {
	if _, exists := operators[operator.code]; exists {
		log.Fatalf("duplicate code: %s", operator)
	}
	operators[operator.code] = operator
}

func init() {
	operators = make(map[byte]*Operator)
	register(&Operator{code: 0xA9, order: lda, addressing: immediate})
	register(&Operator{code: 0xA5, order: lda, addressing: zeropage})
	register(&Operator{code: 0xB5, order: lda, addressing: zeropageX})
	register(&Operator{code: 0xAD, order: lda, addressing: absolute})
	register(&Operator{code: 0xBD, order: lda, addressing: absoluteX})
	register(&Operator{code: 0xB9, order: lda, addressing: absoluteY})
	register(&Operator{code: 0xA1, order: lda, addressing: indirectX})
	register(&Operator{code: 0xB1, order: lda, addressing: indirectY})

	register(&Operator{code: 0xA2, order: ldx, addressing: immediate})
	register(&Operator{code: 0xA6, order: ldx, addressing: zeropage})
	register(&Operator{code: 0xB6, order: ldx, addressing: zeropageY})
	register(&Operator{code: 0xAE, order: ldx, addressing: absolute})
	register(&Operator{code: 0xBE, order: ldx, addressing: absoluteY})

	register(&Operator{code: 0xA0, order: ldy, addressing: immediate})
	register(&Operator{code: 0xA4, order: ldy, addressing: zeropage})
	register(&Operator{code: 0xB4, order: ldy, addressing: zeropageX})
	register(&Operator{code: 0xAC, order: ldy, addressing: absolute})
	register(&Operator{code: 0xBC, order: ldy, addressing: absoluteX})

	register(&Operator{code: 0x85, order: sta, addressing: zeropage})
	register(&Operator{code: 0x95, order: sta, addressing: zeropageX})
	register(&Operator{code: 0x8D, order: sta, addressing: absolute})
	register(&Operator{code: 0x9D, order: sta, addressing: absoluteX})
	register(&Operator{code: 0x99, order: sta, addressing: absoluteY})
	register(&Operator{code: 0x81, order: sta, addressing: indirectX})
	register(&Operator{code: 0x91, order: sta, addressing: indirectY})

	register(&Operator{code: 0x86, order: stx, addressing: zeropage})
	register(&Operator{code: 0x96, order: stx, addressing: zeropageY})
	register(&Operator{code: 0x8E, order: stx, addressing: absolute})

	register(&Operator{code: 0x84, order: sty, addressing: zeropage})
	register(&Operator{code: 0x94, order: sty, addressing: zeropageX})
	register(&Operator{code: 0x8C, order: sty, addressing: absolute})

	register(&Operator{code: 0xAA, order: tax, addressing: implied})

	register(&Operator{code: 0xA8, order: tay, addressing: implied})

	register(&Operator{code: 0xBA, order: tsx, addressing: implied})

	register(&Operator{code: 0x8A, order: txa, addressing: implied})

	register(&Operator{code: 0x9A, order: txs, addressing: implied})

	register(&Operator{code: 0x98, order: tya, addressing: implied})

	register(&Operator{code: 0x69, order: adc, addressing: immediate})
	register(&Operator{code: 0x65, order: adc, addressing: zeropage})
	register(&Operator{code: 0x75, order: adc, addressing: zeropageX})
	register(&Operator{code: 0x6D, order: adc, addressing: absolute})
	register(&Operator{code: 0x7D, order: adc, addressing: absoluteX})
	register(&Operator{code: 0x79, order: adc, addressing: absoluteY})
	register(&Operator{code: 0x61, order: adc, addressing: indirectX})
	register(&Operator{code: 0x71, order: adc, addressing: indirectY})

	register(&Operator{code: 0x29, order: and, addressing: immediate})
	register(&Operator{code: 0x25, order: and, addressing: zeropage})
	register(&Operator{code: 0x35, order: and, addressing: zeropageX})
	register(&Operator{code: 0x2D, order: and, addressing: absolute})
	register(&Operator{code: 0x3D, order: and, addressing: absoluteX})
	register(&Operator{code: 0x39, order: and, addressing: absoluteY})
	register(&Operator{code: 0x21, order: and, addressing: indirectX})
	register(&Operator{code: 0x31, order: and, addressing: indirectY})

	register(&Operator{code: 0x0A, order: asl, addressing: accumulator})
	register(&Operator{code: 0x06, order: asl, addressing: zeropage})
	register(&Operator{code: 0x16, order: asl, addressing: zeropageX})
	register(&Operator{code: 0x0E, order: asl, addressing: absolute})
	register(&Operator{code: 0x1E, order: asl, addressing: absoluteX})

	register(&Operator{code: 0x24, order: bit, addressing: zeropage})
	register(&Operator{code: 0x2C, order: bit, addressing: absolute})

	register(&Operator{code: 0xC9, order: cmp, addressing: immediate})
	register(&Operator{code: 0xC5, order: cmp, addressing: zeropage})
	register(&Operator{code: 0xD5, order: cmp, addressing: zeropageX})
	register(&Operator{code: 0xCD, order: cmp, addressing: absolute})
	register(&Operator{code: 0xDD, order: cmp, addressing: absoluteX})
	register(&Operator{code: 0xD9, order: cmp, addressing: absoluteY})
	register(&Operator{code: 0xC1, order: cmp, addressing: indirectX})
	register(&Operator{code: 0xD1, order: cmp, addressing: indirectY})

	register(&Operator{code: 0xE0, order: cpx, addressing: immediate})
	register(&Operator{code: 0xE4, order: cpx, addressing: zeropage})
	register(&Operator{code: 0xEC, order: cpx, addressing: absolute})

	register(&Operator{code: 0xC0, order: cpy, addressing: immediate})
	register(&Operator{code: 0xC4, order: cpy, addressing: zeropage})
	register(&Operator{code: 0xCC, order: cpy, addressing: absolute})

	register(&Operator{code: 0xC6, order: dec, addressing: zeropage})
	register(&Operator{code: 0xD6, order: dec, addressing: zeropageX})
	register(&Operator{code: 0xCE, order: dec, addressing: absolute})
	register(&Operator{code: 0xDE, order: dec, addressing: absoluteX})

	register(&Operator{code: 0xCA, order: dex, addressing: implied})
	register(&Operator{code: 0x88, order: dey, addressing: implied})

	register(&Operator{code: 0x49, order: eor, addressing: immediate})
	register(&Operator{code: 0x45, order: eor, addressing: zeropage})
	register(&Operator{code: 0x55, order: eor, addressing: zeropageX})
	register(&Operator{code: 0x4D, order: eor, addressing: absolute})
	register(&Operator{code: 0x5D, order: eor, addressing: absoluteX})
	register(&Operator{code: 0x59, order: eor, addressing: absoluteY})
	register(&Operator{code: 0x41, order: eor, addressing: indirectX})
	register(&Operator{code: 0x51, order: eor, addressing: indirectY})

	register(&Operator{code: 0xE6, order: inc, addressing: zeropage})
	register(&Operator{code: 0xF6, order: inc, addressing: zeropageX})
	register(&Operator{code: 0xEE, order: inc, addressing: absolute})
	register(&Operator{code: 0xFE, order: inc, addressing: absoluteX})

	register(&Operator{code: 0xE8, order: inx, addressing: implied})
	register(&Operator{code: 0xC8, order: iny, addressing: implied})

	register(&Operator{code: 0x4A, order: lsr, addressing: accumulator})
	register(&Operator{code: 0x46, order: lsr, addressing: zeropage})
	register(&Operator{code: 0x56, order: lsr, addressing: zeropageX})
	register(&Operator{code: 0x4E, order: lsr, addressing: absolute})
	register(&Operator{code: 0x5E, order: lsr, addressing: absoluteX})

	register(&Operator{code: 0x09, order: ora, addressing: immediate})
	register(&Operator{code: 0x05, order: ora, addressing: zeropage})
	register(&Operator{code: 0x15, order: ora, addressing: zeropageX})
	register(&Operator{code: 0x0D, order: ora, addressing: absolute})
	register(&Operator{code: 0x1D, order: ora, addressing: absoluteX})
	register(&Operator{code: 0x19, order: ora, addressing: absoluteY})
	register(&Operator{code: 0x01, order: ora, addressing: indirectX})
	register(&Operator{code: 0x11, order: ora, addressing: indirectY})

	register(&Operator{code: 0x2A, order: rol, addressing: accumulator})
	register(&Operator{code: 0x26, order: rol, addressing: zeropage})
	register(&Operator{code: 0x36, order: rol, addressing: zeropageX})
	register(&Operator{code: 0x2E, order: rol, addressing: absolute})
	register(&Operator{code: 0x3E, order: rol, addressing: absoluteX})

	register(&Operator{code: 0x6A, order: ror, addressing: accumulator})
	register(&Operator{code: 0x66, order: ror, addressing: zeropage})
	register(&Operator{code: 0x76, order: ror, addressing: zeropageX})
	register(&Operator{code: 0x6E, order: ror, addressing: absolute})
	register(&Operator{code: 0x7E, order: ror, addressing: absoluteX})

	register(&Operator{code: 0xE9, order: sbc, addressing: immediate})
	register(&Operator{code: 0xE5, order: sbc, addressing: zeropage})
	register(&Operator{code: 0xF5, order: sbc, addressing: zeropageX})
	register(&Operator{code: 0xED, order: sbc, addressing: absolute})
	register(&Operator{code: 0xFD, order: sbc, addressing: absoluteX})
	register(&Operator{code: 0xF9, order: sbc, addressing: absoluteY})
	register(&Operator{code: 0xE1, order: sbc, addressing: indirectX})
	register(&Operator{code: 0xF1, order: sbc, addressing: indirectY})

	register(&Operator{code: 0x48, order: pha, addressing: implied})
	register(&Operator{code: 0x08, order: php, addressing: implied})
	register(&Operator{code: 0x68, order: pla, addressing: implied})
	register(&Operator{code: 0x28, order: plp, addressing: implied})

	register(&Operator{code: 0x4C, order: jmp, addressing: absolute})
	register(&Operator{code: 0x6C, order: jmp, addressing: indirect})

	register(&Operator{code: 0x20, order: jsr, addressing: absolute})

	register(&Operator{code: 0x60, order: rts, addressing: implied})
	register(&Operator{code: 0x40, order: rti, addressing: implied})

	register(&Operator{code: 0x90, order: bcc, addressing: relative})
	register(&Operator{code: 0xB0, order: bcs, addressing: relative})
	register(&Operator{code: 0xF0, order: beq, addressing: relative})
	register(&Operator{code: 0x30, order: bmi, addressing: relative})
	register(&Operator{code: 0xD0, order: bne, addressing: relative})
	register(&Operator{code: 0x10, order: bpl, addressing: relative})
	register(&Operator{code: 0x50, order: bvc, addressing: relative})
	register(&Operator{code: 0x70, order: bvs, addressing: relative})

	register(&Operator{code: 0x18, order: clc, addressing: implied})
	register(&Operator{code: 0xD8, order: cld, addressing: implied})
	register(&Operator{code: 0x58, order: cli, addressing: implied})
	register(&Operator{code: 0xB8, order: clv, addressing: implied})

	register(&Operator{code: 0x38, order: sec, addressing: implied})
	register(&Operator{code: 0xF8, order: sed, addressing: implied})
	register(&Operator{code: 0x78, order: sei, addressing: implied})

	register(&Operator{code: 0x00, order: brk, addressing: implied})

	register(&Operator{code: 0xEA, order: nop, addressing: implied})
}
