package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
)

// MazeCollisionSystem ensures entities do not pass through maze walls.
type MazeCollisionSystem struct{}

func (mcs *MazeCollisionSystem) Update(w *ecs.World) {
	mazes := w.GetComponents(reflect.TypeOf(&components.Maze{}))
	players := w.GetComponents(reflect.TypeOf(&components.InputControlled{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))
	velocities := w.GetComponents(reflect.TypeOf(&components.Velocity{}))

	for mazeEntity := range mazes {
		mazeComp := mazes[mazeEntity].(*components.Maze)
		maze := mazeComp.Maze
		cellSize := mazeComp.CellSize

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
			cell := maze.GetCell(col, row)

			// Check if out of maze bounds
			if col < 0 || col >= maze.Cols() || row < 0 || row >= maze.Rows() {
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
				if crossedLeftBoundary(pos, col, cellSize) && maze.GetCellAbove(col, row).HasLeftWall() ||
					crossedRightBoundary(pos, size, col, cellSize) && maze.GetCellAbove(col, row).HasRightWall() {
					vel.DY = 0
					pos.Y = float64(row * cellSize)
				}
			}

			if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellSize) && col < maze.Cols()-1 { // Moving RIGHT
				if crossedTopBoundary(pos, row, cellSize) && maze.GetCellRight(col, row).HasTopWall() ||
					crossedBottomBoundary(pos, size, row, cellSize) && maze.GetCellRight(col, row).HasBottomWall() {
					vel.DX = 0
					pos.X = float64((col+1)*cellSize) - size.Width
				}
			}

			if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellSize) && row < maze.Rows()-1 { // Moving DOWN
				if crossedLeftBoundary(pos, col, cellSize) && maze.GetCellBelow(col, row).HasLeftWall() ||
					crossedRightBoundary(pos, size, col, cellSize) && maze.GetCellBelow(col, row).HasRightWall() {
					vel.DY = 0
					pos.Y = float64((row+1)*cellSize) - size.Height
				}
			}

			if vel.DX < 0 && crossedLeftBoundary(pos, col, cellSize) && col > 0 { // Moving LEFT
				if crossedTopBoundary(pos, row, cellSize) && maze.GetCellLeft(col, row).HasTopWall() ||
					crossedBottomBoundary(pos, size, row, cellSize) && maze.GetCellLeft(col, row).HasBottomWall() {
					vel.DX = 0
					pos.X = float64(col * cellSize)
				}
			}

			// Check collisions with other cells walls based on velocity direction
			if vel.DY < 0 && crossedTopBoundary(pos, row, cellSize) { // Moving UP
				if col > 0 && crossedLeftBoundary(pos, col, cellSize) && maze.GetCellLeft(col, row).HasTopWall() ||
					col < maze.Cols()-1 && crossedRightBoundary(pos, size, col, cellSize) && maze.GetCellRight(col, row).HasTopWall() {
					vel.DY = 0
					pos.Y = float64(row * cellSize)
				}
			}

			if vel.DX > 0 && crossedRightBoundary(pos, size, col, cellSize) { // Moving RIGHT
				if row > 0 && crossedTopBoundary(pos, row, cellSize) && maze.GetCellAbove(col, row).HasRightWall() ||
					row < maze.Rows()-1 && crossedBottomBoundary(pos, size, row, cellSize) && maze.GetCellBelow(col, row).HasRightWall() {
					vel.DX = 0
					pos.X = float64((col+1)*cellSize) - size.Width
				}
			}

			if vel.DY > 0 && crossedBottomBoundary(pos, size, row, cellSize) { // Moving DOWN
				if col > 0 && crossedLeftBoundary(pos, col, cellSize) && maze.GetCellLeft(col, row).HasBottomWall() ||
					col < maze.Cols()-1 && crossedRightBoundary(pos, size, col, cellSize) && maze.GetCellRight(col, row).HasBottomWall() {
					vel.DY = 0
					pos.Y = float64((row+1)*cellSize) - size.Height
				}
			}

			if vel.DX < 0 && crossedLeftBoundary(pos, col, cellSize) { // Moving LEFT
				if row > 0 && crossedTopBoundary(pos, row, cellSize) && maze.GetCellAbove(col, row).HasLeftWall() ||
					row < maze.Rows()-1 && crossedBottomBoundary(pos, size, row, cellSize) && maze.GetCellBelow(col, row).HasLeftWall() {
					vel.DX = 0
					pos.X = float64(col * cellSize)
				}
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
