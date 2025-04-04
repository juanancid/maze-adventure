package components

import "github.com/hajimehoshi/ebiten/v2"

type InputControlled struct {
	MoveLeftKey  ebiten.Key
	MoveRightKey ebiten.Key
	MoveUpKey    ebiten.Key
	MoveDownKey  ebiten.Key
}
