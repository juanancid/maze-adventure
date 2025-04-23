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
func (m *Manager) NextLevel() (level Level, found bool) {
	m.currentLevel++

	if m.currentLevel > len(definitions.LevelRegistry) {
		level = EmptyLevel
		found = false
		return
	}

	levelConfig := definitions.LevelRegistry[m.currentLevel-1]()
	level = Level{
		Number: m.currentLevel,
		Maze:   levelConfig.Maze,
		Player: levelConfig.Player,
		Exit:   levelConfig.Exit,
	}
	found = true
	return
}
