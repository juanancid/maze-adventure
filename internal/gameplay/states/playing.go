package states

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/renderers"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/updaters"
)

type PlayingState struct {
	world        *entities.World
	levelManager *levels.Manager
	updaters     []Updater
	renderers    []Renderer
	eventBus     *events.Bus
}

type Updater interface {
	Update(w *entities.World)
}

type Renderer interface {
	Draw(world *entities.World, screen *ebiten.Image)
}

func NewPlayingState(levelManager *levels.Manager) *PlayingState {
	ps := &PlayingState{
		levelManager: levelManager,
		eventBus:     events.NewBus(),
	}
	ps.loadNextLevel()
	ps.setUpdaters()
	ps.setRenderers()
	ps.setupEventSubscriptions()
	return ps
}

func (ps *PlayingState) OnEnter() {
	// Initialize or reset state explicitly
}

func (ps *PlayingState) OnExit() {
	// Cleanup state explicitly
}

func (ps *PlayingState) Update() error {
	for _, updater := range ps.updaters {
		updater.Update(ps.world)
	}

	ps.eventBus.Process()
	return nil
}

func (ps *PlayingState) Draw(screen *ebiten.Image) {
	for _, renderer := range ps.renderers {
		renderer.Draw(ps.world, screen)
	}
}

// The existing helper methods (loadNextLevel, setUpdaters, etc.) are moved here unchanged.

func (ps *PlayingState) loadNextLevel() {
	levelConfig, err := ps.levelManager.NextLevel()
	if err != nil {
		panic(err)
	}

	ps.world = levels.CreateLevel(levelConfig)
}

func (ps *PlayingState) setUpdaters() {
	ps.updaters = []Updater{
		&updaters.InputControl{},
		&updaters.Movement{},
		&updaters.MazeCollisionSystem{},
		updaters.NewScoreSystem(ps.eventBus),
	}
}

func (ps *PlayingState) setRenderers() {
	ps.renderers = []Renderer{
		&renderers.MazeRenderer{},
		&renderers.SpriteRenderer{},
	}
}

func (ps *PlayingState) setupEventSubscriptions() {
	ps.eventBus.Subscribe(reflect.TypeOf(events.LevelCompletedEvent{}), ps.onLevelCompleted)
}

func (ps *PlayingState) onLevelCompleted(e events.Event) {
	ps.loadNextLevel()
}
