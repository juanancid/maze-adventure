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
	manager      *Manager
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

func NewPlayingState(manager *Manager, levelManager *levels.Manager) *PlayingState {
	ps := &PlayingState{
		manager:      manager,
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
		// No more levels explicitly available, trigger game over explicitly
		ps.eventBus.Publish(events.GameOverEvent{})
		return
	}

	ps.world = levels.CreateLevel(levelConfig)
}

func (ps *PlayingState) setUpdaters() {
	ps.updaters = []Updater{
		updaters.NewInputControl(),
		updaters.NewMovement(),
		updaters.NewMazeCollision(),
		updaters.NewScore(ps.eventBus),
	}
}

func (ps *PlayingState) setRenderers() {
	ps.renderers = []Renderer{
		renderers.NewMaze(),
		renderers.NewSprite(),
		renderers.NewHUD(),
	}
}

func (ps *PlayingState) setupEventSubscriptions() {
	ps.eventBus.Subscribe(reflect.TypeOf(events.LevelCompletedEvent{}), ps.onLevelCompleted)
	ps.eventBus.Subscribe(reflect.TypeOf(events.GameOverEvent{}), ps.onGameOver)
}

func (ps *PlayingState) onLevelCompleted(e events.Event) {
	ps.loadNextLevel()
}

func (ps *PlayingState) onGameOver(e events.Event) {
	gameOver := NewGameOverState(ps.manager)
	ps.manager.ChangeState(gameOver)
}
