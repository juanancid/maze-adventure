package levels

import "github.com/juanancid/maze-adventure/internal/ecs"

func CreateLevelWorld() *ecs.World {
	world := ecs.NewWorld()

	const (
		mazeWidth  = 16
		mazeHeight = 10
		cellSize   = 20
	)

	createPlayer(world, cellSize)
	createMaze(world, mazeWidth, mazeHeight, cellSize)

	return world
}
