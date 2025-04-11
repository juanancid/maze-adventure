package levels

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/mazebuilder"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
)

const (
	playerSpriteFile = "internal/engine/assets/images/player.png"
	exitSpriteFile   = "internal/engine/assets/images/exit.png"
)

func CreateLevel(level *Level) *entities.World {
	world := entities.NewWorld()

	mazeCols := level.Maze.Cols
	mazeRows := level.Maze.Rows

	cellWidth := config.ScreenWidth / mazeCols
	cellHeight := (config.ScreenHeight - config.HudHeight) / mazeRows

	playerSize := level.Player.Size

	createPlayer(world, playerSize, cellWidth, cellHeight)
	createMaze(world, mazeCols, mazeRows, cellWidth, cellHeight)
	createExit(world, level.Exit.Position.X, level.Exit.Position.Y, cellWidth, cellHeight, level.Exit.Size)

	// Add level information to the world
	levelEntity := world.NewEntity()
	world.AddComponent(levelEntity, &components.Level{
		Number: level.Number,
	})

	return world
}

func createPlayer(world *entities.World, playerSize, cellWidth, cellHeight int) entities.Entity {
	player := world.NewEntity()

	world.AddComponent(player, &components.Size{Width: float64(playerSize), Height: float64(playerSize)})
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})

	posX := float64(cellWidth-playerSize) / 2
	posY := float64(cellHeight-playerSize) / 2
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

func createMaze(world *entities.World, mazeCols, mazeRows int, cellWidth, cellHeight int) entities.Entity {
	mazeEntity := world.NewEntity()
	world.AddComponent(mazeEntity, &components.Maze{
		Layout:     mazebuilder.NewMazeLayout(mazeCols, mazeRows),
		CellWidth:  cellWidth,
		CellHeight: cellHeight,
	})

	return mazeEntity
}

func createExit(world *entities.World, mazeCol, mazeRow, cellWidth, cellHeight, exitSize int) entities.Entity {
	exit := world.NewEntity()
	world.AddComponent(exit, &components.Size{Width: float64(exitSize), Height: float64(exitSize)})

	// Calculate the center position of the cell
	cellX := float64(mazeCol * cellWidth)
	cellY := float64(mazeRow * cellHeight)

	// Center the exit in the cell
	posX := cellX + float64(cellWidth-exitSize)/2
	posY := cellY + float64(cellHeight-exitSize)/2

	world.AddComponent(exit, &components.Position{X: posX, Y: posY})

	world.AddComponent(exit, &components.Exit{})

	exitSprite := utils.MustLoadSprite(exitSpriteFile)
	world.AddComponent(exit, &components.Sprite{Image: exitSprite})

	return exit
}
