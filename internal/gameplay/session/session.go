package session

import (
	"github.com/juanancid/maze-adventure/internal/gameplay/config"
)

type GameSession struct {
	Score         int
	CurrentLevel  int
	MaxHearts     int
	CurrentHearts int
	Config        config.GameConfig
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
