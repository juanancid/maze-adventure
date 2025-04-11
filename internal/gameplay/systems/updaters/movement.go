package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
)

type Movement struct{}

func NewMovement() Movement {
	return Movement{}
}

func (ms Movement) Update(w *entities.World) {
	entitiesToMove := w.QueryComponents(&components.Velocity{}, &components.Position{})
	for _, entity := range entitiesToMove {
		moveEntity(w, entity)
	}
}

func moveEntity(w *entities.World, entity entities.Entity) {
	pos := w.GetComponent(entity, reflect.TypeOf(&components.Position{})).(*components.Position)
	vel := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)

	pos.X += vel.DX
	pos.Y += vel.DY
}
