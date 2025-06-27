package components

type cellType int

const (
	// cellTypeRegular is a standard wall cell that doesn't affect the player
	cellTypeRegular cellType = iota
	// cellTypeDeadly is a cell that kills the player on contact
	cellTypeDeadly
	// cellTypeFreezing is a cell that freezes the player temporarily
	cellTypeFreezing
)

// Cell represents a cell in the maze.
type Cell struct {
	walls [4]bool
	Type  cellType
}

func NewRegularCell(walls [4]bool) Cell {
	return newCellWithWalls(walls, cellTypeRegular)
}

func NewDeadlyCell(walls [4]bool) Cell {
	return newCellWithWalls(walls, cellTypeDeadly)
}

func NewFreezingCell(walls [4]bool) Cell {
	return newCellWithWalls(walls, cellTypeFreezing)
}

func newCellWithWalls(walls [4]bool, cellType cellType) Cell {
	return Cell{
		walls: walls,
		Type:  cellType,
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

// IsRegular returns true if the cell is a regular cell
func (c Cell) IsRegular() bool {
	return c.Type == cellTypeRegular
}

// IsDeadly returns true if the cell is a deadly cell
func (c Cell) IsDeadly() bool {
	return c.Type == cellTypeDeadly
}

// IsFreezing returns true if the cell is a freezing cell
func (c Cell) IsFreezing() bool {
	return c.Type == cellTypeFreezing
}

func (c Cell) GetWalls() [4]bool {
	return c.walls
}

// Cols function returns the number of columns in the maze.
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
