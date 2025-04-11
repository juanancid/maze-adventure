package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
)

type ExitCollision struct {
	eventBus *events.Bus
}

func NewExitCollision(eventBus *events.Bus) *ExitCollision {
	return &ExitCollision{
		eventBus: eventBus,
	}
}

func (ec ExitCollision) Update(w *entities.World) {
	// Get all entities with Position, Size, and Exit components
	exits := w.QueryComponents(&components.Position{}, &components.Size{}, &components.Exit{})
	if len(exits) == 0 {
		return
	}

	// Get all entities with Position, Size, and InputControlled components (player)
	players := w.QueryComponents(&components.Position{}, &components.Size{}, &components.InputControlled{})
	if len(players) == 0 {
		return
	}

	// Get the first exit and player (there should only be one of each)
	var exitEntity, playerEntity entities.Entity
	for _, entity := range exits {
		exitEntity = entity
		break
	}
	for _, entity := range players {
		playerEntity = entity
		break
	}

	// Get their components
	exitPos := w.GetComponent(exitEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	exitSize := w.GetComponent(exitEntity, reflect.TypeOf(&components.Size{})).(*components.Size)
	playerPos := w.GetComponent(playerEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	playerSize := w.GetComponent(playerEntity, reflect.TypeOf(&components.Size{})).(*components.Size)

	// Check for collision
	if checkCollision(exitPos, exitSize, playerPos, playerSize) {
		ec.eventBus.Publish(events.LevelCompletedEvent{})
	}
}

func checkCollision(pos1 *components.Position, size1 *components.Size, pos2 *components.Position, size2 *components.Size) bool {
	return pos1.X < pos2.X+size2.Width &&
		pos1.X+size1.Width > pos2.X &&
		pos1.Y < pos2.Y+size2.Height &&
		pos1.Y+size1.Height > pos2.Y
}
