package levels

import "github.com/juanancid/maze-adventure/internal/ecs"

func CreateLevel(world *ecs.World) {
	const (
		mazeWidth  = 16
		mazeHeight = 10
		cellSize   = 20
	)

	createPlayer(world, cellSize)
	createMaze(world, mazeWidth, mazeHeight, cellSize)
}
