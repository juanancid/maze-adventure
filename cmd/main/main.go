package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"maze-adventure/internal/config"
	"maze-adventure/internal/ecs"
	"maze-adventure/internal/ecs/components"
	"maze-adventure/internal/ecs/systems"
	"maze-adventure/internal/maze"
	"maze-adventure/internal/utils"
)

const (
	mazeWidth  = 10 // number of columns
	mazeHeight = 10 // number of rows
	cellSize   = 16 // size of each cell in pixels
)

type Game struct {
	World *ecs.World
	Maze  maze.Maze
}

func NewGame(maze maze.Maze) *Game {
	world := ecs.NewWorld()

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
	playerSprite, err := utils.LoadSprite("internal/assets/images/player.png")
	if err != nil {
		panic(err)
	}
	world.AddComponent(player, &components.Sprite{Image: playerSprite})

	game := &Game{
		World: world,
		Maze:  maze,
	}

	game.World.AddSystem(&systems.InputControl{})
	game.World.AddSystem(&systems.Movement{})

	game.World.AddRenderable(&systems.Renderer{})

	return game
}

func (g *Game) Update() error {
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(screen)

	// Set the wall color.
	wallColor := color.White

	// Iterate over each cell and draw its walls.
	for x := 0; x < g.Maze.Width; x++ {
		for y := 0; y < g.Maze.Height; y++ {
			cell := g.Maze.Grid[x][y]
			// Calculate pixel coordinates.
			x1 := float64(x*cellSize) + 1
			y1 := float64(y*cellSize) + 1
			x2 := float64((x+1)*cellSize) + 1
			y2 := float64((y+1)*cellSize) + 1

			// Draw top wall.
			if cell.Walls[0] {
				ebitenutil.DrawLine(screen, x1, y1, x2, y1, wallColor)
			}
			// Draw right wall.
			if cell.Walls[1] {
				ebitenutil.DrawLine(screen, x2, y1, x2, y2, wallColor)
			}
			// Draw bottom wall.
			if cell.Walls[2] {
				ebitenutil.DrawLine(screen, x2, y2, x1, y2, wallColor)
			}
			// Draw left wall.
			if cell.Walls[3] {
				ebitenutil.DrawLine(screen, x1, y2, x1, y1, wallColor)
			}
		}
	}
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*config.ScaleFactor, config.ScreenHeight*config.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizable(true)

	maze := maze.GenerateMaze(mazeWidth, mazeHeight)
	game := NewGame(maze)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
