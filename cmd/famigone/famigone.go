package main

import (
	"log"
	"path/filepath"

	"github.com/kanataxa/famigone/pkg/ppu"
	"github.com/kanataxa/famigone/pkg/runner"

	"github.com/kanataxa/famigone/pkg/bus"
	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/cpu"
)

func main() {
	c, err := cassette.New(filepath.Join("testdata", "hello_world.nes"))
	if err != nil {
		log.Fatal(err)
	}

	b := bus.NewCPUBus(c)
	cp := cpu.New(b)

	runner.New(cp, &ppu.PPU{}).Run()
}
