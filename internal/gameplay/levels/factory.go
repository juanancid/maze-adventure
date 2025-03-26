package levels

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/layout"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
)

const (
	playerSpriteFile = "internal/engine/assets/images/player.png"
	exitSpriteFile   = "internal/engine/assets/images/exit.png"
)

func createPlayer(world *entities.World, playerSize, cellSize int) entities.Entity {
	player := world.NewEntity()

	world.AddComponent(player, &components.Size{Width: float64(playerSize), Height: float64(playerSize)})
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})

	posX := float64(cellSize-playerSize) / 2
	posY := float64(cellSize-playerSize) / 2
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

func createMaze(world *entities.World, mazeWidth, mazeHeight int, cellSize int) entities.Entity {
	mazeEntity := world.NewEntity()
	world.AddComponent(mazeEntity, &components.Maze{
		Layout:   layout.New(mazeWidth, mazeHeight),
		CellSize: cellSize,
	})

	return mazeEntity
}

func createExit(world *entities.World, mazeCol, mazeRow, cellSize int) entities.Entity {
	exit := world.NewEntity()
	world.AddComponent(exit, &components.Size{Width: float64(cellSize), Height: float64(cellSize)})

	posX := float64(mazeCol * cellSize)
	posY := float64(mazeRow * cellSize)
	world.AddComponent(exit, &components.Position{X: posX, Y: posY})

	world.AddComponent(exit, &components.Exit{})

	exitSprite := utils.MustLoadSprite(exitSpriteFile)
	world.AddComponent(exit, &components.Sprite{Image: exitSprite})

	return exit
}
