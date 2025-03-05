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

	updaters    []Updater
	renderers   []Renderer
	scoreSystem *systems.ScoreSystem
}

type Updater interface {
	Update(w *ecs.World)
}

type Renderer interface {
	Draw(world *ecs.World, screen *ebiten.Image)
}

func NewGame() *Game {
	world := ecs.NewWorld()

	game := &Game{
		world:     world,
		updaters:  make([]Updater, 0),
		renderers: make([]Renderer, 0),
	}

	levels.CreateLevel(world)
	game.addSystems()
	game.addRenderers()

	return game
}

func (g *Game) Update() error {
	if err := g.update(); err != nil {
		return err
	}

	if g.scoreSystem.LevelCompleted {
		g.scoreSystem.LevelCompleted = false

		g.world = ecs.NewWorld()
		levels.CreateLevel(g.world)
	}

	return nil
}

func (g *Game) update() error {
	for _, s := range g.updaters {
		s.Update(g.world)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.renderers {
		r.Draw(g.world, screen)
	}
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) addSystems() {
	g.addSystem(&systems.InputControl{})
	g.addSystem(&systems.Movement{})
	g.addSystem(&systems.MazeCollisionSystem{})

	scoreSystem := &systems.ScoreSystem{LevelCompleted: false}
	g.scoreSystem = scoreSystem
	g.addSystem(scoreSystem)
}
func (g *Game) addSystem(s Updater) {
	g.updaters = append(g.updaters, s)
}

func (g *Game) addRenderers() {
	g.addRenderer(&systems.MazeRenderer{})
	g.addRenderer(&systems.Renderer{})
}

func (g *Game) addRenderer(r Renderer) {
	g.renderers = append(g.renderers, r)
}
