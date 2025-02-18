package systems

import (
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
)

const (
	cellSize = 16 // size of each cell in pixels
)

type MazeRenderer struct{}

func (r *MazeRenderer) Draw(w *ecs.World, screen *ebiten.Image) {
	wallColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	mazes := w.GetComponents(reflect.TypeOf(&components.Maze{}))

	for _, maze := range mazes {
		m := maze.(*components.Maze).Maze
		// Iterate over each cell and draw its walls.
		for y := 0; y < m.Height; y++ {
			for x := 0; x < m.Width; x++ {
				cell := m.Grid[y][x]
				// Calculate pixel coordinates.
				x1 := float64(x*cellSize) + 1
				y1 := float64(y*cellSize) + 1
				x2 := float64((x+1)*cellSize) + 1
				y2 := float64((y+1)*cellSize) + 1

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
