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

		resolveMazeCollisionForEntity(pos, size, vel, maze, mc.eventBus)
	}
}

func resolveMazeCollisionForEntity(pos *components.Position, size *components.Size, vel *components.Velocity, maze *components.Maze, eventBus *events.Bus) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the player is in
	centerX, centerY := entityBounds.center()
	col, row := getCellCoordsFromWorldPos(centerX, centerY, float64(maze.CellWidth), float64(maze.CellHeight))

	// Check if out of mazeLayout bounds
	if !isCellValid(maze.Layout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	// Get the cell at the player's position
	cell := maze.Layout.GetCell(col, row)

	// Check cell type and handle accordingly
	switch cell.GetType() {
	case components.CellTypeDeadly:
		// Only emit damage event if there's an actual wall collision
		if resolveCollisionAgainstCellWalls(pos, size, vel, col, row, maze) {
			eventBus.Publish(events.PlayerDamaged{Amount: 1})
		}
	case components.CellTypeRegular:
		resolveCollisionAgainstCellWalls(pos, size, vel, col, row, maze)
	case components.CellTypeFreezing:
		// TODO: Implement freezing effect
		resolveCollisionAgainstCellWalls(pos, size, vel, col, row, maze)
	}
}

func getCellCoordsFromWorldPos(x, y, cellWidth, cellHeight float64) (col, row int) {
	col = int(x / cellWidth)
	row = int(y / cellHeight)
	return
}

// isCellValid explicitly checks if cell coordinates are valid
func isCellValid(layout components.Layout, col, row int) bool {
	return col >= 0 && col < layout.Cols() && row >= 0 && row < layout.Rows()
}

// resolveCollisionAgainstCellWalls returns true if there was a wall collision
func resolveCollisionAgainstCellWalls(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	collided = false

	// Check current cell boundaries
	if resolveCollisionAtCurrentCell(pos, size, vel, col, row, maze) {
		collided = true
	}

	// Check neighbor cell edges
	if resolveCollisionAtDiagonalEdges(pos, size, vel, col, row, maze) {
		collided = true
	}

	// Check neighbor cell boundaries
	if resolveCollisionAtAdjacentCells(pos, size, vel, col, row, maze) {
		collided = true
	}

	return collided
}

func resolveCollisionAtCurrentCell(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	mazeLayout := maze.Layout
	currentCell := mazeLayout.GetCell(col, row)
	collided = false

	// Check collisions with walls based on velocity direction
	if vel.DY < 0 && isBeyondTopWall(pos, row, maze.CellHeight) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			collided = true
		}
	}

	if vel.DX > 0 && isBeyondRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			collided = true
		}
	}

	if vel.DY > 0 && isBeyondBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			collided = true
		}
	}

	if vel.DX < 0 && isBeyondLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			collided = true
		}
	}

	return collided
}

func resolveCollisionAtDiagonalEdges(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	mazeLayout := maze.Layout
	collided = false

	// Check collisions with edges based on velocity direction
	if vel.DY < 0 && isBeyondTopWall(pos, row, maze.CellHeight) && row > 0 { // Moving UP
		if isBeyondLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			isBeyondRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			collided = true
		}
	}

	if vel.DX > 0 && isBeyondRightWall(pos, size, col, maze.CellWidth) && col < mazeLayout.Cols()-1 { // Moving RIGHT
		if isBeyondTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasTopWall() ||
			isBeyondBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			collided = true
		}
	}

	if vel.DY > 0 && isBeyondBottomWall(pos, size, row, maze.CellHeight) && row < mazeLayout.Rows()-1 { // Moving DOWN
		if isBeyondLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasLeftWall() ||
			isBeyondRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			collided = true
		}
	}

	if vel.DX < 0 && isBeyondLeftWall(pos, col, maze.CellWidth) && col > 0 { // Moving LEFT
		if isBeyondTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			isBeyondBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			collided = true
		}
	}

	return collided
}

func resolveCollisionAtAdjacentCells(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) (collided bool) {
	mazeLayout := maze.Layout
	collided = false

	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && isBeyondTopWall(pos, row, maze.CellHeight) { // Moving UP
		if col > 0 && isBeyondLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && isBeyondRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
			collided = true
		}
	}

	if vel.DX > 0 && isBeyondRightWall(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if row > 0 && isBeyondTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && isBeyondBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
			collided = true
		}
	}

	if vel.DY > 0 && isBeyondBottomWall(pos, size, row, maze.CellHeight) { // Moving DOWN
		if col > 0 && isBeyondLeftWall(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && isBeyondRightWall(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
			collided = true
		}
	}

	if vel.DX < 0 && isBeyondLeftWall(pos, col, maze.CellWidth) { // Moving LEFT
		if row > 0 && isBeyondTopWall(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && isBeyondBottomWall(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
			collided = true
		}
	}

	return collided
}

func isBeyondTopWall(pos *components.Position, row, cellHeight int) bool {
	return pos.Y < float64(row*cellHeight)
}

func isBeyondRightWall(pos *components.Position, size *components.Size, col, cellWidth int) bool {
	return pos.X+size.Width > float64((col+1)*cellWidth)
}

func isBeyondBottomWall(pos *components.Position, size *components.Size, row, cellHeight int) bool {
	return pos.Y+size.Height > float64((row+1)*cellHeight)
}

func isBeyondLeftWall(pos *components.Position, col, cellWidth int) bool {
	return pos.X < float64(col*cellWidth)
}
