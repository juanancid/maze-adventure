package queries

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
)

func GetMazeComponent(world *entities.World) (*components.Maze, bool) {
	mazeType := reflect.TypeOf(&components.Maze{})
	mazes := world.Query(mazeType)
	if len(mazes) == 0 {
		return nil, false
	}

	comp := world.GetComponent(mazes[0], mazeType).(*components.Maze)
	return comp, true
}

func GetPlayerEntity(world *entities.World) (entities.Entity, bool) {
	playerType := reflect.TypeOf(&components.InputControlled{})
	players := world.Query(playerType)
	if len(players) == 0 {
		return 0, false
	}

	return players[0], true
}

func GetExitEntity(world *entities.World) (entities.Entity, bool) {
	exitType := reflect.TypeOf(&components.Exit{})
	exits := world.Query(exitType)
	if len(exits) == 0 {
		return 0, false
	}

	return exits[0], true
}
