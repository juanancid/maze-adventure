package components

// CellType represents the type of a cell in the maze
type CellType int

const (
	// CellTypeRegular is a standard wall cell that doesn't affect the player
	CellTypeRegular CellType = iota
	// CellTypeDeadly is a cell that kills the player on contact
	CellTypeDeadly
	// CellTypeFreezing is a cell that freezes the player temporarily
	CellTypeFreezing
)

// Cell represents a cell in the maze.
type Cell struct {
	walls [4]bool
	Type  CellType
}

func NewCell(walls [4]bool) Cell {
	return Cell{
		walls: walls,
		Type:  CellTypeRegular, // Default to regular cell type
	}
}

// HasTopWall returns true if the cell has a top wall.
func (c Cell) HasTopWall() bool {
	return c.walls[0]
}

// HasRightWall returns true if the cell has a right wall.
func (c Cell) HasRightWall() bool {
	return c.walls[1]
}

// HasBottomWall returns true if the cell has a bottom wall.
func (c Cell) HasBottomWall() bool {
	return c.walls[2]
}

// HasLeftWall returns true if the cell has a left wall.
func (c Cell) HasLeftWall() bool {
	return c.walls[3]
}

// SetType sets the type of the cell
func (c *Cell) SetType(cellType CellType) {
	c.Type = cellType
}

// GetType returns the type of the cell
func (c Cell) GetType() CellType {
	return c.Type
}

// Cols returns the number of columns in the maze.
func (m Layout) Cols() int {
	return m.cols
}

// Rows returns the number of rows in the maze.
func (m Layout) Rows() int {
	return m.rows
}

// GetCell returns the cell at the given coordinates.
func (m Layout) GetCell(x, y int) Cell {
	return m.grid[y][x]
}

func (m Layout) SetCell(x, y int, cell Cell) {
	m.grid[y][x] = cell
}

// GetCellAbove returns the cell above the given coordinates.
func (m Layout) GetCellAbove(x, y int) Cell {
	return m.GetCell(x, y-1)
}

// GetCellRight returns the cell to the right of the given coordinates.
func (m Layout) GetCellRight(x, y int) Cell {
	return m.GetCell(x+1, y)
}

// GetCellBelow returns the cell below the given coordinates.
func (m Layout) GetCellBelow(x, y int) Cell {
	return m.GetCell(x, y+1)
}

// GetCellLeft returns the cell to the left of the given coordinates.
func (m Layout) GetCellLeft(x, y int) Cell {
	return m.GetCell(x-1, y)
}
