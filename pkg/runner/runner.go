package runner

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kanataxa/famigone/pkg/bus"
	"github.com/kanataxa/famigone/pkg/cassette"
	"github.com/kanataxa/famigone/pkg/cpu"
	"github.com/kanataxa/famigone/pkg/ppu"
	"golang.org/x/xerrors"
)

type NESRunner struct {
	cpu *cpu.CPU
	ppu *ppu.PPU
}

func New(path string) (*NESRunner, error) {
	c, err := cassette.New(path)
	if err != nil {
		return nil, xerrors.Errorf("failed to init cassette: %w", err)
	}

	p := ppu.New()
	b := bus.NewCPUBus(c, p)
	cp := cpu.New(b)
	return &NESRunner{
		cpu: cp,
		ppu: p,
	}, nil
}

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

func (nr *NESRunner) Run() error {
	if err := ebiten.RunGame(nr); err != nil {
		return xerrors.Errorf("failed to run: %w", err)
	}
	return nil
}

func (nr *NESRunner) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (nr *NESRunner) Update() error {
	nr.cpu.Run()
	nr.ppu.Run()
	return nil
}

func (nr *NESRunner) Draw(screen *ebiten.Image) {
}
