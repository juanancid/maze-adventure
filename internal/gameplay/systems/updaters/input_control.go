package updaters

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
)

type InputControl struct{}

func NewInputControl() InputControl {
	return InputControl{}
}

func (is InputControl) Update(w *entities.World) {
	entitiesToControl := w.QueryComponents(&components.InputControlled{}, &components.Velocity{})
	for _, entity := range entitiesToControl {
		handleInput(w, entity)
	}
}

func handleInput(w *entities.World, entity entities.Entity) {
	control := w.GetComponent(entity, reflect.TypeOf(&components.InputControlled{})).(*components.InputControlled)
	velocity := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)

	updateVelocityFromInput(control, velocity)
}

func updateVelocityFromInput(control *components.InputControlled, vel *components.Velocity) {
	vel.DX, vel.DY = 0, 0
	if ebiten.IsKeyPressed(control.MoveLeftKey) {
		vel.DX = -1
	}
	if ebiten.IsKeyPressed(control.MoveRightKey) {
		vel.DX = 1
	}
	if ebiten.IsKeyPressed(control.MoveUpKey) {
		vel.DY = -1
	}
	if ebiten.IsKeyPressed(control.MoveDownKey) {
		vel.DY = 1
	}
}
