package systems

import (
	"github.com/juanancid/maze-adventure/internal/queries"
	"reflect"

	"github.com/juanancid/maze-adventure/internal/components"
	"github.com/juanancid/maze-adventure/internal/entities"
	"github.com/juanancid/maze-adventure/internal/events"
)

type ScoreSystem struct {
	eventBus *events.Bus
}

func NewScoreSystem(eventBus *events.Bus) *ScoreSystem {
	return &ScoreSystem{
		eventBus: eventBus,
	}
}

func (ss *ScoreSystem) Update(w *entities.World) {
	maze, ok := queries.GetMaze(w)
	if !ok {
		return
	}
	mazeLayout := maze.Layout
	cellSize := maze.CellSize

	inputControlledEntities := w.GetComponents(reflect.TypeOf(&components.InputControlled{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))

	for entity := range inputControlledEntities {
		pos := positions[entity].(*components.Position)
		size := sizes[entity].(*components.Size)

		// Compute the player's center
		centerX := pos.X + size.Width/2
		centerY := pos.Y + size.Height/2

		// Determine the cell the player is in
		col := int(centerX / float64(cellSize))
		row := int(centerY / float64(cellSize))

		if col == mazeLayout.Cols()-1 && row == mazeLayout.Rows()-1 {
			ss.eventBus.Publish(events.LevelCompletedEvent{})
		}
	}
}
