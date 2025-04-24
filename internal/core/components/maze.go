package components

import (
	"github.com/juanancid/maze-adventure/internal/engine/mazebuilder"
)

type Maze struct {
	Layout     mazebuilder.Layout
	CellWidth  int
	CellHeight int
}
