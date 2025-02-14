package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juanancid/maze-adventure/internal/ecs/systems"

	"github.com/juanancid/maze-adventure/internal/config"
	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/levels"
)

type Game struct {
	World *ecs.World
}

func (g *Game) Update() error {
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(screen)
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*config.ScaleFactor, config.ScreenHeight*config.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizable(true)

	game := newGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func newGame() *Game {
	world := ecs.NewWorld()

	levels.CreateLevel(world)

	world.AddSystem(&systems.InputControl{})
	world.AddSystem(&systems.Movement{})

	world.AddRenderable(&systems.MazeRenderer{})
	world.AddRenderable(&systems.Renderer{})

	game := &Game{
		World: world,
	}
	return game
}
