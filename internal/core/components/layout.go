package components

// Layout represents a maze with a 2D grid of cells.
type Layout struct {
	cols int
	rows int
	grid [][]Cell
}

func NewLayout(cols, rows int, grid [][]Cell) Layout {
	return Layout{
		cols: cols,
		rows: rows,
		grid: grid,
	}
}
