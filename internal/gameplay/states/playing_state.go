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

func (s *PlayingState) OnEnter() {
	// Initialize or reset state explicitly
}

func (s *PlayingState) OnExit() {
	// Cleanup state explicitly
}

func (s *PlayingState) Update() error {
	for _, updater := range s.updaters {
		updater.Update(s.world)
	}

	s.eventBus.Process()
	return nil
}

func (s *PlayingState) Draw(screen *ebiten.Image) {
	for _, renderer := range s.renderers {
		renderer.Draw(s.world, screen)
	}
}

// The existing helper methods (loadNextLevel, setUpdaters, etc.) are moved here unchanged.

func (s *PlayingState) loadNextLevel() {
	levelConfig, hasMore := s.levelManager.NextLevel()

	if !hasMore {
		// No more levels to load, trigger game complete event
		s.eventBus.Publish(events.GameComplete{})
		return
	}

	s.world = levels.CreateLevel(levelConfig)
}

func (s *PlayingState) setUpdaters() {
	s.updaters = []Updater{
		updaters.NewInputControl(),
		updaters.NewMovement(),
		updaters.NewMazeCollision(),
		updaters.NewExitCollision(s.eventBus),
		updaters.NewCollectiblePickup(),
	}
}

func (s *PlayingState) setRenderers() {
	s.renderers = []Renderer{
		renderers.NewMaze(),
		renderers.NewSprite(),
		renderers.NewHUD(),
	}
}

func (s *PlayingState) setupEventSubscriptions() {
	s.eventBus.Subscribe(reflect.TypeOf(events.LevelCompletedEvent{}), s.onLevelCompleted)
	s.eventBus.Subscribe(reflect.TypeOf(events.GameComplete{}), s.onGameCompleted)
}

func (s *PlayingState) onLevelCompleted(e events.Event) {
	s.loadNextLevel()
}

func (s *PlayingState) onGameCompleted(e events.Event) {
	endState := NewEndState(s.manager)
	s.manager.ChangeState(endState)
}
