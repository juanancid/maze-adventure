package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	engineconfig "github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/debug"
	"github.com/juanancid/maze-adventure/internal/engine/input"
	gameplayconfig "github.com/juanancid/maze-adventure/internal/gameplay/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/states"
)

type Game struct {
	stateManager *states.Manager
	debugSystem  *debug.System
	config       gameplayconfig.GameConfig
}

func NewGame(config gameplayconfig.GameConfig) *Game {
	var levelManager *levels.Manager
	var err error

	// Create level manager with appropriate starting level
	if config.StartingLevel > 1 {
		levelManager, err = levels.NewManagerWithStartingLevel(config.StartingLevel)
		if err != nil {
			// Fallback to default manager if there's an error
			levelManager = levels.NewManager()
		}
	} else {
		levelManager = levels.NewManager()
	}

	stateManager := states.NewManager(nil)
	bootState := states.NewBootState(stateManager, levelManager, config)
	stateManager.ChangeState(bootState)

	inputHandler := input.NewHandler()
	debugSystem := debug.NewSystem(inputHandler)

	return &Game{
		stateManager: stateManager,
		debugSystem:  debugSystem,
		config:       config,
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
	return engineconfig.ScreenWidth, engineconfig.ScreenHeight
}
