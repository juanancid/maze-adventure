package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

type Movement struct{}

func NewMovement() Movement {
	return Movement{}
}

func (ms Movement) Update(wold *entities.World, gameSession *session.GameSession) {
	entitiesToMove := wold.QueryComponents(&components.Velocity{}, &components.Position{})
	for _, entity := range entitiesToMove {
		moveEntity(wold, entity)
	}
}

func moveEntity(w *entities.World, entity entities.Entity) {
	pos := w.GetComponent(entity, reflect.TypeOf(&components.Position{})).(*components.Position)
	vel := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)

	pos.X += vel.DX
	pos.Y += vel.DY
}
