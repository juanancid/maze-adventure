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

func createPlayer(world *ecs.World, cellSize int) ecs.Entity {
	const (
		playerWidth  = 12
		playerHeight = 12
	)

	player := world.NewEntity()

	world.AddComponent(player, &components.Size{Width: playerWidth, Height: playerHeight})
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})

	posX := float64(cellSize-playerWidth) / 2
	posY := float64(cellSize-playerHeight) / 2
	world.AddComponent(player, &components.Position{X: posX, Y: posY})

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
