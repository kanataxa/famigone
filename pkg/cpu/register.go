package cpu

type Register struct {
	A  int8
	X  int8
	Y  int8
	S  int8
	P  StatusRegister
	PC int16
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
