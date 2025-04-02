package states

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	current State
}

func NewManager(initial State) *Manager {
	if initial != nil {
		initial.OnEnter()
	}

	return &Manager{current: initial}
}

func (m *Manager) ChangeState(next State) {
	if m.current != nil {
		m.current.OnExit()
	}

	m.current = next

	if m.current != nil {
		m.current.OnEnter()
	}
}

func (m *Manager) Update() error {
	return m.current.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.current.Draw(screen)
}
