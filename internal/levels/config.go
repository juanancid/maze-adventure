package levels

type levelConfig struct {
	mazeWidth  int
	mazeHeight int
	cellSize   int
	playerSize int
}

var levelConfigs = map[int]levelConfig{
	0: {mazeWidth: 8, mazeHeight: 5, cellSize: 40, playerSize: 12},
	1: {mazeWidth: 16, mazeHeight: 10, cellSize: 20, playerSize: 12},

	// Default level
	-1: {mazeWidth: 32, mazeHeight: 20, cellSize: 10, playerSize: 6},
}
