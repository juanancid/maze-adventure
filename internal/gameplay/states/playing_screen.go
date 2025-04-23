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

type PlayingScreen struct {
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

func NewPlayingScreen(manager *Manager, levelManager *levels.Manager) *PlayingScreen {
	ps := &PlayingScreen{
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

func (s *PlayingScreen) OnEnter() {
	// Initialize or reset state explicitly
}

func (s *PlayingScreen) OnExit() {
	// Cleanup state explicitly
}

func (s *PlayingScreen) Update() error {
	for _, updater := range s.updaters {
		updater.Update(s.world)
	}

	s.eventBus.Process()
	return nil
}

func (s *PlayingScreen) Draw(screen *ebiten.Image) {
	for _, renderer := range s.renderers {
		renderer.Draw(s.world, screen)
	}
}

// The existing helper methods (loadNextLevel, setUpdaters, etc.) are moved here unchanged.

func (s *PlayingScreen) loadNextLevel() {
	levelConfig, hasMore := s.levelManager.NextLevel()

	if !hasMore {
		// No more levels to load, trigger game complete event
		s.eventBus.Publish(events.GameComplete{})
		return
	}

	s.world = levels.CreateLevel(levelConfig)
}

func (s *PlayingScreen) setUpdaters() {
	s.updaters = []Updater{
		updaters.NewInputControl(),
		updaters.NewMovement(),
		updaters.NewMazeCollision(),
		updaters.NewExitCollision(s.eventBus),
		updaters.NewCollectiblePickup(),
	}
}

func (s *PlayingScreen) setRenderers() {
	s.renderers = []Renderer{
		renderers.NewMaze(),
		renderers.NewSprite(),
		renderers.NewHUD(),
	}
}

func (s *PlayingScreen) setupEventSubscriptions() {
	s.eventBus.Subscribe(reflect.TypeOf(events.LevelCompletedEvent{}), s.onLevelCompleted)
	s.eventBus.Subscribe(reflect.TypeOf(events.GameComplete{}), s.onGameCompleted)
}

func (s *PlayingScreen) onLevelCompleted(e events.Event) {
	s.loadNextLevel()
}

func (s *PlayingScreen) onGameCompleted(e events.Event) {
	endScreen := NewEndScreen(s.manager)
	s.manager.ChangeState(endScreen)
}
