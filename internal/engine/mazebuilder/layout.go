package mazebuilder

// Layout represents a maze with a 2D grid of cells.
type Layout struct {
	cols int
	rows int
	grid [][]Cell
}

// NewMazeLayout creates a new maze with the given width and height.
func NewMazeLayout(cols, rows int) Layout {
	bGrid := initializeBuilderGrid(cols, rows)

	startCol, startRow := 0, 0
	carveMazePaths(startCol, startRow, cols, rows, bGrid)

	return convertBuilderGridToLayout(bGrid, cols, rows)
}
