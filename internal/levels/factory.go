package levels

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
	"github.com/juanancid/maze-adventure/internal/maze"
	"github.com/juanancid/maze-adventure/internal/utils"
)

const (
	playerSpriteFile = "internal/assets/images/player.png"
)

func createPlayer(world *ecs.World) ecs.Entity {
	player := world.NewEntity()

	world.AddComponent(player, &components.Position{X: 100, Y: 100})
	world.AddComponent(player, &components.Velocity{})
	world.AddComponent(player, &components.Size{Width: 12, Height: 12})

	world.AddComponent(player, &components.InputControlled{
		MoveLeftKey:  ebiten.KeyLeft,
		MoveRightKey: ebiten.KeyRight,
		MoveUpKey:    ebiten.KeyUp,
		MoveDownKey:  ebiten.KeyDown,
	})

	playerSprite := utils.MustLoadSprite(playerSpriteFile)
	world.AddComponent(player, &components.Sprite{Image: playerSprite})

	return player
}

func createMaze(world *ecs.World, mazeWidth, mazeHeight int, cellSize int) ecs.Entity {
	mazeEntity := world.NewEntity()
	world.AddComponent(mazeEntity, &components.Maze{
		Maze:     maze.New(mazeWidth, mazeHeight),
		CellSize: cellSize,
	})

	return mazeEntity
}
