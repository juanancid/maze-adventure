package entities

import (
	"reflect"
)

type World struct {
	nextEntityID Entity
	components   map[reflect.Type]map[Entity]Component
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

func (w *World) Query(types ...reflect.Type) EntityList {
	if len(types) == 0 {
		return nil
	}

	matching := make(map[Entity]bool)
	firstType := types[0]
	for e := range w.components[firstType] {
		matching[e] = true
	}

	for _, t := range types[1:] {
		for e := range matching {
			if _, exists := w.components[t][e]; !exists {
				delete(matching, e)
			}
		}
	}

	result := make([]Entity, 0, len(matching))
	for e := range matching {
		result = append(result, e)
	}

	return result
}

func (w *World) QueryComponents(components ...Component) EntityList {
	var types []reflect.Type
	for _, c := range components {
		types = append(types, reflect.TypeOf(c))
	}
	return w.Query(types...)
}

func (w *World) RemoveEntity(entity Entity) {
	for _, componentMap := range w.components {
		delete(componentMap, entity)
	}
}
