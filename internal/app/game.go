package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/juanancid/maze-adventure/internal/debug"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/states"
)

type Game struct {
	stateManager *states.Manager
	showDebug    bool
}

func NewGame() *Game {
	levelManager := levels.NewManager()

	stateManager := states.NewManager(nil)
	bootState := states.NewBootState(stateManager, levelManager)
	stateManager.ChangeState(bootState)

	return &Game{
		stateManager: stateManager,
		showDebug:    false,
	}
}

func (g *Game) Update() error {
	debug.TrackFPS()

	if ebiten.IsKeyPressed(ebiten.KeyMeta) && ebiten.IsKeyPressed(ebiten.KeyD) {
		g.showDebug = !g.showDebug
	}

	return g.stateManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stateManager.Draw(screen)

	if g.showDebug {
		ebitenutil.DebugPrint(screen, debug.GetDebugInfo())
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
