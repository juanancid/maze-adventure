package mazebuilder

import (
	"math/rand"
)

// MazeLayout represents a maze with a 2D grid of cells.
type MazeLayout struct {
	cols int
	rows int
	grid MazeGrid
}

// MazeGrid represents a 2D grid of cells.
type MazeGrid []MazeRow

// MazeRow represents a row of cells.
type MazeRow []*MazeCell

// MazeCell represents a cell in the maze.
type MazeCell struct {
	x, y    int
	visited bool
	// Walls: [top, right, bottom, left]
	Walls [4]bool
}

// HasTopWall returns true if the cell has a top wall.
func (c *MazeCell) HasTopWall() bool {
	return c.Walls[0]
}

// HasRightWall returns true if the cell has a right wall.
func (c *MazeCell) HasRightWall() bool {
	return c.Walls[1]
}

// HasBottomWall returns true if the cell has a bottom wall.
func (c *MazeCell) HasBottomWall() bool {
	return c.Walls[2]
}

// HasLeftWall returns true if the cell has a left wall.
func (c *MazeCell) HasLeftWall() bool {
	return c.Walls[3]
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
	return m.grid[y][x]
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

// NewMazeLayout creates a new maze with the given width and height.
func NewMazeLayout(cols, rows int) MazeLayout {
	grid := initializeGrid(cols, rows)

	startCol, startRow := 0, 0
	maze := carveMaze(startCol, startRow, cols, rows, grid)

	return maze
}

func initializeGrid(cols, rows int) MazeGrid {
	grid := make(MazeGrid, rows)

	for row := 0; row < rows; row++ {
		grid[row] = make(MazeRow, cols)

		for col := 0; col < cols; col++ {
			grid[row][col] = &MazeCell{
				x:       col,
				y:       row,
				visited: false,
				Walls:   [4]bool{true, true, true, true},
			}
		}
	}

	return grid
}

func carveMaze(startCol, startRow int, cols, rows int, grid MazeGrid) MazeLayout {
	var (
		dx = [4]int{0, 1, 0, -1}
		dy = [4]int{-1, 0, 1, 0}
	)

	stack := MazeRow{}

	start := grid[startRow][startCol]
	start.visited = true
	stack = append(stack, start)

	for len(stack) > 0 {
		// Peek at the top of the stack.
		current := stack[len(stack)-1]

		// Collect all unvisited neighbors.
		var neighbors MazeRow
		var directions []int
		for dir := 0; dir < 4; dir++ {
			nx := current.x + dx[dir]
			ny := current.y + dy[dir]
			if inBounds(nx, ny, cols, rows) && !grid[ny][nx].visited {
				neighbors = append(neighbors, grid[ny][nx])
				directions = append(directions, dir)
			}
		}

		if len(neighbors) > 0 {
			// Randomly choose one unvisited neighbor.
			idx := rand.Intn(len(neighbors))
			neighbor := neighbors[idx]
			dir := directions[idx]

			// Carve the path between current and neighbor.
			removeWall(current, neighbor, dir)

			// Mark neighbor as visited and push it onto the stack.
			neighbor.visited = true
			stack = append(stack, neighbor)
		} else {
			// Backtrack if no unvisited neighbors.
			stack = stack[:len(stack)-1]
		}
	}

	return MazeLayout{
		cols: cols,
		rows: rows,
		grid: grid,
	}
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func removeWall(current, neighbor *MazeCell, dir int) {
	current.Walls[dir] = false
	neighbor.Walls[(dir+2)%4] = false // Remove the opposite wall in neighbor.
}
