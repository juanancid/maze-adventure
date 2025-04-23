package levels

import "github.com/juanancid/maze-adventure/internal/gameplay/levels/definitions"

// Level represents a game level configuration.
type Level struct {
	Number int
	Maze   definitions.MazeConfig
	Player definitions.PlayerConfig
	Exit   definitions.ExitConfig
}

// EmptyLevel is a placeholder for an empty level.
var EmptyLevel = Level{}
