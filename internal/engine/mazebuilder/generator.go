package mazebuilder

import (
	"math/rand"
	"time"

	"github.com/juanancid/maze-adventure/internal/core/components"
)

func newMazeLayout(cols, rows int, extraConnectionChance float64) components.Layout {
	bGrid := initializeBuilderGrid(cols, rows)

	startCol, startRow := 0, 0
	carveMazePaths(startCol, startRow, cols, rows, bGrid)
	addExtraConnections(bGrid, extraConnectionChance)

	return convertBuilderGridToLayout(bGrid, cols, rows)
}

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

func addExtraConnections(grid builderGrid, chance float64) {
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}

	rows := len(grid)
	cols := len(grid[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			for dir := 0; dir < 4; dir++ {
				nx := x + dx[dir]
				ny := y + dy[dir]

				if inBounds(nx, ny, cols, rows) && r.Float64() < chance {
					c := grid[y][x]
					n := grid[ny][nx]
					if wallsBetween(c, n) { // Only break walls that still exist
						removeWall(c, n, dir)
					}
				}
			}
		}
	}
}

func wallsBetween(a, b *builderCell) bool {
	// They must be adjacent
	dx := a.x - b.x
	dy := a.y - b.y

	if dx == 1 {
		return a.walls[3] && b.walls[1]
	} // left
	if dx == -1 {
		return a.walls[1] && b.walls[3]
	} // right
	if dy == 1 {
		return a.walls[0] && b.walls[2]
	} // up
	if dy == -1 {
		return a.walls[2] && b.walls[0]
	} // down

	return false
}

func convertBuilderGridToLayout(grid builderGrid, cols, rows int) components.Layout {
	finalGrid := make([][]components.Cell, rows)
	for y := range grid {
		finalGrid[y] = make([]components.Cell, cols)
		for x := range grid[y] {
			finalGrid[y][x] = components.NewRegularCell(grid[y][x].walls)
		}
	}

	return components.NewLayout(cols, rows, finalGrid)
}
