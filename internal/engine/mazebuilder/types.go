package mazebuilder

// MazeLayout represents a maze with a 2D grid of cells.
type MazeLayout struct {
	cols int
	rows int
	grid [][]MazeCell
}

// MazeCell represents a cell in the maze.
type MazeCell struct {
	walls [4]bool
}

// HasTopWall returns true if the cell has a top wall.
func (c *MazeCell) HasTopWall() bool {
	return c.walls[0]
}

// HasRightWall returns true if the cell has a right wall.
func (c *MazeCell) HasRightWall() bool {
	return c.walls[1]
}

// HasBottomWall returns true if the cell has a bottom wall.
func (c *MazeCell) HasBottomWall() bool {
	return c.walls[2]
}

// HasLeftWall returns true if the cell has a left wall.
func (c *MazeCell) HasLeftWall() bool {
	return c.walls[3]
}

// Cols returns the number of columns in the maze.
func (m MazeLayout) Cols() int {
	return m.cols
}

// Rows returns the number of rows in the maze.
func (m MazeLayout) Rows() int {
	return m.rows
}

// GetCell returns the cell at the given coordinates.
func (m MazeLayout) GetCell(x, y int) *MazeCell {
	return &m.grid[y][x]
}

// GetCellAbove returns the cell above the given coordinates.
func (m MazeLayout) GetCellAbove(x, y int) *MazeCell {
	return m.GetCell(x, y-1)
}

// GetCellRight returns the cell to the right of the given coordinates.
func (m MazeLayout) GetCellRight(x, y int) *MazeCell {
	return m.GetCell(x+1, y)
}

// GetCellBelow returns the cell below the given coordinates.
func (m MazeLayout) GetCellBelow(x, y int) *MazeCell {
	return m.GetCell(x, y+1)
}

// GetCellLeft returns the cell to the left of the given coordinates.
func (m MazeLayout) GetCellLeft(x, y int) *MazeCell {
	return m.GetCell(x-1, y)
}
