package levels

import (
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

// NextLevel returns the next level configuration.
// It returns the level, a boolean indicating if there is a next level,
// and an error if there was a problem loading the level.
func (m *Manager) NextLevel() (level Config, levelNumber int, found bool) {
	m.currentLevel++

	if m.currentLevel > len(definitions.LevelRegistry) {
		level = EmptyLevel
		levelNumber = 0
		found = false
		return
	}

	levelConfig := definitions.LevelRegistry[m.currentLevel-1]()
	level = Config{
		Number:       m.currentLevel,
		Maze:         levelConfig.Maze,
		Player:       levelConfig.Player,
		Exit:         levelConfig.Exit,
		Collectibles: levelConfig.Collectibles,
	}
	levelNumber = m.currentLevel
	found = true
	return
}
