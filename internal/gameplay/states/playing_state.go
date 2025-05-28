package states

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/renderers"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/updaters"
)

type PlayingState struct {
	stateManager *Manager
	levelManager *levels.Manager
	config       config.GameConfig

	gameSession *session.GameSession
	eventBus    *events.Bus

	world     *entities.World
	updaters  []Updater
	renderers []Renderer
}

type Updater interface {
	Update(world *entities.World, gameSession *session.GameSession)
}

type Renderer interface {
	Draw(world *entities.World, gameSession *session.GameSession, screen *ebiten.Image)
}

func NewPlayingState(stateManager *Manager, levelManager *levels.Manager, config config.GameConfig) *PlayingState {
	ps := &PlayingState{
		stateManager: stateManager,
		levelManager: levelManager,
		config:       config,
		gameSession:  session.NewGameSession(config),
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
		updater.Update(s.world, s.gameSession)
	}

	s.eventBus.Process()
	return nil
}

func (s *PlayingState) Draw(screen *ebiten.Image) {
	for _, renderer := range s.renderers {
		renderer.Draw(s.world, s.gameSession, screen)
	}
}

// Helper methods

func (s *PlayingState) loadNextLevel() {
	levelConfig, levelNumber, hasMore := s.levelManager.NextLevel()

	if !hasMore {
		// No more levels to load, trigger game complete event
		s.eventBus.Publish(events.GameComplete{})
		return
	}

	s.gameSession.CurrentLevel = levelNumber
	s.world = levels.CreateLevel(levelConfig)
}

func (s *PlayingState) setUpdaters() {
	s.updaters = []Updater{
		updaters.NewInputControl(),
		updaters.NewMovement(),
		updaters.NewMazeCollision(s.eventBus),
		updaters.NewExitCollision(s.eventBus),
		updaters.NewCollectiblePickup(s.eventBus),
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
	s.eventBus.Subscribe(reflect.TypeOf(events.CollectiblePicked{}), s.OnCollectiblePicked)
	s.eventBus.Subscribe(reflect.TypeOf(events.LevelCompletedEvent{}), s.onLevelCompleted)
	s.eventBus.Subscribe(reflect.TypeOf(events.GameComplete{}), s.onGameCompleted)
	s.eventBus.Subscribe(reflect.TypeOf(events.PlayerDamaged{}), s.onPlayerDamaged)
}

func (s *PlayingState) OnCollectiblePicked(e events.Event) {
	utils.PlaySound(utils.SoundCollectibleBip)
	s.gameSession.Score += e.(events.CollectiblePicked).Value
}

func (s *PlayingState) onLevelCompleted(e events.Event) {
	utils.PlaySound(utils.SoundLevelCompleted)
	s.loadNextLevel()
}

func (s *PlayingState) onGameCompleted(e events.Event) {
	endState := NewEndState(s.stateManager)
	s.stateManager.ChangeState(endState)
}

func (s *PlayingState) onPlayerDamaged(e events.Event) {
	s.gameSession.TakeDamage()

	// If player has no hearts left, game over
	if !s.gameSession.IsAlive() {
		// TODO: Implement game over state
		panic("player damaged")
	}
}
