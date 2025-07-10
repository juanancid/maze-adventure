package levels

import (
	"fmt"

	"github.com/juanancid/maze-adventure/internal/gameplay/levels/definitions"
)

type Manager struct {
	currentLevel int
}

func NewManager() *Manager {
	return &Manager{
		currentLevel: 0,
	}
}

// NewManagerWithStartingLevel creates a new level manager starting at a specific level
func NewManagerWithStartingLevel(startingLevel int) (*Manager, error) {
	if startingLevel < 1 || startingLevel > len(definitions.LevelRegistry) {
		return nil, fmt.Errorf("invalid starting level %d: must be between 1 and %d", startingLevel, len(definitions.LevelRegistry))
	}

	return &Manager{
		currentLevel: startingLevel - 1, // Set to one before the desired level so NextLevel() returns the correct level
	}, nil
}

// NextLevel returns the next level configuration.
// It returns the level, a boolean indicating if there is a next level,
// and an error if there was a problem loading the level.
func (m *Manager) NextLevel() (levelConfig definitions.LevelConfig, levelNumber int, found bool) {
	m.currentLevel++

	if m.currentLevel > len(definitions.LevelRegistry) {
		levelConfig = definitions.EmptyLevelConfig
		levelNumber = 0
		found = false
		return
	}

	levelConfig = definitions.LevelRegistry[m.currentLevel-1]()
	levelNumber = m.currentLevel
	found = true
	return
}

// GetCurrentLevel returns the current level configuration without advancing.
// This is useful for restarting the current level.
func (m *Manager) GetCurrentLevel() (levelConfig definitions.LevelConfig, levelNumber int, found bool) {
	if m.currentLevel <= 0 || m.currentLevel > len(definitions.LevelRegistry) {
		levelConfig = definitions.EmptyLevelConfig
		levelNumber = 0
		found = false
		return
	}

	levelConfig = definitions.LevelRegistry[m.currentLevel-1]()
	levelNumber = m.currentLevel
	found = true
	return
}

// GetCurrentLevelNumber returns the current level number (1-based)
func (m *Manager) GetCurrentLevelNumber() int {
	return m.currentLevel
}

// GetTotalLevels returns the total number of available levels
func (m *Manager) GetTotalLevels() int {
	return len(definitions.LevelRegistry)
}

// IsValidLevel checks if a level number is valid
func IsValidLevel(levelNumber int) bool {
	return levelNumber >= 1 && levelNumber <= len(definitions.LevelRegistry)
}
