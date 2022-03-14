package cassette

import (
	"encoding/binary"
	"io/ioutil"
	"os"

	"github.com/kanataxa/famigone/pkg/memory"
	"github.com/pkg/errors"
)

const (
	headerSize = 0x0010
	prgROMSize = 0x4000
	chrROMSize = 0x2000
)

var (
	iNESLineHead = string([]byte{0x4E, 0x45, 0x53, 0x1A})
)

/*
See: https://wiki.nesdev.com/w/index.php/INES#iNES_file_format
The format of the header is as follows:

0-3: Constant $4E $45 $53 $1A ("NES" followed by MS-DOS end-of-file)
4: Size of PRG ROM in 16 KB units
5: Size of CHR ROM in 8 KB units (Value 0 means the board uses CHR RAM)
6: Flags 6 - Mapper, mirroring, battery, trainer
7: Flags 7 - Mapper, VS/Playchoice, NES 2.0
8: Flags 8 - PRG-RAM size (rarely used extension)
9: Flags 9 - TV system (rarely used extension)
10: Flags 10 - TV system, PRG-RAM presence (unofficial, rarely used extension)
11-15: Unused padding (should be filled with zero, but some rippers put their name across bytes 7-15)
*/
type iNESHeader struct {
	NES      [4]byte
	SizePRG  byte
	SizeCHR  byte
	Mapper   byte
	VSMapper byte
	SizeRAM  byte
	_        [7]byte
}

func (i *iNESHeader) validate() error {
	if i == nil {
		return errors.New("header is nil")
	}

	if nes := string(i.NES[:]); nes != iNESLineHead {
		return errors.Errorf("invalid 0-3 header: [%s]. must be NES", nes)
	}
	return nil
}

type Cassette struct {
	header *iNESHeader
	prg    *memory.ROM
	chr    *memory.ROM
}

func (c *Cassette) ProgramROM() *memory.ROM {
	return c.prg
}

func (c *Cassette) CharacterROM() *memory.ROM {
	return c.chr
}

func newHeader(path string) (*iNESHeader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()

	var header iNESHeader
	if err := binary.Read(file, binary.LittleEndian, &header); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := header.validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	return &header, nil
}

func New(path string) (*Cassette, error) {
	header, err := newHeader(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	prg := data[headerSize : headerSize+int(header.SizePRG)*prgROMSize]
	chr := data[len(prg) : len(prg)+int(header.SizeCHR)*chrROMSize]

	return &Cassette{
		header: header,
		prg:    memory.NewROM(prg),
		chr:    memory.NewROM(chr),
	}, nil
}
