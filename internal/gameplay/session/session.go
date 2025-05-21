package session

type GameSession struct {
	Score         int
	CurrentLevel  int
	MaxHearts     int
	CurrentHearts int
}

// NewGameSession creates a new game session with the specified health configuration
func NewGameSession(maxHearts int) *GameSession {
	return &GameSession{
		Score:         0,
		CurrentLevel:  0,
		MaxHearts:     maxHearts,
		CurrentHearts: maxHearts,
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
