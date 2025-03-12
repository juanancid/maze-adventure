package levels

import "github.com/juanancid/maze-adventure/internal/ecs"

func CreateLevelWorld(level int) *ecs.World {
	world := ecs.NewWorld()

	switch level {
	case 0:
		createLevelWorld0(world)
	case 1:
		createLevelWorld1(world)
	default:
		createLevelWorldDefault(world)
	}
	return world
}

func createLevelWorld0(world *ecs.World) {
	const (
		mazeWidth  = 8
		mazeHeight = 5
		cellSize   = 40
		playerSize = 12
	)

	createPlayer(world, playerSize, cellSize)
	createMaze(world, mazeWidth, mazeHeight, cellSize)
	createExit(world, mazeWidth-1, mazeHeight-1, cellSize)
}

func createLevelWorld1(world *ecs.World) {
	const (
		mazeWidth  = 16
		mazeHeight = 10
		cellSize   = 20
		playerSize = 12
	)

	createPlayer(world, playerSize, cellSize)
	createMaze(world, mazeWidth, mazeHeight, cellSize)
	createExit(world, mazeWidth-1, mazeHeight-1, cellSize)
}

func createLevelWorldDefault(world *ecs.World) {
	const (
		mazeWidth  = 32
		mazeHeight = 20
		cellSize   = 10
		playerSize = 6
	)

	createPlayer(world, playerSize, cellSize)
	createMaze(world, mazeWidth, mazeHeight, cellSize)
	createExit(world, mazeWidth-1, mazeHeight-1, cellSize)
}
