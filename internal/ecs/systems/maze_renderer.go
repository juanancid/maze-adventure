package systems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/juanancid/maze-adventure/internal/ecs"
)

type MazeRenderer struct{}

func (r *MazeRenderer) Draw(w *ecs.World, screen *ebiten.Image) {
	mazeLayout := w.Maze().Layout
	cellSize := w.Maze().CellSize

	wallColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	// Iterate over each cell and draw its walls.
	for row := 0; row < mazeLayout.Rows(); row++ {
		for col := 0; col < mazeLayout.Cols(); col++ {
			cell := mazeLayout.GetCell(col, row)

			// Calculate pixel coordinates.
			x1 := float64(col*cellSize) + 1
			y1 := float64(row*cellSize) + 1
			x2 := float64((col+1)*cellSize) + 1
			y2 := float64((row+1)*cellSize) + 1

			// Draw top wall.
			if cell.Walls[0] {
				vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y1), 1, wallColor, false)
			}
			// Draw right wall.
			if cell.Walls[1] {
				vector.StrokeLine(screen, float32(x2-1), float32(y1), float32(x2-1), float32(y2), 1, wallColor, false)
			}
			// Draw bottom wall.
			if cell.Walls[2] {
				vector.StrokeLine(screen, float32(x2), float32(y2-1), float32(x1), float32(y2-1), 1, wallColor, false)
			}
			// Draw left wall.
			if cell.Walls[3] {
				vector.StrokeLine(screen, float32(x1), float32(y2), float32(x1), float32(y1), 1, wallColor, false)
			}
		}
	}
}
