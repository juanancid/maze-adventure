package entities

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
)

// EntityList helps iterate explicitly over entities
type EntityList []Entity

func (entities EntityList) GetPosition(w *World, entity Entity) *components.Position {
	return w.GetComponent(entity, reflect.TypeOf(&components.Position{})).(*components.Position)
}

func (entities EntityList) GetSize(w *World, entity Entity) *components.Size {
	return w.GetComponent(entity, reflect.TypeOf(&components.Size{})).(*components.Size)
}

func (entities EntityList) GetVelocity(w *World, entity Entity) *components.Velocity {
	return w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)
}
