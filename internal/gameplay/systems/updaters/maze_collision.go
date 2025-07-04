package updaters

import (
	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
	"time"
)

// MazeCollision ensures entities do not pass through maze walls.
type MazeCollision struct {
	eventBus *events.Bus
}

func NewMazeCollision(eventBus *events.Bus) MazeCollision {
	return MazeCollision{
		eventBus: eventBus,
	}
}

// Update checks for player collisions with the maze and handles resulting effects.
func (mc MazeCollision) Update(world *entities.World, gameSession *session.GameSession) {
	maze, ok := queries.GetMazeComponent(world)
	if !ok {
		return
	}

	entityList := world.QueryComponents(&components.Position{}, &components.Size{}, &components.Velocity{})
	for _, entity := range entityList {
		pos := entityList.GetPosition(world, entity)
		size := entityList.GetSize(world, entity)
		vel := entityList.GetVelocity(world, entity)

		enforcePlayerMazeCollisions(pos, size, vel, gameSession, maze, mc.eventBus)
	}
}

// enforcePlayerMazeCollisions handles collision and cell effects for player entities
func enforcePlayerMazeCollisions(pos *components.Position, size *components.Size, vel *components.Velocity, gameSession *session.GameSession, maze *components.Maze, eventBus *events.Bus) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the player is in
	centerX, centerY := entityBounds.center()
	col, row := convertWorldPositionToCellCoordinates(centerX, centerY, float64(maze.CellWidth), float64(maze.CellHeight))

	// Check if out of mazeLayout bounds
	if !isCellWithinMazeBounds(maze.Layout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	// Check if player has moved to a different cell
	if gameSession.HasCellChanged(col, row) {
		gameSession.SetCell(col, row)

	}

	// Handle wall collisions
	wallCollisionOccurred := preventAllWallPenetrations(pos, size, vel, col, row, maze)
	if wallCollisionOccurred {
		cell := maze.Layout.GetCell(col, row)

		if cell.IsFreezing() && gameSession.CanApplyFreezeEffect() {
			// Emit freeze event when entering freezing cell
			eventBus.Publish(events.PlayerFrozen{Duration: int(session.DefaultFreezeDuration / time.Millisecond)})
		}
		if cell.IsDeadly() && gameSession.CanApplyDamageEffect() {
			// Emit damage event when entering deadly cell (with cooldown check)
			eventBus.Publish(events.PlayerDamaged{Amount: 1})
		}
	}
}

func convertWorldPositionToCellCoordinates(x, y, cellWidth, cellHeight float64) (col, row int) {
	col = int(x / cellWidth)
	row = int(y / cellHeight)
	return
}

// isCellWithinMazeBounds checks if the cell coordinates are within the maze bounds
func isCellWithinMazeBounds(layout components.Layout, col, row int) bool {
	return col >= 0 && col < layout.Cols() && row >= 0 && row < layout.Rows()
}

// preventAllWallPenetrations returns true if there was a wall collision
func preventAllWallPenetrations(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (wallCollisionOccurred bool) {
	wallCollisionOccurred = false

	// Check current cell boundaries
	if preventWallPenetrationInCurrentCell(pos, size, vel, col, row, maze) {
		wallCollisionOccurred = true
	}

	// Check neighbor cell edges
	if preventWallPenetrationAtDiagonalEdges(pos, size, vel, col, row, maze) {
		wallCollisionOccurred = true
	}

	// Check neighbor cell boundaries
	if preventWallPenetrationInAdjacentCells(pos, size, vel, col, row, maze) {
		wallCollisionOccurred = true
	}

	return wallCollisionOccurred
}

func preventWallPenetrationInCurrentCell(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (wallCollisionOccurred bool) {
	mazeLayout := maze.Layout
	currentCell := mazeLayout.GetCell(col, row)
	wallCollisionOccurred = false

	// Check collisions with walls based on the velocity direction
	if vel.DY < 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			wallCollisionOccurred = true
		}
	}

	if vel.DX > 0 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			wallCollisionOccurred = true
		}
	}

	if vel.DY > 0 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			wallCollisionOccurred = true
		}
	}

	if vel.DX < 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			wallCollisionOccurred = true
		}
	}

	return wallCollisionOccurred
}

func preventWallPenetrationAtDiagonalEdges(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (wallCollisionOccurred bool) {
	mazeLayout := maze.Layout
	wallCollisionOccurred = false

	// Check collisions with edges based on the velocity direction
	if vel.DY < 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) && row > 0 { // Moving UP
		if isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			wallCollisionOccurred = true
		}
	}

	if vel.DX > 0 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) && col < mazeLayout.Cols()-1 { // Moving RIGHT
		if isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasTopWall() ||
			isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			wallCollisionOccurred = true
		}
	}

	if vel.DY > 0 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && row < mazeLayout.Rows()-1 { // Moving DOWN
		if isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasLeftWall() ||
			isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			wallCollisionOccurred = true
		}
	}

	if vel.DX < 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) && col > 0 { // Moving LEFT
		if isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			wallCollisionOccurred = true
		}
	}

	return wallCollisionOccurred
}

func preventWallPenetrationInAdjacentCells(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (wallCollisionOccurred bool) {
	mazeLayout := maze.Layout
	wallCollisionOccurred = false

	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) { // Moving UP
		if col > 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			wallCollisionOccurred = true
		}
	}

	if vel.DX > 0 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if row > 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			wallCollisionOccurred = true
		}
	}

	if vel.DY > 0 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if col > 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			wallCollisionOccurred = true
		}
	}

	if vel.DX < 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if row > 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			wallCollisionOccurred = true
		}
	}

	return wallCollisionOccurred
}

// isCollidingWithTopWall checks if entity collides with the top wall of a cell
func isCollidingWithTopWall(pos *components.Position, row, cellHeight int) bool {
	return pos.Y < float64(row*cellHeight)
}

// isCollidingWithRightWall checks if entity collides with the right wall of a cell
func isCollidingWithRightWall(pos *components.Position, size *components.Size, col, cellWidth int) bool {
	return pos.X+size.Width > float64((col+1)*cellWidth)
}

// isCollidingWithBottomWall checks if entity collides with the bottom wall of a cell
func isCollidingWithBottomWall(pos *components.Position, size *components.Size, row, cellHeight int) bool {
	return pos.Y+size.Height > float64((row+1)*cellHeight)
}

// isCollidingWithLeftWall checks if entity collides with the left wall of a cell
func isCollidingWithLeftWall(pos *components.Position, col, cellWidth int) bool {
	return pos.X < float64(col*cellWidth)
}
