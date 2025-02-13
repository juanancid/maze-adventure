package maze

import (
	"math/rand"
	"time"
)

// Maze represents a maze with a 2D grid of cells.
type Maze struct {
	Width  int
	Height int
	Grid   Grid
}

// Grid represents a 2D grid of cells.
type Grid []CellColumn

// CellColumn represents a column of cells in the maze.
type CellColumn []*Cell

// Cell represents a cell in the maze.
type Cell struct {
	x, y    int
	visited bool
	// Walls: [top, right, bottom, left]
	Walls [4]bool
}

// Direction vectors: 0 = top, 1 = right, 2 = bottom, 3 = left.
var dx = [4]int{0, 1, 0, -1}
var dy = [4]int{-1, 0, 1, 0}

func New(width, height int) Maze {
	grid := initGrid(width, height)
	generateMazeDFS(0, 0, width, height, grid)

	return Maze{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}

// initGrid initializes the Grid with cells with all Walls intact.
func initGrid(width, height int) Grid {
	grid := make(Grid, width)
	for x := 0; x < width; x++ {
		grid[x] = make(CellColumn, height)
		for y := 0; y < height; y++ {
			grid[x][y] = &Cell{
				x:       x,
				y:       y,
				visited: false,
				Walls:   [4]bool{true, true, true, true},
			}
		}
	}

	return grid
}

// inBounds checks if the given coordinates are within the Grid.
func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

// removeWall removes the wall between two adjacent cells.
// 'dir' is the direction from the current cell to the neighbor.
func removeWall(current, neighbor *Cell, dir int) {
	current.Walls[dir] = false
	neighbor.Walls[(dir+2)%4] = false // Remove the opposite wall in neighbor.
}

// generateMazeDFS creates a maze using an iterative DFS algorithm.
func generateMazeDFS(startX, startY, mazeWidth, mazeHeight int, grid Grid) {
	stack := CellColumn{}

	start := grid[startX][startY]
	start.visited = true
	stack = append(stack, start)

	// Seed the random generator.
	rand.Seed(time.Now().UnixNano())

	for len(stack) > 0 {
		// Peek at the top of the stack.
		current := stack[len(stack)-1]

		// Collect all unvisited neighbors.
		var neighbors CellColumn
		var directions []int
		for dir := 0; dir < 4; dir++ {
			nx := current.x + dx[dir]
			ny := current.y + dy[dir]
			if inBounds(nx, ny, mazeWidth, mazeHeight) && !grid[nx][ny].visited {
				neighbors = append(neighbors, grid[nx][ny])
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
