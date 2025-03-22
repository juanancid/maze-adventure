package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/components"
	"github.com/juanancid/maze-adventure/internal/ecs"
)

// MazeCollisionSystem ensures entities do not pass through maze walls.
type MazeCollisionSystem struct{}

func (mcs *MazeCollisionSystem) Update(w *ecs.World) {
	mazeLayout := w.Maze().Layout
	cellSize := w.Maze().CellSize

	players := w.GetComponents(reflect.TypeOf(&components.InputControlled{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))
	velocities := w.GetComponents(reflect.TypeOf(&components.Velocity{}))

	for player := range players {
		pos := positions[player].(*components.Position)
		size := sizes[player].(*components.Size)
		vel := velocities[player].(*components.Velocity)

		// Compute the player's center
		centerX := pos.X + size.Width/2
		centerY := pos.Y + size.Height/2

		// Determine the cell the player is in
		col := int(centerX / float64(cellSize))
		row := int(centerY / float64(cellSize))
		cell := mazeLayout.GetCell(col, row)

		// Check if out of mazeLayout bounds
		if col < 0 || col >= mazeLayout.Cols() || row < 0 || row >= mazeLayout.Rows() {
			vel.DX = 0
			vel.DY = 0
			continue
		}

		// Check collisions with walls based on velocity direction
		if vel.DY < 0 && crossedTopBoundary(pos, row, cellSize) { // Moving UP
			if cell.HasTopWall() {
				vel.DY = 0
				pos.Y = float64(row * cellSize)
			}
		}

		if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellSize) { // Moving RIGHT
			if cell.HasRightWall() {
				vel.DX = 0
				pos.X = float64((col+1)*cellSize) - size.Width
			}
		}

		if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellSize) { // Moving DOWN
			if cell.HasBottomWall() {
				vel.DY = 0
				pos.Y = float64((row+1)*cellSize) - size.Height
			}
		}

		if vel.DX < 0 && crossedLeftBoundary(pos, col, cellSize) { // Moving LEFT
			if cell.HasLeftWall() {
				vel.DX = 0
				pos.X = float64(col * cellSize)
			}
		}

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
