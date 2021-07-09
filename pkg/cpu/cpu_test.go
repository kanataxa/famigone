package cpu

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kanataxa/famigone/pkg/bus"
	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/ppu"
)

type NesTestLog struct {
	PC              string
	OPCode          string
	AddressingValue string
	OpType          string
	A               string
	X               string
	Y               string
	P               string
	SP              string
	PPU             []string
	CYC             string
}

func (n *NesTestLog) String() string {
	return fmt.Sprintf("%s %s %s %s A:%s X:%s Y:%s P:%s\n",
		n.PC,
		n.OPCode,
		n.AddressingValue,
		n.OpType,
		n.A,
		n.X,
		n.Y,
		n.P,
	)
}

func parse(t string) *NesTestLog {
	text := []rune(t)
	var result []string
	var tok []string
	for i := 0; i < len(text); i++ {
		r := string(text[i])
		if strings.ReplaceAll(r, " ", "") == "" {
			if len(tok) > 0 {
				result = append(result, strings.Join(tok, ""))
				tok = []string{}
			}
			continue
		}
		tok = append(tok, r)
	}

	return &NesTestLog{
		PC:     result[0],
		OPCode: result[1],
	}
}

func loadNesTestLog() ([]*NesTestLog, error) {
	f, err := os.Open(filepath.Join("..", "..", "testdata", "nestest.log"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var log []*NesTestLog
	for scanner.Scan() {
		text := scanner.Text()
		log = append(log, parse(text))
	}
	return log, nil
}

func TestCPU(t *testing.T) {
	ca, err := cassette.New(filepath.Join("..", "..", "testdata", "nestest.nes"))
	if err != nil {
		t.Fatal(err)
	}
	p := ppu.New()
	b := bus.NewCPUBus(ca, p)
	c := New(b)
	c.register.PC = uint16(0xC000)

	expectedLogs, err := loadNesTestLog()
	if err != nil {
		t.Fatal(err)
	}

	var executedLogs []string
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(strings.Join(executedLogs, ""))
			t.Error(r)
		}
	}()
	for i := 0; ; i++ {
		currentPC := c.register.PC
		capture := c.register.P.String()
		op := Lookup(c.Current())
		c.Run()
		log := c.Log(op, currentPC, capture)
		if log.PC != expectedLogs[i].PC {
			t.Errorf("invalid pc. line: %d expected: %s, but: %s", i, expectedLogs[i].PC, log.PC)
			fmt.Println(strings.Join(executedLogs, ""))
			return
		}
		if log.OPCode != expectedLogs[i].OPCode {
			t.Errorf("invalid opcode. line: %d expected: %s, but: %s", i, expectedLogs[i].OPCode, log.OPCode)
			fmt.Println(strings.Join(executedLogs, ""))
			return
		}
		executedLogs = append(executedLogs, log.String())

	}

}

func (c *CPU) Log(op *Operator, pc uint16, statusCapture string) *NesTestLog {
	latest := c.register.PC
	c.register.PC = pc
	addressingValue := c.addressingValue()
	/*
		fmt.Printf("%s %s %s %s A:%s X:%s Y:%s\n",
			strings.ToUpper(fmt.Sprintf("%04x", pc)),
			strings.ToUpper(fmt.Sprintf("%02x", op.code)),
			strings.ToUpper(fmt.Sprintf("%04x", addressingValue)),
			op.order,
			strings.ToUpper(fmt.Sprintf("%02x", c.register.A)),
			strings.ToUpper(fmt.Sprintf("%02x", c.register.X)),
			strings.ToUpper(fmt.Sprintf("%02x", c.register.Y)),
		)
	*/
	c.register.PC = latest

	return &NesTestLog{
		PC:              strings.ToUpper(fmt.Sprintf("%04x", pc)),
		OPCode:          strings.ToUpper(fmt.Sprintf("%02x", op.code)),
		AddressingValue: strings.ToUpper(fmt.Sprintf("%04x", addressingValue)),
		OpType:          op.order.String(),
		A:               strings.ToUpper(fmt.Sprintf("%02x", c.register.A)),
		X:               strings.ToUpper(fmt.Sprintf("%02x", c.register.X)),
		Y:               strings.ToUpper(fmt.Sprintf("%02x", c.register.Y)),
		P:               strings.ToUpper(statusCapture),
		SP:              strings.ToUpper(fmt.Sprintf("%02x", c.register.S)),
	}
}
