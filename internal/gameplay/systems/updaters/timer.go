package updaters

import (
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// Timer handles the level timer countdown
type Timer struct {
	eventBus *events.Bus
}

// NewTimer creates a new timer updater
func NewTimer(eventBus *events.Bus) *Timer {
	return &Timer{
		eventBus: eventBus,
	}
}

// Update decreases the timer and publishes timer expired event when needed
func (t *Timer) Update(world *entities.World, gameSession *session.GameSession) {
	if !gameSession.TimerEnabled {
		return
	}

	// Get delta time from Ebiten (1/60 for 60 FPS)
	deltaTime := 1.0 / 60.0

	// Store previous state to detect expiration
	wasExpired := gameSession.IsTimerExpired()

	// Update the timer
	gameSession.UpdateTimer(deltaTime)

	// Check if timer just expired
	if !wasExpired && gameSession.IsTimerExpired() {
		t.eventBus.Publish(events.TimerExpired{})
	}
}
