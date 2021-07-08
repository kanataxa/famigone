package cpu

import (
	"fmt"
	"github.com/kanataxa/famigone/pkg/bus"
	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/ppu"
	"log"
	"path/filepath"
	"strings"
	"testing"
)


func TestCPU(t *testing.T) {
	ca, err := cassette.New(filepath.Join("..","..","testdata", "nestest.nes"))
	if err != nil {
		log.Fatal(err)
	}

	p := ppu.New()
	b := bus.NewCPUBus(ca, p)
	c := New(b)
	c.register.PC = uint16(0xC000)
	for {
		currentPC := c.register.PC
		op := Lookup(c.Current())
		c.Run()
		c.Log(op, currentPC)

	}

}

func (c *CPU) Log(op *Operator, pc uint16) {
	latest := c.register.PC
	c.register.PC = pc
	addressingValue := c.addressingValue()
	fmt.Printf("%s %s %s %s A:%s X:%s Y:%s\n",
		strings.ToUpper(fmt.Sprintf("%04x", pc)),
		strings.ToUpper(fmt.Sprintf("%02x", op.code)),
		strings.ToUpper(fmt.Sprintf("%04x", addressingValue)),
		op.order,
		strings.ToUpper(fmt.Sprintf("%02x", c.register.A)),
		strings.ToUpper(fmt.Sprintf("%02x", c.register.X)),
		strings.ToUpper(fmt.Sprintf("%02x", c.register.Y)),
	)
	c.register.PC = latest
}
