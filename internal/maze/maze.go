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
type Grid []CellRow

// CellRow represents a column of cells in the maze.
type CellRow []*Cell

// Cell represents a cell in the maze.
type Cell struct {
	x, y    int
	visited bool
	// Walls: [top, right, bottom, left]
	Walls [4]bool
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
	startCol, startRow := 0, 0
	maze := initMaze(cols, rows)

	generateMaze(startCol, startRow, maze)

	return maze
}

func initMaze(cols, rows int) Maze {
	grid := make(Grid, rows)

	for row := 0; row < rows; row++ {
		grid[row] = make(CellRow, cols)

		for col := 0; col < cols; col++ {
			grid[row][col] = &Cell{
				x:       col,
				y:       row,
				visited: false,
				Walls:   [4]bool{true, true, true, true},
			}
		}
	}

	return Maze{
		cols: cols,
		rows: rows,
		grid: grid,
	}
}

// generateMaze creates a maze using an iterative DFS algorithm.
func generateMaze(startCol, startRow int, maze Maze) {
	var (
		dx = [4]int{0, 1, 0, -1}
		dy = [4]int{-1, 0, 1, 0}
	)

	stack := CellRow{}

	start := maze.grid[startCol][startRow]
	start.visited = true
	stack = append(stack, start)

	for len(stack) > 0 {
		// Peek at the top of the stack.
		current := stack[len(stack)-1]

		// Collect all unvisited neighbors.
		var neighbors CellRow
		var directions []int
		for dir := 0; dir < 4; dir++ {
			nx := current.x + dx[dir]
			ny := current.y + dy[dir]
			if inBounds(nx, ny, maze.cols, maze.rows) && !maze.grid[ny][nx].visited {
				neighbors = append(neighbors, maze.grid[ny][nx])
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
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

// removeWall removes the wall between two adjacent cells.
// 'dir' is the direction from the current cell to the neighbor.
func removeWall(current, neighbor *Cell, dir int) {
	current.Walls[dir] = false
	neighbor.Walls[(dir+2)%4] = false // Remove the opposite wall in neighbor.
}
