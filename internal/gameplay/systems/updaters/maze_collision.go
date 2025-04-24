package updaters

import (
	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/engine/mazebuilder"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// MazeCollision ensures entities do not pass through maze walls.
type MazeCollision struct{}

func NewMazeCollision() MazeCollision {
	return MazeCollision{}
}

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

		handleEntityCollision(pos, size, vel, maze)
	}
}

func handleEntityCollision(pos *components.Position, size *components.Size, vel *components.Velocity, maze *components.Maze) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the player is in
	centerX, centerY := entityBounds.center()
	col, row := cellIndices(centerX, centerY, float64(maze.CellWidth), float64(maze.CellHeight))

	// Check if out of mazeLayout bounds
	if !isCellValid(maze.Layout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	checkAndResolveWallCollision(pos, size, vel, col, row, maze)
}

func cellIndices(x, y, cellWidth, cellHeight float64) (col, row int) {
	col = int(x / cellWidth)
	row = int(y / cellHeight)
	return
}

// isCellValid explicitly checks if cell coordinates are valid
func isCellValid(layout mazebuilder.Layout, col, row int) bool {
	return col >= 0 && col < layout.Cols() && row >= 0 && row < layout.Rows()
}

func checkAndResolveWallCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	checkCurrentCellBoundaryCollision(pos, size, vel, col, row, maze)
	checkNeighborCellEdgeCollision(pos, size, vel, col, row, maze)
	checkNeighborCellBoundaryCollision(pos, size, vel, col, row, maze)
}

func checkCurrentCellBoundaryCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	mazeLayout := maze.Layout
	currentCell := mazeLayout.GetCell(col, row)

	// Check collisions with walls based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, maze.CellHeight) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, maze.CellHeight) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, maze.CellWidth) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
		}
	}
}

func checkNeighborCellEdgeCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	mazeLayout := maze.Layout

	// Check collisions with edges based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, maze.CellHeight) && row > 0 { // Moving UP
		if crossedLeftBoundary(pos, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			crossedRightBoundary(pos, size, col, maze.CellWidth) && mazeLayout.GetCellAbove(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, maze.CellWidth) && col < mazeLayout.Cols()-1 { // Moving RIGHT
		if crossedTopBoundary(pos, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasTopWall() ||
			crossedBottomBoundary(pos, size, row, maze.CellHeight) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, maze.CellHeight) && row < mazeLayout.Rows()-1 { // Moving DOWN
		if crossedLeftBoundary(pos, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasLeftWall() ||
			crossedRightBoundary(pos, size, col, maze.CellWidth) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, maze.CellWidth) && col > 0 { // Moving LEFT
		if crossedTopBoundary(pos, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			crossedBottomBoundary(pos, size, row, maze.CellHeight) && mazeLayout.GetCellLeft(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
		}
	}
}

func checkNeighborCellBoundaryCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, maze *components.Maze) {
	mazeLayout := maze.Layout

	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, maze.CellHeight) { // Moving UP
		if col > 0 && crossedLeftBoundary(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && crossedRightBoundary(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * maze.CellHeight)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, maze.CellWidth) { // Moving RIGHT
		if row > 0 && crossedTopBoundary(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && crossedBottomBoundary(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*maze.CellWidth) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, maze.CellHeight) { // Moving DOWN
		if col > 0 && crossedLeftBoundary(pos, col, maze.CellWidth) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && crossedRightBoundary(pos, size, col, maze.CellWidth) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*maze.CellHeight) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, maze.CellWidth) { // Moving LEFT
		if row > 0 && crossedTopBoundary(pos, row, maze.CellHeight) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && crossedBottomBoundary(pos, size, row, maze.CellHeight) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * maze.CellWidth)
		}
	}
}

func crossedTopBoundary(pos *components.Position, row, cellHeight int) bool {
	return pos.Y < float64(row*cellHeight)
}

func crossedRightBoundary(pos *components.Position, size *components.Size, col, cellWidth int) bool {
	return pos.X+size.Width > float64((col+1)*cellWidth)
}

func crossedBottomBoundary(pos *components.Position, size *components.Size, row, cellHeight int) bool {
	return pos.Y+size.Height > float64((row+1)*cellHeight)
}

func crossedLeftBoundary(pos *components.Position, col, cellWidth int) bool {
	return pos.X < float64(col*cellWidth)
}
