package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// PatrollerCollision handles collision detection between player and patrollers
type PatrollerCollision struct {
	eventBus *events.Bus
}

// NewPatrollerCollision creates a new patroller collision system
func NewPatrollerCollision(eventBus *events.Bus) PatrollerCollision {
	return PatrollerCollision{
		eventBus: eventBus,
	}
}

// Update checks for collisions between player and patrollers
func (pc PatrollerCollision) Update(world *entities.World, gameSession *session.GameSession) {
	// Get all player entities (entities with InputControlled component indicate player)
	playerEntities := world.QueryComponents(&components.Position{}, &components.Size{}, &components.InputControlled{})
	if len(playerEntities) == 0 {
		return // No player found
	}

	// Get all patroller entities
	patrollerEntities := world.QueryComponents(&components.Patroller{}, &components.Position{}, &components.Size{})

	// Check collisions between player and each patroller
	for _, playerEntity := range playerEntities {
		playerPos := world.GetComponent(playerEntity, reflect.TypeOf(&components.Position{}))
		playerSize := world.GetComponent(playerEntity, reflect.TypeOf(&components.Size{}))

		if playerPos == nil || playerSize == nil {
			continue
		}

		playerPosition := playerPos.(*components.Position)
		playerSizeComp := playerSize.(*components.Size)

		for _, patrollerEntity := range patrollerEntities {
			patrollerComp := world.GetComponent(patrollerEntity, reflect.TypeOf(&components.Patroller{}))
			patrollerPos := world.GetComponent(patrollerEntity, reflect.TypeOf(&components.Position{}))
			patrollerSize := world.GetComponent(patrollerEntity, reflect.TypeOf(&components.Size{}))

			if patrollerComp == nil || patrollerPos == nil || patrollerSize == nil {
				continue
			}

			patroller := patrollerComp.(*components.Patroller)
			patrollerPosition := patrollerPos.(*components.Position)
			patrollerSizeComp := patrollerSize.(*components.Size)

			// Only check collision with active patrollers
			if !patroller.IsPatrollerActive() {
				continue
			}

			// Check if player and patroller are colliding
			if isColliding(playerPosition, playerSizeComp, patrollerPosition, patrollerSizeComp) {
				// Emit damage event using existing event system
				pc.eventBus.Publish(events.PlayerDamaged{Amount: patroller.GetDamage()})
			}
		}
	}
}

// isColliding checks if two rectangular entities are overlapping
func isColliding(pos1 *components.Position, size1 *components.Size, pos2 *components.Position, size2 *components.Size) bool {
	// Calculate bounding boxes
	left1 := pos1.X
	right1 := pos1.X + size1.Width
	top1 := pos1.Y
	bottom1 := pos1.Y + size1.Height

	left2 := pos2.X
	right2 := pos2.X + size2.Width
	top2 := pos2.Y
	bottom2 := pos2.Y + size2.Height

	// Check for overlap
	return !(right1 <= left2 || right2 <= left1 || bottom1 <= top2 || bottom2 <= top1)
}
