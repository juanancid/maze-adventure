package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
	"github.com/juanancid/maze-adventure/internal/events"
)

type ScoreSystem struct{}

func (ss *ScoreSystem) Update(w *ecs.World) {
	mazeLayout := w.Maze().Layout
	cellSize := w.Maze().CellSize

	inputControlledEntities := w.GetComponents(reflect.TypeOf(&components.InputControlled{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))

	for entity, _ := range inputControlledEntities {
		pos := positions[entity].(*components.Position)
		size := sizes[entity].(*components.Size)

		// Compute the player's center
		centerX := pos.X + size.Width/2
		centerY := pos.Y + size.Height/2

		// Determine the cell the player is in
		col := int(centerX / float64(cellSize))
		row := int(centerY / float64(cellSize))

		if col == mazeLayout.Cols()-1 && row == mazeLayout.Rows()-1 {
			w.EmitEvent(events.LevelCompletedEvent{})
			return
		}
	}
}
