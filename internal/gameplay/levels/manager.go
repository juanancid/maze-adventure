package levels

type Manager struct {
	current int
}

func NewManager() *Manager {
	return &Manager{current: 1}
}

func (m *Manager) NextLevel() (*Level, error) {
	level, err := Load(m.current)
	if err != nil {
		return nil, err
	}
	m.current++
	return level, nil
}
