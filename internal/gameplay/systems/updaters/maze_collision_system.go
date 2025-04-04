package updaters

import (
	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/engine/layout"
)

// MazeCollisionSystem ensures entities do not pass through maze walls.
type MazeCollisionSystem struct{}

func (mcs *MazeCollisionSystem) Update(w *entities.World) {
	maze, ok := queries.GetMaze(w)
	if !ok {
		return
	}
	mazeLayout := maze.Layout
	cellSize := maze.CellSize

	entityList := w.QueryComponents(&components.Position{}, &components.Size{}, &components.Velocity{})

	for _, entity := range entityList {
		pos := entityList.GetPosition(w, entity)
		size := entityList.GetSize(w, entity)
		vel := entityList.GetVelocity(w, entity)

		handleEntityCollision(pos, size, vel, mazeLayout, cellSize)
	}
}

func handleEntityCollision(pos *components.Position, size *components.Size, vel *components.Velocity, mazeLayout layout.Layout, cellSize int) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the player is in
	centerX, centerY := entityBounds.center()
	col, row := cellIndices(centerX, centerY, float64(cellSize))

	// Check if out of mazeLayout bounds
	if !isCellValid(mazeLayout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	checkAndResolveWallCollision(pos, size, vel, col, row, cellSize, mazeLayout)
}

func cellIndices(x, y, cellSize float64) (col, row int) {
	col = int(x / cellSize)
	row = int(y / cellSize)
	return
}

// isCellValid explicitly checks if cell coordinates are valid
func isCellValid(layout layout.Layout, col, row int) bool {
	return col >= 0 && col < layout.Cols() && row >= 0 && row < layout.Rows()
}

func checkAndResolveWallCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellSize int, mazeLayout layout.Layout) {
	checkCurrentCellBoundaryCollision(pos, size, vel, col, row, cellSize, mazeLayout)
	checkNeighborCellEdgeCollision(pos, size, vel, col, row, cellSize, mazeLayout)
	checkNeighborCellBoundaryCollision(pos, size, vel, col, row, cellSize, mazeLayout)
}

func checkCurrentCellBoundaryCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellSize int, mazeLayout layout.Layout) {
	currentCell := mazeLayout.GetCell(col, row)

	// Check collisions with walls based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, cellSize) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * cellSize)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellSize) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*cellSize) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellSize) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*cellSize) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, cellSize) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * cellSize)
		}
	}
}

func checkNeighborCellEdgeCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellSize int, mazeLayout layout.Layout) {
	// Check collisions with edges based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, cellSize) && row > 0 { // Moving UP
		if crossedLeftBoundary(pos, col, cellSize) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			crossedRightBoundary(pos, size, col, cellSize) && mazeLayout.GetCellAbove(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64(row * cellSize)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellSize) && col < mazeLayout.Cols()-1 { // Moving RIGHT
		if crossedTopBoundary(pos, row, cellSize) && mazeLayout.GetCellRight(col, row).HasTopWall() ||
			crossedBottomBoundary(pos, size, row, cellSize) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64((col+1)*cellSize) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellSize) && row < mazeLayout.Rows()-1 { // Moving DOWN
		if crossedLeftBoundary(pos, col, cellSize) && mazeLayout.GetCellBelow(col, row).HasLeftWall() ||
			crossedRightBoundary(pos, size, col, cellSize) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*cellSize) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, cellSize) && col > 0 { // Moving LEFT
		if crossedTopBoundary(pos, row, cellSize) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			crossedBottomBoundary(pos, size, row, cellSize) && mazeLayout.GetCellLeft(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64(col * cellSize)
		}
	}
}

func checkNeighborCellBoundaryCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellSize int, mazeLayout layout.Layout) {
	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, cellSize) { // Moving UP
		if col > 0 && crossedLeftBoundary(pos, col, cellSize) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && crossedRightBoundary(pos, size, col, cellSize) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * cellSize)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellSize) { // Moving RIGHT
		if row > 0 && crossedTopBoundary(pos, row, cellSize) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && crossedBottomBoundary(pos, size, row, cellSize) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*cellSize) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellSize) { // Moving DOWN
		if col > 0 && crossedLeftBoundary(pos, col, cellSize) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && crossedRightBoundary(pos, size, col, cellSize) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*cellSize) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, cellSize) { // Moving LEFT
		if row > 0 && crossedTopBoundary(pos, row, cellSize) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && crossedBottomBoundary(pos, size, row, cellSize) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * cellSize)
		}
	}
}

func crossedTopBoundary(pos *components.Position, row, cellSize int) bool {
	return pos.Y < float64(row*cellSize)
}

func crossedRightBoundary(pos *components.Position, size *components.Size, col, cellSize int) bool {
	return pos.X+size.Width > float64((col+1)*cellSize)
}

func crossedBottomBoundary(pos *components.Position, size *components.Size, row, cellSize int) bool {
	return pos.Y+size.Height > float64((row+1)*cellSize)
}

func crossedLeftBoundary(pos *components.Position, col, cellSize int) bool {
	return pos.X < float64(col*cellSize)
}
