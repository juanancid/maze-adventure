package app

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/states"
)

type Game struct {
	stateManager *states.Manager
}

func NewGame() *Game {
	levelManager := levels.NewManager()
	initialState := states.NewPlayingState(levelManager)

	return &Game{
		stateManager: states.NewManager(initialState),
	}
}

func (g *Game) Update() error {
	return g.stateManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stateManager.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
