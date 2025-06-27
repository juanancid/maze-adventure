package updaters

import (
	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
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

		enforceEntityMazeCollisions(pos, size, vel, maze, mc.eventBus)
	}
}

func enforceEntityMazeCollisions(pos *components.Position, size *components.Size, vel *components.Velocity, maze *components.Maze, eventBus *events.Bus) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the player is in
	centerX, centerY := entityBounds.center()
	col, row := convertWorldPositionToCellCoordinates(centerX, centerY, float64(maze.CellWidth), float64(maze.CellHeight))

	// Check if out of mazeLayout bounds
	if !isCellWithinMazeBounds(maze.Layout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	// Get the cell at the player's position
	cell := maze.Layout.GetCell(col, row)

	// Check the cell type and handle accordingly
	if cell.IsDeadly() {
		// Only emit damage event if there's an actual wall collision
		if preventAllWallPenetrations(pos, size, vel, col, row, maze) {
			eventBus.Publish(events.PlayerDamaged{Amount: 1})
			// Move the player to the center of the cell to prevent immediate re-collision
			centerEntityInCell(pos, size, col, row, maze.CellWidth, maze.CellHeight)
		}
	} else if cell.IsRegular() {
		preventAllWallPenetrations(pos, size, vel, col, row, maze)
	} else if cell.IsFreezing() {
		// Freezing cells still have wall collision, but also apply speed reduction
		preventAllWallPenetrations(pos, size, vel, col, row, maze)
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
func preventAllWallPenetrations(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	collided = false

	// Check current cell boundaries
	if preventWallPenetrationInCurrentCell(pos, size, vel, col, row, maze) {
		collided = true
	}

	// Check neighbor cell edges
	if preventWallPenetrationAtDiagonalEdges(pos, size, vel, col, row, maze) {
		collided = true
	}

	// Check neighbor cell boundaries
	if preventWallPenetrationInAdjacentCells(pos, size, vel, col, row, maze) {
		collided = true
	}

	return collided
}

func preventWallPenetrationInCurrentCell(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	mazeLayout := maze.Layout
	currentCell := mazeLayout.GetCell(col, row)
	collided = false

	// Check collisions with walls based on the velocity direction
	if vel.DY < 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			collided = true
		}
	}

	if vel.DX > 0 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			collided = true
		}
	}

	if vel.DY > 0 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			collided = true
		}
	}

	if vel.DX < 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			collided = true
		}
	}

	return collided
}

func preventWallPenetrationAtDiagonalEdges(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	mazeLayout := maze.Layout
	collided = false

	// Check collisions with edges based on the velocity direction
	if vel.DY < 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) && row > 0 { // Moving UP
		if isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			collided = true
		}
	}

	if vel.DX > 0 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) && col < mazeLayout.Cols()-1 { // Moving RIGHT
		if isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasTopWall() ||
			isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			collided = true
		}
	}

	if vel.DY > 0 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && row < mazeLayout.Rows()-1 { // Moving DOWN
		if isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasLeftWall() ||
			isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			collided = true
		}
	}

	if vel.DX < 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) && col > 0 { // Moving LEFT
		if isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			collided = true
		}
	}

	return collided
}

func preventWallPenetrationInAdjacentCells(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	mazeLayout := maze.Layout
	collided = false

	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) { // Moving UP
		if col > 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			collided = true
		}
	}

	if vel.DX > 0 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if row > 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			collided = true
		}
	}

	if vel.DY > 0 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if col > 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && isCollidingWithRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			collided = true
		}
	}

	if vel.DX < 0 && isCollidingWithLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if row > 0 && isCollidingWithTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && isCollidingWithBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			collided = true
		}
	}

	return collided
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

// centerEntityInCell positions the entity at the center of the specified cell
func centerEntityInCell(pos *components.Position, size *components.Size, col, row, cellWidth, cellHeight int) {
	// Calculate the center position of the cell
	centerX := float64(col*cellWidth) + float64(cellWidth)/2
	centerY := float64(row*cellHeight) + float64(cellHeight)/2

	// Adjust position to center the entity
	pos.X = centerX - size.Width/2
	pos.Y = centerY - size.Height/2
}
