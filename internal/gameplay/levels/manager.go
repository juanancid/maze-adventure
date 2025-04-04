package levels

import "errors"

const MaxLevels = 1

type Manager struct {
	current int
}

func NewManager() *Manager {
	return &Manager{current: 1}
}

func (m *Manager) NextLevel() (*Level, error) {
	if m.current > MaxLevels {
		return nil, errors.New("no more levels")
	}

	level, err := Load(m.current)
	if err != nil {
		return nil, err
	}
	m.current++
	return level, nil
}
