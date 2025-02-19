package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
)

// MazeCollisionSystem ensures entities do not pass through maze walls.
type MazeCollisionSystem struct{}

func (mcs *MazeCollisionSystem) Update(w *ecs.World) {
	players := w.GetComponents(reflect.TypeOf(&components.InputControlled{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))
	velocities := w.GetComponents(reflect.TypeOf(&components.Velocity{}))
	mazes := w.GetComponents(reflect.TypeOf(&components.Maze{}))

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

			// Check if out of maze bounds
			if col < 0 || col >= maze.Cols() || row < 0 || row >= maze.Rows() {
				vel.DX = 0
				vel.DY = 0
				continue
			}

			// Check collisions with walls based on velocity direction
			if vel.DY < 0 && collidesMovingUp(mazeComp, pos, col, row) { // Moving UP
				vel.DY = 0
				pos.Y = float64(row*cellSize) + 1
			}
			if vel.DX > 0 && collidesMovingRight(mazeComp, pos, col, row) { // Moving RIGHT
				vel.DX = 0
				pos.X = float64((col+1)*cellSize) - size.Width - 1
			}
			if vel.DY > 0 && collidesMovingDown(mazeComp, pos, col, row) { // Moving DOWN
				vel.DY = 0
				pos.Y = float64((row+1)*cellSize) - size.Height - 1
			}
			if vel.DX < 0 && collidesMovingLeft(mazeComp, pos, col, row) { // Moving LEFT
				vel.DX = 0
				pos.X = float64(col*cellSize) + 1
			}
		}
	}
}

func collidesMovingUp(mazeComp *components.Maze, pos *components.Position, col, row int) bool {
	maze := mazeComp.Maze
	cellSize := mazeComp.CellSize

	cell := maze.GetCell(col, row)

	return cell.Walls[0] && pos.Y < float64(row*cellSize)
}

func collidesMovingRight(mazeComp *components.Maze, pos *components.Position, col, row int) bool {
	maze := mazeComp.Maze
	cellSize := mazeComp.CellSize

	cell := maze.GetCell(col, row)

	return cell.Walls[1] && pos.X > float64((col+1)*cellSize)
}

func collidesMovingDown(mazeComp *components.Maze, pos *components.Position, col, row int) bool {
	maze := mazeComp.Maze
	cellSize := mazeComp.CellSize

	cell := maze.GetCell(col, row)

	return cell.Walls[2] && pos.Y > float64((row+1)*cellSize)
}

func collidesMovingLeft(mazeComp *components.Maze, pos *components.Position, col, row int) bool {
	maze := mazeComp.Maze
	cellSize := mazeComp.CellSize

	cell := maze.GetCell(col, row)

	return cell.Walls[3] && pos.X < float64(col*cellSize)
}
