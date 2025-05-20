package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/debug"
	"github.com/juanancid/maze-adventure/internal/engine/input"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/states"
)

type Game struct {
	stateManager *states.Manager
	debugSystem  *debug.System
}

func NewGame() *Game {
	levelManager := levels.NewManager()

	stateManager := states.NewManager(nil)
	bootState := states.NewBootState(stateManager, levelManager)
	stateManager.ChangeState(bootState)

	inputHandler := input.NewHandler()
	debugSystem := debug.NewSystem(inputHandler)

	return &Game{
		stateManager: stateManager,
		debugSystem:  debugSystem,
	}
}

func (g *Game) Update() error {
	g.debugSystem.Update()
	return g.stateManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stateManager.Draw(screen)

	if g.debugSystem.IsDebugEnabled() {
		ebitenutil.DebugPrint(screen, debug.GetDebugInfo())
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
