package states

import (
	"log"
	"reflect"
	"time"

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
	world, err := levels.CreateLevel(levelConfig)
	if err != nil {
		// Critical error: level creation failed
		log.Printf("CRITICAL: Failed to create level %d: %v", levelNumber, err)
		log.Printf("This is a serious bug that needs investigation. Keeping current level to prevent crash.")

		// Reset the timer for current level as a fallback
		if s.world != nil {
			s.gameSession.SetTimer(levelConfig.Timer)
		} else {
			// If we don't even have a world, this is catastrophic - trigger game complete to exit gracefully
			log.Printf("CATASTROPHIC: No world available, ending game to prevent crash")
			s.eventBus.Publish(events.GameComplete{})
		}
		return
	}
	s.world = world

	// Initialize the timer for this level
	s.gameSession.SetTimer(levelConfig.Timer)
}

func (s *PlayingState) setUpdaters() {
	s.updaters = []Updater{
		updaters.NewInputControl(),
		updaters.NewMovement(),
		updaters.NewMazeCollision(s.eventBus),
		updaters.NewExitCollision(s.eventBus),
		updaters.NewCollectiblePickup(s.eventBus),
		updaters.NewTimer(s.eventBus),
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
	s.eventBus.Subscribe(reflect.TypeOf(events.TimerExpired{}), s.onTimerExpired)
	s.eventBus.Subscribe(reflect.TypeOf(events.PlayerFrozen{}), s.onPlayerFrozen)
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
	victoryState := NewVictoryState(s.stateManager)
	s.stateManager.ChangeState(victoryState)
}

func (s *PlayingState) onPlayerDamaged(e events.Event) {
	// Check if damage can be applied (respects cooldown)
	if !s.gameSession.CanApplyDamageEffect() {
		return // Skip damage if in cooldown period
	}

	utils.PlaySound(utils.SoundDamage)
	s.gameSession.ApplyDamageWithCooldown()

	// If player has no hearts left, game over
	if !s.gameSession.IsAlive() {
		s.triggerGameOver()
	}
}

func (s *PlayingState) triggerGameOver() {
	gameOverState := NewGameOverState(s.stateManager)
	s.stateManager.ChangeState(gameOverState)
}

func (s *PlayingState) onTimerExpired(e events.Event) {
	utils.PlaySound(utils.SoundDamage)

	// Player takes damage when timer expires
	s.gameSession.TakeDamage()

	// Check if player is still alive
	if !s.gameSession.IsAlive() {
		// Player died, trigger game over
		s.triggerGameOver()
	} else {
		// Player is still alive, reset timer for current level
		s.resetTimerForCurrentLevel()
	}
}

func (s *PlayingState) onPlayerFrozen(e events.Event) {
	utils.PlaySound(utils.SoundFreeze)
	frozenEvent := e.(events.PlayerFrozen)
	duration := time.Duration(frozenEvent.Duration) * time.Millisecond
	s.gameSession.StartFreeze(duration)
}

func (s *PlayingState) resetTimerForCurrentLevel() {
	levelConfig, _, hasLevel := s.levelManager.GetCurrentLevel()

	if !hasLevel {
		// No current level, shouldn't happen
		return
	}

	// Reset the timer for this level without recreating the world
	s.gameSession.SetTimer(levelConfig.Timer)
}
