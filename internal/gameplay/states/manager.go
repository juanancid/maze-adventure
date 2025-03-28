package states

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	current State
}

func NewManager(initial State) *Manager {
	initial.OnEnter()
	return &Manager{current: initial}
}

func (m *Manager) ChangeState(next State) {
	m.current.OnExit()
	m.current = next
	m.current.OnEnter()
}

func (m *Manager) Update() error {
	return m.current.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.current.Draw(screen)
}
