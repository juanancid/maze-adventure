package session

import (
	"fmt"
	"time"

	"github.com/juanancid/maze-adventure/internal/gameplay/config"
)

const (
	// DefaultFreezeDuration is the default time a player is frozen when entering a freezing cell
	DefaultFreezeDuration = 2500 * time.Millisecond // 2.5 seconds
	// DefaultDamageCooldown is the default time before a player can take damage again
	DefaultDamageCooldown = 1500 * time.Millisecond // 1.5 seconds
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
	// Freeze fields
	IsFrozen          bool          // Whether the player is currently frozen
	FreezeStartTime   time.Time     // When the freeze effect started
	FreezeDuration    time.Duration // How long the freeze effect lasts
	LastFreezeCellCol int           // Last cell where freeze was applied (to prevent re-triggering)
	LastFreezeCellRow int           // Last cell where freeze was applied (to prevent re-triggering)
	CurrentCellCol    int           // Current cell column position
	CurrentCellRow    int           // Current cell row position
	// Damage cooldown fields
	LastDamageTime    time.Time     // When the player last took damage
	DamageCooldown    time.Duration // How long to wait before taking damage again
	LastDamageCellCol int           // Last cell where damage was applied (to prevent re-triggering)
	LastDamageCellRow int           // Last cell where damage was applied (to prevent re-triggering)
}

// NewGameSession creates a new game session with the specified configuration
func NewGameSession(config config.GameConfig) *GameSession {
	return &GameSession{
		Score:             0,
		CurrentLevel:      0,
		MaxHearts:         config.StartingHearts,
		CurrentHearts:     config.StartingHearts,
		Config:            config,
		LastFreezeCellCol: -1, // -1 indicates no previous freeze cell
		LastFreezeCellRow: -1,
		CurrentCellCol:    -1, // -1 indicates uninitialized
		CurrentCellRow:    -1,
		DamageCooldown:    DefaultDamageCooldown,
		LastDamageCellCol: -1, // -1 indicates no previous damage cell
		LastDamageCellRow: -1,
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

// Freeze-related methods

// SetCell updates the player's current cell position
func (g *GameSession) SetCell(col, row int) {
	g.CurrentCellCol = col
	g.CurrentCellRow = row
}

// HasCellChanged checks if the player has moved to a different cell
func (g *GameSession) HasCellChanged(col, row int) bool {
	return g.CurrentCellCol != col || g.CurrentCellRow != row
}

// StartFreeze immobilizes the player for the specified duration
func (g *GameSession) StartFreeze(duration time.Duration) {
	g.IsFrozen = true
	g.FreezeStartTime = time.Now()
	g.FreezeDuration = duration
	g.LastFreezeCellCol = g.CurrentCellCol
	g.LastFreezeCellRow = g.CurrentCellRow
}

// UpdateFreezeState checks if the freeze duration has expired and updates state accordingly
func (g *GameSession) UpdateFreezeState() {
	if g.IsFrozen && time.Since(g.FreezeStartTime) >= g.FreezeDuration {
		g.IsFrozen = false
	}
}

// CanApplyFreezeEffect checks if freeze can be applied (not already frozen and not in the same cell as last freeze)
func (g *GameSession) CanApplyFreezeEffect() bool {
	return !g.IsFrozen
}

// IsImmobilized returns true if the player cannot move due to freeze effect
func (g *GameSession) IsImmobilized() bool {
	return g.IsFrozen
}

// Damage cooldown methods

// CanApplyDamageEffect checks if damage can be applied (not in cooldown and not in the same cell as last damage)
func (g *GameSession) CanApplyDamageEffect() bool {
	// Check if enough time has passed since last damage
	timeSinceLastDamage := time.Since(g.LastDamageTime)
	cooldownExpired := timeSinceLastDamage >= g.DamageCooldown

	return cooldownExpired
}

// ApplyDamageWithCooldown applies damage and sets the cooldown
func (g *GameSession) ApplyDamageWithCooldown() {
	g.TakeDamage()
	g.LastDamageTime = time.Now()
	g.LastDamageCellCol = g.CurrentCellCol
	g.LastDamageCellRow = g.CurrentCellRow
}
