package components

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/maze"
)

type Position struct {
	X, Y float64
}

type Velocity struct {
	DX, DY float64
}

type Size struct {
	Width, Height float64
}

type InputControlled struct {
	MoveLeftKey  ebiten.Key
	MoveRightKey ebiten.Key
	MoveUpKey    ebiten.Key
	MoveDownKey  ebiten.Key
}

type Sprite struct {
	Image *ebiten.Image
}

type Maze struct {
	Maze     maze.Maze
	CellSize int
}
