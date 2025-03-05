package ecs

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/ecs/events"
)

type World struct {
	nextEntityID Entity
	components   map[reflect.Type]map[Entity]Component
	events       []events.Event
}

type Entity int

type Component interface{}

func NewWorld() *World {
	return &World{
		nextEntityID: 0,
		components:   make(map[reflect.Type]map[Entity]Component),
	}
}

func (w *World) NewEntity() Entity {
	id := w.nextEntityID
	w.nextEntityID++
	return id
}

func (w *World) AddComponent(entity Entity, component Component) {
	componentType := reflect.TypeOf(component)
	if w.components[componentType] == nil {
		w.components[componentType] = make(map[Entity]Component)
	}
	w.components[componentType][entity] = component
}

func (w *World) GetComponent(entity Entity, componentType reflect.Type) Component {
	return w.components[componentType][entity]
}

func (w *World) GetComponents(componentType reflect.Type) map[Entity]Component {
	return w.components[componentType]
}

func (w *World) EmitEvent(e events.Event) {
	w.events = append(w.events, e)
}

func (w *World) DrainEvents() []events.Event {
	ee := w.events
	w.events = nil
	return ee
}
