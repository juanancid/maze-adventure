package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"maze-adventure/internal/config"
	"maze-adventure/internal/ecs"
	"maze-adventure/internal/ecs/components"
	"maze-adventure/internal/ecs/systems"
	"maze-adventure/internal/utils"
)

type Game struct {
	World *ecs.World
}

func NewGame() *Game {
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
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*config.ScaleFactor, config.ScreenHeight*config.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizable(true)

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
