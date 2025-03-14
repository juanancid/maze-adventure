package levels

import "github.com/juanancid/maze-adventure/internal/ecs"

func CreateLevelWorld(level int) *ecs.World {
	world := ecs.NewWorld()

	config, ok := levelConfigs[level]
	if !ok {
		config = levelConfigs[-1] // use default config if not found.
	}

	createPlayer(world, config.playerSize, config.cellSize)
	createMaze(world, config.mazeWidth, config.mazeHeight, config.cellSize)
	createExit(world, config.mazeWidth-1, config.mazeHeight-1, config.cellSize)

	return world
}
