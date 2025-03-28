package states

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
	OnEnter()
	OnExit()
}
