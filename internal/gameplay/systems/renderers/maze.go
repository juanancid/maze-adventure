package renderers

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/engine/config"
)

type Maze struct{}

func NewMaze() Maze {
	return Maze{}
}

func (r Maze) Draw(w *entities.World, screen *ebiten.Image) {
	maze, ok := queries.GetMazeComponent(w)
	if !ok {
		return
	}

	mazeLayout := maze.Layout
	cellWidth := maze.CellWidth
	cellHeight := maze.CellHeight

	// Fill the entire maze area with background color
	bgColor := color.RGBA{R: 0x12, G: 0x18, B: 0x21, A: 0xFF}
	screen.Fill(bgColor)

	wallColor := color.RGBA{R: 0x36, G: 0x9b, B: 0x48, A: 0xFF}

	// Iterate over each cell and draw its walls.
	for row := 0; row < mazeLayout.Rows(); row++ {
		for col := 0; col < mazeLayout.Cols(); col++ {
			cell := mazeLayout.GetCell(col, row)

			// Calculate pixel coordinates.
			x1 := float64(col*cellWidth) + 1
			y1 := float64(row*cellHeight+config.HudHeight) + 1
			x2 := float64((col+1)*cellWidth) + 1
			y2 := float64((row+1)*cellHeight+config.HudHeight) + 1

			if cell.HasTopWall() {
				vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y1), 1, wallColor, false)
			}

			if cell.HasRightWall() {
				vector.StrokeLine(screen, float32(x2-1), float32(y1), float32(x2-1), float32(y2), 1, wallColor, false)
			}

			if cell.HasBottomWall() {
				vector.StrokeLine(screen, float32(x2), float32(y2-1), float32(x1), float32(y2-1), 1, wallColor, false)
			}

			if cell.HasLeftWall() {
				vector.StrokeLine(screen, float32(x1), float32(y2), float32(x1), float32(y1), 1, wallColor, false)
			}
		}
	}
}
