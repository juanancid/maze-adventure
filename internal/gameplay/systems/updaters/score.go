package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
)

type Score struct {
	eventBus *events.Bus
}

func NewScore(eventBus *events.Bus) *Score {
	return &Score{
		eventBus: eventBus,
	}
}

func (s Score) Update(w *entities.World) {
	maze, ok := queries.GetMaze(w)
	if !ok {
		return
	}
	mazeLayout := maze.Layout
	cellWidth := maze.CellWidth
	cellHeight := maze.CellHeight

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
		col := int(centerX / float64(cellWidth))
		row := int(centerY / float64(cellHeight))

		if col == mazeLayout.Cols()-1 && row == mazeLayout.Rows()-1 {
			s.eventBus.Publish(events.LevelCompletedEvent{})
		}
	}
}
