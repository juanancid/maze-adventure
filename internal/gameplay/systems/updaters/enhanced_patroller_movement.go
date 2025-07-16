package updaters

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// EnhancedPatrollerMovement handles advanced movement patterns for patroller NPCs
type EnhancedPatrollerMovement struct {
	startTime time.Time
}

// NewEnhancedPatrollerMovement creates a new enhanced patroller movement system
func NewEnhancedPatrollerMovement() EnhancedPatrollerMovement {
	return EnhancedPatrollerMovement{
		startTime: time.Now(),
	}
}

// Update applies enhanced movement patterns to all patroller entities
func (epm EnhancedPatrollerMovement) Update(world *entities.World, gameSession *session.GameSession) {
	maze, ok := queries.GetMazeComponent(world)
	if !ok {
		return // No maze available
	}

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

		// Apply movement pattern based on patroller type
		epm.applyEnhancedMovementPattern(patroller, position, velocity, maze)
	}
}

// applyEnhancedMovementPattern applies the appropriate movement pattern
func (epm EnhancedPatrollerMovement) applyEnhancedMovementPattern(patroller *components.Patroller, position *components.Position, velocity *components.Velocity, maze *components.Maze) {
	elapsed := time.Since(epm.startTime).Seconds()

	switch patroller.PatrolType {
	case components.PatrolPatternRandom:
		epm.applyRandomMovement(patroller, position, velocity, maze, elapsed)
	case components.PatrolPatternLinear:
		epm.applyLinearMovement(patroller, position, velocity, maze, elapsed)
	case components.PatrolPatternPerimeter:
		epm.applyPerimeterMovement(patroller, position, velocity, maze, elapsed)
	case components.PatrolPatternCross:
		epm.applyCrossMovement(patroller, position, velocity, maze, elapsed)
	default:
		// Fallback to random movement
		epm.applyRandomMovement(patroller, position, velocity, maze, elapsed)
	}
}

// applyRandomMovement implements random directional movement at intersections
func (epm EnhancedPatrollerMovement) applyRandomMovement(patroller *components.Patroller, position *components.Position, velocity *components.Velocity, maze *components.Maze, elapsed float64) {
	// Check if we should change direction (every 1-3 seconds randomly)
	timeSinceLastChange := elapsed - patroller.State.LastDirectionChange
	shouldChangeDirection := timeSinceLastChange > (1.0 + rand.Float64()*2.0)

	if shouldChangeDirection || (velocity.DX == 0 && velocity.DY == 0) {
		// Get current cell position
		col, row := epm.getEntityCellPosition(position, maze)

		// Get available directions (not blocked by walls)
		availableDirections := epm.getAvailableDirections(col, row, maze)

		if len(availableDirections) > 0 {
			// Choose a random available direction
			newDirection := availableDirections[rand.Intn(len(availableDirections))]
			patroller.State.CurrentDirection = newDirection
			patroller.State.LastDirectionChange = elapsed
		}
	}

	// Apply movement in current direction
	epm.applyDirectionalMovement(patroller, velocity)
}

// applyLinearMovement implements back-and-forth movement along corridors
func (epm EnhancedPatrollerMovement) applyLinearMovement(patroller *components.Patroller, position *components.Position, velocity *components.Velocity, maze *components.Maze, elapsed float64) {
	col, row := epm.getEntityCellPosition(position, maze)

	// Check if we hit a wall and need to reverse direction
	if epm.isDirectionBlocked(col, row, patroller.State.CurrentDirection, maze) {
		// Reverse direction
		patroller.State.CurrentDirection = (patroller.State.CurrentDirection + 2) % 4
		patroller.State.LastDirectionChange = elapsed
	}

	// Apply movement in current direction
	epm.applyDirectionalMovement(patroller, velocity)
}

// applyPerimeterMovement implements wall-following behavior
func (epm EnhancedPatrollerMovement) applyPerimeterMovement(patroller *components.Patroller, position *components.Position, velocity *components.Velocity, maze *components.Maze, elapsed float64) {
	col, row := epm.getEntityCellPosition(position, maze)

	// Try to keep a wall on the right side (right-hand rule)
	rightDirection := (patroller.State.CurrentDirection + 1) % 4
	forwardDirection := patroller.State.CurrentDirection

	if !epm.isDirectionBlocked(col, row, rightDirection, maze) {
		// Turn right if possible
		patroller.State.CurrentDirection = rightDirection
	} else if epm.isDirectionBlocked(col, row, forwardDirection, maze) {
		// Turn left if forward is blocked
		patroller.State.CurrentDirection = (patroller.State.CurrentDirection + 3) % 4
	}
	// Otherwise continue forward

	patroller.State.LastDirectionChange = elapsed
	epm.applyDirectionalMovement(patroller, velocity)
}

// applyCrossMovement implements alternating horizontal and vertical movement
func (epm EnhancedPatrollerMovement) applyCrossMovement(patroller *components.Patroller, position *components.Position, velocity *components.Velocity, maze *components.Maze, elapsed float64) {
	timeSinceLastChange := elapsed - patroller.State.LastDirectionChange

	// Change between horizontal and vertical every 2-4 seconds
	shouldChangePhase := timeSinceLastChange > (2.0 + rand.Float64()*2.0)

	if shouldChangePhase {
		patroller.State.MovementPhase = (patroller.State.MovementPhase + 1) % 2
		patroller.State.LastDirectionChange = elapsed

		col, row := epm.getEntityCellPosition(position, maze)

		if patroller.State.MovementPhase == 0 {
			// Horizontal phase - choose left or right
			if !epm.isDirectionBlocked(col, row, 1, maze) { // Right
				patroller.State.CurrentDirection = 1
			} else if !epm.isDirectionBlocked(col, row, 3, maze) { // Left
				patroller.State.CurrentDirection = 3
			}
		} else {
			// Vertical phase - choose up or down
			if !epm.isDirectionBlocked(col, row, 0, maze) { // Up
				patroller.State.CurrentDirection = 0
			} else if !epm.isDirectionBlocked(col, row, 2, maze) { // Down
				patroller.State.CurrentDirection = 2
			}
		}
	}

	// Check if current direction is blocked and reverse if needed
	col, row := epm.getEntityCellPosition(position, maze)
	if epm.isDirectionBlocked(col, row, patroller.State.CurrentDirection, maze) {
		patroller.State.CurrentDirection = (patroller.State.CurrentDirection + 2) % 4
	}

	epm.applyDirectionalMovement(patroller, velocity)
}

// Helper functions

// getEntityCellPosition returns the cell coordinates of an entity
func (epm EnhancedPatrollerMovement) getEntityCellPosition(position *components.Position, maze *components.Maze) (col, row int) {
	col = int(position.X / float64(maze.CellWidth))
	row = int(position.Y / float64(maze.CellHeight))
	return
}

// getAvailableDirections returns directions that are not blocked by walls
func (epm EnhancedPatrollerMovement) getAvailableDirections(col, row int, maze *components.Maze) []int {
	var directions []int

	// Check all four directions (0=up, 1=right, 2=down, 3=left)
	for dir := 0; dir < 4; dir++ {
		if !epm.isDirectionBlocked(col, row, dir, maze) {
			directions = append(directions, dir)
		}
	}

	return directions
}

// isDirectionBlocked checks if movement in a direction is blocked by a wall
func (epm EnhancedPatrollerMovement) isDirectionBlocked(col, row, direction int, maze *components.Maze) bool {
	// Check maze bounds first
	if col < 0 || col >= maze.Layout.Cols() || row < 0 || row >= maze.Layout.Rows() {
		return true
	}

	cell := maze.Layout.GetCell(col, row)

	switch direction {
	case 0: // Up
		return cell.HasTopWall()
	case 1: // Right
		return cell.HasRightWall()
	case 2: // Down
		return cell.HasBottomWall()
	case 3: // Left
		return cell.HasLeftWall()
	default:
		return true
	}
}

// applyDirectionalMovement sets velocity based on current direction
func (epm EnhancedPatrollerMovement) applyDirectionalMovement(patroller *components.Patroller, velocity *components.Velocity) {
	speed := patroller.Speed

	switch patroller.State.CurrentDirection {
	case 0: // Up
		velocity.DX = 0
		velocity.DY = -speed
	case 1: // Right
		velocity.DX = speed
		velocity.DY = 0
	case 2: // Down
		velocity.DX = 0
		velocity.DY = speed
	case 3: // Left
		velocity.DX = -speed
		velocity.DY = 0
	default:
		velocity.DX = 0
		velocity.DY = 0
	}
}
