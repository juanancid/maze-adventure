package levels

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/mazebuilder"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels/definitions"
)

const (
	playerSpriteFile = "internal/engine/assets/images/player.png"
	exitSpriteFile   = "internal/engine/assets/images/exit.png"
)

func CreateLevel(levelConfig definitions.LevelConfig) *entities.World {
	world := entities.NewWorld()

	mazeCols := levelConfig.Maze.Cols
	mazeRows := levelConfig.Maze.Rows

	cellWidth := config.ScreenWidth / mazeCols
	cellHeight := (config.ScreenHeight - config.HudHeight) / mazeRows

	playerSize := levelConfig.Player.Size

	createPlayer(world, playerSize, cellWidth, cellHeight)
	createMaze(world, mazeCols, mazeRows, cellWidth, cellHeight)
	createExit(world, levelConfig.Exit.Position.X, levelConfig.Exit.Position.Y, cellWidth, cellHeight, levelConfig.Exit.Size)
	createCollectibles(world, levelConfig)

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

func createCollectibles(world *entities.World, levelConfig definitions.LevelConfig) {
	mazeCols := levelConfig.Maze.Cols
	mazeRows := levelConfig.Maze.Rows

	cellWidth := config.ScreenWidth / mazeCols
	cellHeight := (config.ScreenHeight - config.HudHeight) / mazeRows

	for i := 0; i < levelConfig.Collectibles.Number; i++ {
		// Generate random cell coordinates within maze bounds
		row := rand.Intn(mazeRows)
		col := rand.Intn(mazeCols)

		// Create a collectible at the random cell
		createCollectible(world, row, col, cellWidth, cellHeight, levelConfig.Collectibles.Value, levelConfig.Collectibles.Size)
	}
}

func createCollectible(world *entities.World, row, col, cellWidth, cellHeight, value, size int) {
	collectible := world.NewEntity()

	// Calculate the center position of the cell
	cellX := float64(col * cellWidth)
	cellY := float64(row * cellHeight)

	// Center the collectible in the cell
	x := cellX + float64(cellWidth-size)/2
	y := cellY + float64(cellHeight-size)/2

	world.AddComponent(collectible, &components.Position{X: x, Y: y})
	world.AddComponent(collectible, &components.Size{Width: float64(size), Height: float64(size)})
	world.AddComponent(collectible, &components.Collectible{
		Kind:  components.CollectibleScore,
		Value: value,
	})
	world.AddComponent(collectible, &components.Sprite{
		Image: utils.MustLoadSprite("internal/engine/assets/images/collectible-score.png"),
	})
}
