package updaters

import (
	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/engine/mazebuilder"
)

// MazeCollisionSystem ensures entities do not pass through maze walls.
type MazeCollisionSystem struct{}

func (mcs *MazeCollisionSystem) Update(w *entities.World) {
	maze, ok := queries.GetMaze(w)
	if !ok {
		return
	}
	mazeLayout := maze.Layout
	cellWidth := maze.CellWidth
	cellHeight := maze.CellHeight

	entityList := w.QueryComponents(&components.Position{}, &components.Size{}, &components.Velocity{})

	for _, entity := range entityList {
		pos := entityList.GetPosition(w, entity)
		size := entityList.GetSize(w, entity)
		vel := entityList.GetVelocity(w, entity)

		handleEntityCollision(pos, size, vel, mazeLayout, cellWidth, cellHeight)
	}
}

func handleEntityCollision(pos *components.Position, size *components.Size, vel *components.Velocity, mazeLayout mazebuilder.Layout, cellWidth, cellHeight int) {
	entityBounds := newBoundingBox(pos, size)

	// Determine the cell the player is in
	centerX, centerY := entityBounds.center()
	col, row := cellIndices(centerX, centerY, float64(cellWidth), float64(cellHeight))

	// Check if out of mazeLayout bounds
	if !isCellValid(mazeLayout, col, row) {
		vel.DX, vel.DY = 0, 0
		return
	}

	checkAndResolveWallCollision(pos, size, vel, col, row, cellWidth, cellHeight, mazeLayout)
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

func checkAndResolveWallCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellWidth, cellHeight int, mazeLayout mazebuilder.Layout) {
	checkCurrentCellBoundaryCollision(pos, size, vel, col, row, cellWidth, cellHeight, mazeLayout)
	checkNeighborCellEdgeCollision(pos, size, vel, col, row, cellWidth, cellHeight, mazeLayout)
	checkNeighborCellBoundaryCollision(pos, size, vel, col, row, cellWidth, cellHeight, mazeLayout)
}

func checkCurrentCellBoundaryCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellWidth, cellHeight int, mazeLayout mazebuilder.Layout) {
	currentCell := mazeLayout.GetCell(col, row)

	// Check collisions with walls based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, cellHeight) { // Moving UP
		if currentCell.HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * cellHeight)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellWidth) { // Moving RIGHT
		if currentCell.HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*cellWidth) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellHeight) { // Moving DOWN
		if currentCell.HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*cellHeight) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, cellWidth) { // Moving LEFT
		if currentCell.HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * cellWidth)
		}
	}
}

func checkNeighborCellEdgeCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellWidth, cellHeight int, mazeLayout mazebuilder.Layout) {
	// Check collisions with edges based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, cellHeight) && row > 0 { // Moving UP
		if crossedLeftBoundary(pos, col, cellWidth) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			crossedRightBoundary(pos, size, col, cellWidth) && mazeLayout.GetCellAbove(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64(row * cellHeight)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellWidth) && col < mazeLayout.Cols()-1 { // Moving RIGHT
		if crossedTopBoundary(pos, row, cellHeight) && mazeLayout.GetCellRight(col, row).HasTopWall() ||
			crossedBottomBoundary(pos, size, row, cellHeight) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64((col+1)*cellWidth) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellHeight) && row < mazeLayout.Rows()-1 { // Moving DOWN
		if crossedLeftBoundary(pos, col, cellWidth) && mazeLayout.GetCellBelow(col, row).HasLeftWall() ||
			crossedRightBoundary(pos, size, col, cellWidth) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*cellHeight) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, cellWidth) && col > 0 { // Moving LEFT
		if crossedTopBoundary(pos, row, cellHeight) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			crossedBottomBoundary(pos, size, row, cellHeight) && mazeLayout.GetCellLeft(col, row).HasBottomWall() {
			vel.DX = 0
			pos.X = float64(col * cellWidth)
		}
	}
}

func checkNeighborCellBoundaryCollision(pos *components.Position, size *components.Size, vel *components.Velocity, col, row int, cellWidth, cellHeight int, mazeLayout mazebuilder.Layout) {
	// Check collisions with other cells walls based on velocity direction
	if vel.DY < 0 && crossedTopBoundary(pos, row, cellHeight) { // Moving UP
		if col > 0 && crossedLeftBoundary(pos, col, cellWidth) && mazeLayout.GetCellLeft(col, row).HasTopWall() ||
			col < mazeLayout.Cols()-1 && crossedRightBoundary(pos, size, col, cellWidth) && mazeLayout.GetCellRight(col, row).HasTopWall() {
			vel.DY = 0
			pos.Y = float64(row * cellHeight)
		}
	}

	if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellWidth) { // Moving RIGHT
		if row > 0 && crossedTopBoundary(pos, row, cellHeight) && mazeLayout.GetCellAbove(col, row).HasRightWall() ||
			row < mazeLayout.Rows()-1 && crossedBottomBoundary(pos, size, row, cellHeight) && mazeLayout.GetCellBelow(col, row).HasRightWall() {
			vel.DX = 0
			pos.X = float64((col+1)*cellWidth) - size.Width
		}
	}

	if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellHeight) { // Moving DOWN
		if col > 0 && crossedLeftBoundary(pos, col, cellWidth) && mazeLayout.GetCellLeft(col, row).HasBottomWall() ||
			col < mazeLayout.Cols()-1 && crossedRightBoundary(pos, size, col, cellWidth) && mazeLayout.GetCellRight(col, row).HasBottomWall() {
			vel.DY = 0
			pos.Y = float64((row+1)*cellHeight) - size.Height
		}
	}

	if vel.DX < 0 && crossedLeftBoundary(pos, col, cellWidth) { // Moving LEFT
		if row > 0 && crossedTopBoundary(pos, row, cellHeight) && mazeLayout.GetCellAbove(col, row).HasLeftWall() ||
			row < mazeLayout.Rows()-1 && crossedBottomBoundary(pos, size, row, cellHeight) && mazeLayout.GetCellBelow(col, row).HasLeftWall() {
			vel.DX = 0
			pos.X = float64(col * cellWidth)
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
