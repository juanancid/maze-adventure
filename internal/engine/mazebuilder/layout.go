package mazebuilder

// MazeLayout represents a maze with a 2D grid of cells.
type MazeLayout struct {
	cols int
	rows int
	grid [][]MazeCell
}

// NewMazeLayout creates a new maze with the given width and height.
func NewMazeLayout(cols, rows int) MazeLayout {
	grid := initializeGrid(cols, rows)

	startCol, startRow := 0, 0
	carveMaze(startCol, startRow, cols, rows, grid)

	return convertToLayout(grid, cols, rows)
}
