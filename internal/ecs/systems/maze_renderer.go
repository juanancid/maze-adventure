package systems

import (
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
)

type MazeRenderer struct{}

func (r *MazeRenderer) Draw(w *ecs.World, screen *ebiten.Image) {
	wallColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	mazes := w.GetComponents(reflect.TypeOf(&components.Maze{}))

	for _, maze := range mazes {
		m := maze.(*components.Maze).Maze
		cellSize := maze.(*components.Maze).CellSize

		// Iterate over each cell and draw its walls.
		for row := 0; row < m.Rows(); row++ {
			for col := 0; col < m.Cols(); col++ {
				cell := m.GetCell(col, row)

				// Calculate pixel coordinates.
				x1 := float64(col*cellSize) + 1
				y1 := float64(row*cellSize) + 1
				x2 := float64((col+1)*cellSize) + 1
				y2 := float64((row+1)*cellSize) + 1

				// Draw top wall.
				if cell.Walls[0] {
					ebitenutil.DrawLine(screen, x1, y1, x2, y1, wallColor)
				}
				// Draw right wall.
				if cell.Walls[1] {
					ebitenutil.DrawLine(screen, x2, y1, x2, y2, wallColor)
				}
				// Draw bottom wall.
				if cell.Walls[2] {
					ebitenutil.DrawLine(screen, x2, y2, x1, y2, wallColor)
				}
				// Draw left wall.
				if cell.Walls[3] {
					ebitenutil.DrawLine(screen, x1, y2, x1, y1, wallColor)
				}
			}
		}
	}
}
