package maze

import (
	"math/rand"
)

// Maze represents a maze with a 2D grid of cells.
type Maze struct {
	cols int
	rows int
	grid Grid
}

// Grid represents a 2D grid of cells.
type Grid []Row

// Row represents a row of cells.
type Row []*Cell

// Cell represents a cell in the maze.
type Cell struct {
	x, y    int
	visited bool
	// Walls: [top, right, bottom, left]
	Walls [4]bool
}

// HasTopWall returns true if the cell has a top wall.
func (c *Cell) HasTopWall() bool {
	return c.Walls[0]
}

// HasRightWall returns true if the cell has a right wall.
func (c *Cell) HasRightWall() bool {
	return c.Walls[1]
}

// HasBottomWall returns true if the cell has a bottom wall.
func (c *Cell) HasBottomWall() bool {
	return c.Walls[2]
}

// HasLeftWall returns true if the cell has a left wall.
func (c *Cell) HasLeftWall() bool {
	return c.Walls[3]
}

// Cols returns the number of columns in the maze.
func (m Maze) Cols() int {
	return m.cols
}

// Rows returns the number of rows in the maze.
func (m Maze) Rows() int {
	return m.rows
}

// GetCell returns the cell at the given coordinates.
func (m Maze) GetCell(x, y int) *Cell {
	return m.grid[y][x]
}

// New creates a new maze with the given width and height.
func New(cols, rows int) Maze {
	grid := initializeGrid(cols, rows)

	startCol, startRow := 0, 0
	maze := carveMaze(startCol, startRow, cols, rows, grid)

	return maze
}

func initializeGrid(cols, rows int) Grid {
	grid := make(Grid, rows)

	for row := 0; row < rows; row++ {
		grid[row] = make(Row, cols)

		for col := 0; col < cols; col++ {
			grid[row][col] = &Cell{
				x:       col,
				y:       row,
				visited: false,
				Walls:   [4]bool{true, true, true, true},
			}
		}
	}

	return grid
}

func carveMaze(startCol, startRow int, cols, rows int, grid Grid) Maze {
	var (
		dx = [4]int{0, 1, 0, -1}
		dy = [4]int{-1, 0, 1, 0}
	)

	stack := Row{}

	start := grid[startRow][startCol]
	start.visited = true
	stack = append(stack, start)

	for len(stack) > 0 {
		// Peek at the top of the stack.
		current := stack[len(stack)-1]

		// Collect all unvisited neighbors.
		var neighbors Row
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

	return Maze{
		cols: cols,
		rows: rows,
		grid: grid,
	}
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func removeWall(current, neighbor *Cell, dir int) {
	current.Walls[dir] = false
	neighbor.Walls[(dir+2)%4] = false // Remove the opposite wall in neighbor.
}
