package levels

import "github.com/juanancid/maze-adventure/internal/gameplay/levels/definitions"

// Config represents a game level configuration.
type Config struct {
	Number       int
	Maze         definitions.MazeConfig
	Player       definitions.PlayerConfig
	Exit         definitions.ExitConfig
	Collectibles definitions.Collectibles
}

// EmptyLevel is a placeholder for an empty level.
var EmptyLevel = Config{}
