package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// PatrollerMazeCollision ensures patrollers do not pass through maze walls
// but allows them to move through passages freely
type PatrollerMazeCollision struct{}

// NewPatrollerMazeCollision creates a new patroller maze collision system
func NewPatrollerMazeCollision() PatrollerMazeCollision {
	return PatrollerMazeCollision{}
}

// Update checks for patroller collisions with maze walls and prevents penetration
func (pmc PatrollerMazeCollision) Update(world *entities.World, gameSession *session.GameSession) {
	maze, ok := queries.GetMazeComponent(world)
	if !ok {
		return
	}

	// Get all patroller entities with position, size, and velocity
	patrollerEntities := world.QueryComponents(&components.Patroller{}, &components.Position{}, &components.Size{}, &components.Velocity{})

	for _, entity := range patrollerEntities {
		patrollerComp := world.GetComponent(entity, reflect.TypeOf(&components.Patroller{}))
		positionComp := world.GetComponent(entity, reflect.TypeOf(&components.Position{}))
		sizeComp := world.GetComponent(entity, reflect.TypeOf(&components.Size{}))
		velocityComp := world.GetComponent(entity, reflect.TypeOf(&components.Velocity{}))

		if patrollerComp == nil || positionComp == nil || sizeComp == nil || velocityComp == nil {
			continue
		}

		patroller := patrollerComp.(*components.Patroller)
		position := positionComp.(*components.Position)
		size := sizeComp.(*components.Size)
		velocity := velocityComp.(*components.Velocity)

		// Only handle collision for active patrollers
		if !patroller.IsPatrollerActive() {
			continue
		}

		// Enforce maze collision for this patroller
		pmc.enforcePatrollerMazeCollisions(position, size, velocity, maze)
	}
}

// enforcePatrollerMazeCollisions handles wall collision for patroller entities
// This is similar to the player collision logic but without cell effects
func (pmc PatrollerMazeCollision) enforcePatrollerMazeCollisions(pos *components.Position, size *components.Size, vel *components.Velocity, maze *components.Maze) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the patroller is in
	centerX, centerY := entityBounds.center()
	col, row := pmc.convertWorldPositionToCellCoordinates(centerX, centerY, float64(maze.CellWidth), float64(maze.CellHeight))

	// Check if out of maze bounds - stop movement if outside
	if !pmc.isCellWithinMazeBounds(maze.Layout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	// Handle wall collisions (without cell effects)
	pmc.preventAllWallPenetrations(pos, size, vel, col, row, maze)
}

// convertWorldPositionToCellCoordinates converts world coordinates to cell coordinates
func (pmc PatrollerMazeCollision) convertWorldPositionToCellCoordinates(x, y, cellWidth, cellHeight float64) (col, row int) {
	col = int(x / cellWidth)
	row = int(y / cellHeight)
	return
}

// isCellWithinMazeBounds checks if the given cell coordinates are within maze bounds
func (pmc PatrollerMazeCollision) isCellWithinMazeBounds(layout components.Layout, col, row int) bool {
	return col >= 0 && col < layout.Cols() && row >= 0 && row < layout.Rows()
}

// preventAllWallPenetrations prevents patroller from passing through walls
// This reuses the same logic as player collision but without cell effects
func (pmc PatrollerMazeCollision) preventAllWallPenetrations(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	// Check current cell boundaries
	pmc.preventWallPenetrationInCurrentCell(pos, size, vel, col, row, maze)

	// Check neighbor cell edges
	pmc.preventWallPenetrationAtDiagonalEdges(pos, size, vel, col, row, maze)

	// Check neighbor cell boundaries
	pmc.preventWallPenetrationInAdjacentCells(pos, size, vel, col, row, maze)
}

// preventWallPenetrationInCurrentCell prevents penetration of walls in the current cell
func (pmc PatrollerMazeCollision) preventWallPenetrationInCurrentCell(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	mazeLayout := maze.Layout
	currentCell := mazeLayout.GetCell(col, row)

	// Check collisions with walls based on the velocity direction
	if vel.DY < 0 && pmc.isCollidingWithTopWall(pos, row, maze.CellHeight) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
		}
	}

	if vel.DX > 0 && pmc.isCollidingWithRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
		}
	}

	if vel.DY > 0 && pmc.isCollidingWithBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
		}
	}

	if vel.DX < 0 && pmc.isCollidingWithLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
		}
	}
}

// preventWallPenetrationAtDiagonalEdges prevents penetration at diagonal cell edges
func (pmc PatrollerMazeCollision) preventWallPenetrationAtDiagonalEdges(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	mazeLayout := maze.Layout

	// Check diagonal collisions when moving diagonally
	if vel.DX > 0 && vel.DY < 0 { // Moving UP-RIGHT
		if pmc.isCollidingWithTopWall(pos, row, maze.CellHeight) && pmc.isCollidingWithRightWall(pos, size, col, maze.CellWidth) {
			if col < mazeLayout.Cols()-1 && row > 0 {
				topRightCell := mazeLayout.GetCell(col+1, row-1)
				if topRightCell.HasLeftWall() || topRightCell.HasBottomWall() {
					vel.DX = 0
					vel.DY = 0
				}
			}
		}
	}

	if vel.DX > 0 && vel.DY > 0 { // Moving DOWN-RIGHT
		if pmc.isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && pmc.isCollidingWithRightWall(pos, size, col, maze.CellWidth) {
			if col < mazeLayout.Cols()-1 && row < mazeLayout.Rows()-1 {
				bottomRightCell := mazeLayout.GetCell(col+1, row+1)
				if bottomRightCell.HasLeftWall() || bottomRightCell.HasTopWall() {
					vel.DX = 0
					vel.DY = 0
				}
			}
		}
	}

	if vel.DX < 0 && vel.DY > 0 { // Moving DOWN-LEFT
		if pmc.isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && pmc.isCollidingWithLeftWall(pos, col, maze.CellWidth) {
			if col > 0 && row < mazeLayout.Rows()-1 {
				bottomLeftCell := mazeLayout.GetCell(col-1, row+1)
				if bottomLeftCell.HasRightWall() || bottomLeftCell.HasTopWall() {
					vel.DX = 0
					vel.DY = 0
				}
			}
		}
	}

	if vel.DX < 0 && vel.DY < 0 { // Moving UP-LEFT
		if pmc.isCollidingWithTopWall(pos, row, maze.CellHeight) && pmc.isCollidingWithLeftWall(pos, col, maze.CellWidth) {
			if col > 0 && row > 0 {
				topLeftCell := mazeLayout.GetCell(col-1, row-1)
				if topLeftCell.HasRightWall() || topLeftCell.HasBottomWall() {
					vel.DX = 0
					vel.DY = 0
				}
			}
		}
	}
}

// preventWallPenetrationInAdjacentCells prevents penetration of walls in adjacent cells
func (pmc PatrollerMazeCollision) preventWallPenetrationInAdjacentCells(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	mazeLayout := maze.Layout

	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && pmc.isCollidingWithTopWall(pos, row, maze.CellHeight) { // Moving UP
		if col > 0 && pmc.isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && pmc.isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
		}
	}

	if vel.DX > 0 && pmc.isCollidingWithRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if row > 0 && pmc.isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && pmc.isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
		}
	}

	if vel.DY > 0 && pmc.isCollidingWithBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if col > 0 && pmc.isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && pmc.isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
		}
	}

	if vel.DX < 0 && pmc.isCollidingWithLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if row > 0 && pmc.isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && pmc.isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
		}
	}
}

// Wall collision detection helper functions (reused from maze_collision.go logic)

func (pmc PatrollerMazeCollision) isCollidingWithTopWall(pos *components.Position, row, cellHeight int) bool {
	return pos.Y < float64(row*cellHeight)
}

func (pmc PatrollerMazeCollision) isCollidingWithRightWall(pos *components.Position, size *components.Size, col, cellWidth int) bool {
	return pos.X+size.Width > float64((col+1)*cellWidth)
}

func (pmc PatrollerMazeCollision) isCollidingWithBottomWall(pos *components.Position, size *components.Size, row, cellHeight int) bool {
	return pos.Y+size.Height > float64((row+1)*cellHeight)
}

func (pmc PatrollerMazeCollision) isCollidingWithLeftWall(pos *components.Position, col, cellWidth int) bool {
	return pos.X < float64(col*cellWidth)
}
