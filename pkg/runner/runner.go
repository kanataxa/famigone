package runner

import (
	"github.com/kanataxa/famigone/pkg/cpu"
	"github.com/kanataxa/famigone/pkg/ppu"
)

type Runner interface {
	Run() error
}

type NESRunner struct {
	cpu Runner
	ppu Runner
}

func (r *NESRunner) Run() error {
	for {
		r.cpu.Run()
		r.ppu.Run()
	}
}

func New(cpu *cpu.CPU, ppu *ppu.PPU) Runner {
	return &NESRunner{
		cpu: cpu,
		ppu: ppu,
	}
}
