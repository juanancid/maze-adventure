package definitions

// LevelConfig represents the configuration for a game level
type LevelConfig struct {
	Maze         MazeConfig
	Player       PlayerConfig
	Exit         ExitConfig
	Collectibles Collectibles
}

// MazeConfig defines the maze dimensions
type MazeConfig struct {
	Cols int
	Rows int
}

// PlayerConfig defines the player properties
type PlayerConfig struct {
	Size int
}

// ExitConfig defines the exit properties
type ExitConfig struct {
	Position Coordinate
	Size     int
}

// Coordinate represents a position in the maze
type Coordinate struct {
	X int
	Y int
}

type Collectibles struct {
	Number int
	Size   int
	Value  int
}
