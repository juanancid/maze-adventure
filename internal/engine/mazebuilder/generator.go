package mazebuilder

import (
	"math/rand"
)

type builderCell struct {
	x, y    int
	visited bool
	walls   [4]bool
}

type builderGrid [][]*builderCell

func initializeBuilderGrid(cols, rows int) builderGrid {
	grid := make(builderGrid, rows)

	for row := 0; row < rows; row++ {
		grid[row] = make([]*builderCell, cols)

		for col := 0; col < cols; col++ {
			grid[row][col] = &builderCell{
				x:       col,
				y:       row,
				visited: false,
				walls:   [4]bool{true, true, true, true},
			}
		}
	}

	return grid
}

func carveMazePaths(startCol, startRow int, cols, rows int, grid builderGrid) {
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}

	stack := []*builderCell{grid[startRow][startCol]}

	start := grid[startRow][startCol]
	start.visited = true
	stack = append(stack, start)

	for len(stack) > 0 {
		// Peek at the top of the stack.
		current := stack[len(stack)-1]

		// Collect all unvisited neighbors.
		var neighbors []*builderCell
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
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func removeWall(current, neighbor *builderCell, dir int) {
	current.walls[dir] = false
	neighbor.walls[(dir+2)%4] = false // Remove the opposite wall in neighbor.
}

func convertBuilderGridToLayout(grid builderGrid, cols, rows int) Layout {
	finalGrid := make([][]Cell, rows)
	for y := range grid {
		finalGrid[y] = make([]Cell, cols)
		for x := range grid[y] {
			finalGrid[y][x] = Cell{walls: grid[y][x].walls}
		}
	}
	return Layout{cols: cols, rows: rows, grid: finalGrid}
}
