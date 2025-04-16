package levels

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Manager struct {
	currentLevel int
}

func NewManager() *Manager {
	return &Manager{
		currentLevel: 1,
	}
}

// NextLevel loads the next level configuration.
// It returns the level, a boolean indicating if there is a next level,
// and an error if there was a problem loading the level.
func (m *Manager) NextLevel() (*Level, bool, error) {
	if !m.hasLevel(m.currentLevel) {
		return nil, false, nil // no more levels
	}

	level, err := m.loadLevel(m.currentLevel)
	if err != nil {
		return nil, false, err // loading error
	}

	m.currentLevel++
	return level, true, nil // whether there's a *next* one
}

func (m *Manager) loadLevel(levelNumber int) (*Level, error) {
	filePath := m.getLevelFilePath(levelNumber)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var lvl Level
	if err = yaml.Unmarshal(data, &lvl); err != nil {
		return nil, err
	}

	lvl.Number = levelNumber
	return &lvl, nil
}

func (m *Manager) hasLevel(n int) bool {
	levelPath := m.getLevelFilePath(n)
	_, err := os.Stat(levelPath)

	return err == nil
}

func (m *Manager) getLevelFilePath(levelNumber int) string {
	return fmt.Sprintf("internal/gameplay/levels/configs/level_%02d.yaml", levelNumber)
}
