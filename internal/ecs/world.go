package ecs

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	currentLevel int
	nextEntityID Entity
	components   map[reflect.Type]map[Entity]Component
	systems      []System
	renderables  []Renderable
}

type Entity int

type Component interface{}

type System interface {
	Update(w *World)
}

type Renderable interface {
	Draw(world *World, screen *ebiten.Image)
}

func NewWorld() *World {
	return &World{
		nextEntityID: 0,
		components:   make(map[reflect.Type]map[Entity]Component),
		systems:      make([]System, 0),
	}
}

func (w *World) NewEntity() Entity {
	id := w.nextEntityID
	w.nextEntityID++
	return id
}

func (w *World) RemoveEntity(entity Entity) {
	// Remove the entity's components from each component map
	for componentType, entityMap := range w.components {
		delete(entityMap, entity)
		// Clean up empty maps to save memory
		if len(entityMap) == 0 {
			delete(w.components, componentType)
		}
	}
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

func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) Update() {
	for _, s := range w.systems {
		s.Update(w)
	}
}

func (w *World) GetSystems() []System {
	return w.systems
}

func (w *World) GetSystem(target System) System {
	for _, system := range w.systems {
		if reflect.TypeOf(system) == reflect.TypeOf(target) {
			return system
		}
	}
	return nil // Return nil if the system isn't found
}

func (w *World) AddRenderable(r Renderable) {
	w.renderables = append(w.renderables, r)
}

func (w *World) GetRenderable(target Renderable) Renderable {
	for _, renderable := range w.renderables {
		if reflect.TypeOf(renderable) == reflect.TypeOf(target) {
			return renderable
		}
	}
	return nil // Return nil if the renderable isn't found
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, r := range w.renderables {
		r.Draw(w, screen)
	}
}

func (w *World) SetCurrentLevel(level int) {
	w.currentLevel = level
}

func (w *World) GetCurrentLevel() int {
	return w.currentLevel
}
