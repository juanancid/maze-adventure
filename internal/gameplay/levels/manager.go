package levels

type Manager struct {
	CurrentLevel int
}

func NewManager() *Manager {
	return &Manager{CurrentLevel: 1}
}

func (m *Manager) NextLevel() (*Level, error) {
	level, err := Load(m.CurrentLevel)
	if err != nil {
		return nil, err
	}
	m.CurrentLevel++
	return level, nil
}
