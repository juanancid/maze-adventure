package mazebuilder

import (
	"math/rand"
)

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
