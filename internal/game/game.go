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
	game := newEmptyGame()

	world := levels.CreateLevelWorld()
	game.setWorld(world)

	game.addSystems()
	game.addRenderers()

	return game
}

func newEmptyGame() *Game {
	return &Game{}
}

func (g *Game) setWorld(world *ecs.World) {
	g.world = world
}

func (g *Game) Update() error {
	if err := g.update(); err != nil {
		return err
	}

	if g.scoreSystem.LevelCompleted {
		g.scoreSystem.LevelCompleted = false
		g.setWorld(levels.CreateLevelWorld())
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
	g.updaters = make([]Updater, 0)

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
	g.renderers = make([]Renderer, 0)

	g.addRenderer(&systems.MazeRenderer{})
	g.addRenderer(&systems.Renderer{})
}

func (g *Game) addRenderer(r Renderer) {
	g.renderers = append(g.renderers, r)
}
