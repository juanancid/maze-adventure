package maze

import (
	"math/rand"
)

// Maze represents a maze with a 2D grid of cells.
type Maze struct {
	Width  int
	Height int
	Grid   Grid
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

// New creates a new maze with the given width and height.
func New(width, height int) Maze {
	startX, startY := 0, 0
	maze := initMaze(width, height)

	generateMaze(startX, startY, maze)

	return maze
}

func initMaze(width, height int) Maze {
	grid := make(Grid, height)

	for y := 0; y < height; y++ {
		grid[y] = make(CellRow, width)

		for x := 0; x < width; x++ {
			grid[y][x] = &Cell{
				x:       x,
				y:       y,
				visited: false,
				Walls:   [4]bool{true, true, true, true},
			}
		}
	}

	return Maze{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}

// generateMaze creates a maze using an iterative DFS algorithm.
func generateMaze(startX, startY int, maze Maze) {
	var (
		dx = [4]int{0, 1, 0, -1}
		dy = [4]int{-1, 0, 1, 0}
	)

	stack := CellRow{}

	start := maze.Grid[startX][startY]
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
			if inBounds(nx, ny, maze.Width, maze.Height) && !maze.Grid[ny][nx].visited {
				neighbors = append(neighbors, maze.Grid[ny][nx])
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
