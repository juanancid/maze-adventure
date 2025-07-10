package updaters

import (
	"math"
	"reflect"
	"time"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// PatrollerMovement handles movement patterns for patroller NPCs
type PatrollerMovement struct {
	startTime time.Time
}

// NewPatrollerMovement creates a new patroller movement system
func NewPatrollerMovement() PatrollerMovement {
	return PatrollerMovement{
		startTime: time.Now(),
	}
}

// Update applies movement patterns to all patroller entities
func (pm PatrollerMovement) Update(world *entities.World, gameSession *session.GameSession) {
	// Get all patroller entities with movement components
	patrollerEntities := world.QueryComponents(&components.Patroller{}, &components.Position{}, &components.Velocity{})

	for _, entity := range patrollerEntities {
		patrollerComp := world.GetComponent(entity, reflect.TypeOf(&components.Patroller{}))
		positionComp := world.GetComponent(entity, reflect.TypeOf(&components.Position{}))
		velocityComp := world.GetComponent(entity, reflect.TypeOf(&components.Velocity{}))

		if patrollerComp == nil || positionComp == nil || velocityComp == nil {
			continue
		}

		patroller := patrollerComp.(*components.Patroller)
		position := positionComp.(*components.Position)
		velocity := velocityComp.(*components.Velocity)

		// Only move active patrollers
		if !patroller.IsPatrollerActive() {
			velocity.DX = 0
			velocity.DY = 0
			continue
		}

		// Apply movement pattern based on patroller ID and time
		pm.applyMovementPattern(patroller, position, velocity)
	}
}

// applyMovementPattern applies a specific movement pattern to a patroller
func (pm PatrollerMovement) applyMovementPattern(patroller *components.Patroller, position *components.Position, velocity *components.Velocity) {
	// Calculate elapsed time for smooth movement
	elapsed := time.Since(pm.startTime).Seconds()

	// Different movement patterns based on patroller ID
	switch patroller.ID % 3 {
	case 0:
		// Horizontal back-and-forth movement
		pm.applyHorizontalPatrol(patroller, velocity, elapsed)
	case 1:
		// Vertical back-and-forth movement
		pm.applyVerticalPatrol(patroller, velocity, elapsed)
	case 2:
		// Circular movement pattern
		pm.applyCircularPatrol(patroller, velocity, elapsed)
	}
}

// applyHorizontalPatrol creates a horizontal back-and-forth movement
func (pm PatrollerMovement) applyHorizontalPatrol(patroller *components.Patroller, velocity *components.Velocity, elapsed float64) {
	// Use sine wave for smooth back-and-forth movement
	speed := patroller.Speed * 0.3       // Slower than player
	direction := math.Sin(elapsed * 2.0) // 2.0 controls frequency

	velocity.DX = direction * speed
	velocity.DY = 0
}

// applyVerticalPatrol creates a vertical back-and-forth movement
func (pm PatrollerMovement) applyVerticalPatrol(patroller *components.Patroller, velocity *components.Velocity, elapsed float64) {
	// Use sine wave for smooth back-and-forth movement
	speed := patroller.Speed * 0.3       // Slower than player
	direction := math.Sin(elapsed * 1.5) // Different frequency for variety

	velocity.DX = 0
	velocity.DY = direction * speed
}

// applyCircularPatrol creates a circular movement pattern
func (pm PatrollerMovement) applyCircularPatrol(patroller *components.Patroller, velocity *components.Velocity, elapsed float64) {
	// Create circular motion using sine and cosine
	speed := patroller.Speed * 0.2 // Even slower for circular motion
	angle := elapsed * 1.0         // Controls rotation speed

	velocity.DX = math.Cos(angle) * speed
	velocity.DY = math.Sin(angle) * speed
}
