package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/components"
	"github.com/juanancid/maze-adventure/internal/entities"
)

type Movement struct{}

func (ms *Movement) Update(w *entities.World) {
	for entity, entityVelocity := range w.GetComponents(reflect.TypeOf(&components.Velocity{})) {
		velocity := entityVelocity.(*components.Velocity)
		position := w.GetComponent(entity, reflect.TypeOf(&components.Position{})).(*components.Position)

		position.X += velocity.DX
		position.Y += velocity.DY
	}
}
