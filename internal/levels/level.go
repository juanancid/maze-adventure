package levels

import "github.com/juanancid/maze-adventure/internal/ecs"

func CreateLevel(world *ecs.World) {
	const (
		mazeWidth  = 10
		mazeHeight = 10
	)

	createPlayer(world)
	createMaze(world, mazeWidth, mazeHeight)
}
