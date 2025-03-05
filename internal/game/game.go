package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/config"
	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/systems"
	"github.com/juanancid/maze-adventure/internal/levels"
)

type Game struct {
	world *ecs.World

	scoreSystem *systems.ScoreSystem
}

func NewGame() *Game {
	world := ecs.NewWorld()

	game := &Game{
		world: world,
	}

	levels.CreateLevel(world)
	game.addSystems(world)
	AddRenderers(world)

	return game
}

func (g *Game) Update() error {
	g.world.Update()

	if g.scoreSystem.LevelCompleted {
		g.scoreSystem.LevelCompleted = false

		g.world = ecs.NewWorld()
		levels.CreateLevel(g.world)
		g.addSystems(g.world)
		AddRenderers(g.world)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) addSystems(world *ecs.World) {
	world.AddSystem(&systems.InputControl{})
	world.AddSystem(&systems.Movement{})
	world.AddSystem(&systems.MazeCollisionSystem{})

	scoreSystem := &systems.ScoreSystem{LevelCompleted: false}
	g.scoreSystem = scoreSystem
	world.AddSystem(scoreSystem)
}

func AddRenderers(world *ecs.World) {
	world.AddRenderer(&systems.MazeRenderer{})
	world.AddRenderer(&systems.Renderer{})
}
