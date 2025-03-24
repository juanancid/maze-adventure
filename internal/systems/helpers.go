package systems

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/components"
	"github.com/juanancid/maze-adventure/internal/entities"
)

func getMaze(world *entities.World) (*components.Maze, bool) {
	mazeType := reflect.TypeOf(&components.Maze{})
	mazes := world.Query(mazeType)
	if len(mazes) == 0 {
		return nil, false
	}
	comp := world.GetComponent(mazes[0], mazeType).(*components.Maze)
	return comp, true
}
