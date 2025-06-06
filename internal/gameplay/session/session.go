package session

import (
	"fmt"

	"github.com/juanancid/maze-adventure/internal/gameplay/config"
)

type GameSession struct {
	Score         int
	CurrentLevel  int
	MaxHearts     int
	CurrentHearts int
	Config        config.GameConfig
	// Timer fields
	TimerEnabled   bool    // Whether the current level has a timer
	TimerRemaining float64 // Remaining time in seconds (float for smooth countdown)
	TimerTotal     int     // Total time for the level in seconds
}

// NewGameSession creates a new game session with the specified configuration
func NewGameSession(config config.GameConfig) *GameSession {
	return &GameSession{
		Score:         0,
		CurrentLevel:  0,
		MaxHearts:     config.StartingHearts,
		CurrentHearts: config.StartingHearts,
		Config:        config,
	}
}

// TakeDamage reduces the player's health by one heart
func (g *GameSession) TakeDamage() {
	if g.CurrentHearts > 0 {
		g.CurrentHearts--
	}
}

// Heal restores one heart of health
func (g *GameSession) Heal() {
	if g.CurrentHearts < g.MaxHearts {
		g.CurrentHearts++
	}
}

// IsAlive returns true if the player has any hearts remaining
func (g *GameSession) IsAlive() bool {
	return g.CurrentHearts > 0
}

// SetTimer initializes the timer for a level
func (g *GameSession) SetTimer(timerSeconds int) {
	if timerSeconds > 0 {
		g.TimerEnabled = true
		g.TimerTotal = timerSeconds
		g.TimerRemaining = float64(timerSeconds)
	} else {
		g.TimerEnabled = false
		g.TimerTotal = 0
		g.TimerRemaining = 0
	}
}

// UpdateTimer decreases the timer by the given delta time
func (g *GameSession) UpdateTimer(deltaTime float64) {
	if g.TimerEnabled && g.TimerRemaining > 0 {
		g.TimerRemaining -= deltaTime
		if g.TimerRemaining < 0 {
			g.TimerRemaining = 0
		}
	}
}

// IsTimerExpired returns true if the timer has reached zero
func (g *GameSession) IsTimerExpired() bool {
	return g.TimerEnabled && g.TimerRemaining <= 0
}

// GetTimerDisplayTime returns the timer in MM:SS format
func (g *GameSession) GetTimerDisplayTime() string {
	if !g.TimerEnabled {
		return ""
	}

	totalSeconds := int(g.TimerRemaining)
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
