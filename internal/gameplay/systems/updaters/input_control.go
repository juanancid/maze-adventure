package updaters

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
)

const (
	vx = 1
	vy = 1
)

type InputControl struct{}

func (is *InputControl) Update(w *entities.World) {
	inputControlledEntities := w.GetComponents(reflect.TypeOf(&components.InputControlled{}))

	for entity, entityControl := range inputControlledEntities {
		control := entityControl.(*components.InputControlled)
		velocity := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)

		velocity.DX = 0
		velocity.DY = 0

		if ebiten.IsKeyPressed(control.MoveLeftKey) {
			velocity.DX = -vx
		}
		if ebiten.IsKeyPressed(control.MoveRightKey) {
			velocity.DX = vx
		}
		if ebiten.IsKeyPressed(control.MoveUpKey) {
			velocity.DY = -vy
		}
		if ebiten.IsKeyPressed(control.MoveDownKey) {
			velocity.DY = vy
		}
	}
}
