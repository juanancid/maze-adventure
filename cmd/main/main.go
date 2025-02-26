package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/config"
	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/systems"
	"github.com/juanancid/maze-adventure/internal/levels"
)

type Game struct {
	world *ecs.World
}

func (g *Game) Update() error {
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*config.ScaleFactor, config.ScreenHeight*config.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func newGame() *Game {
	world := ecs.NewWorld()

	levels.CreateLevel(world)
	addSystems(world)
	AddRenderers(world)

	game := &Game{
		world: world,
	}
	return game
}

func addSystems(world *ecs.World) {
	world.AddSystem(&systems.InputControl{})
	world.AddSystem(&systems.Movement{})
	world.AddSystem(&systems.MazeCollisionSystem{})
}

func AddRenderers(world *ecs.World) {
	world.AddRenderer(&systems.MazeRenderer{})
	world.AddRenderer(&systems.Renderer{})
}
