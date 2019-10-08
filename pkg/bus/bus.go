package bus

type Bus interface {
	Read(uint16) byte
	Write(uint16, byte)
}
