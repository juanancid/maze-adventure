package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
)

type Movement struct{}

func (ms *Movement) Update(w *ecs.World) {
	for entity, entityVelocity := range w.GetComponents(reflect.TypeOf(&components.Velocity{})) {
		velocity := entityVelocity.(*components.Velocity)
		position := w.GetComponent(entity, reflect.TypeOf(&components.Position{})).(*components.Position)

		position.X += velocity.DX
		position.Y += velocity.DY
	}
}
