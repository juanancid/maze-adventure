package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juanancid/maze-adventure/internal/ecs/events"

	"github.com/juanancid/maze-adventure/internal/config"
	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/systems"
	"github.com/juanancid/maze-adventure/internal/levels"
)

type Game struct {
	world *ecs.World

	updaters  []Updater
	renderers []Renderer
}

type Updater interface {
	Update(w *ecs.World)
}

type Renderer interface {
	Draw(world *ecs.World, screen *ebiten.Image)
}

func NewGame() *Game {
	game := newEmptyGame()

	game.setWorld(levels.CreateLevelWorld())
	game.setUpdaters()
	game.setRenderers()

	return game
}

func newEmptyGame() *Game {
	return &Game{}
}

func (g *Game) setWorld(world *ecs.World) {
	g.world = world
}

func (g *Game) setUpdaters() {
	g.updaters = make([]Updater, 0)

	g.addUpdater(&systems.InputControl{})
	g.addUpdater(&systems.Movement{})
	g.addUpdater(&systems.MazeCollisionSystem{})
	g.addUpdater(&systems.ScoreSystem{})
}

func (g *Game) addUpdater(s Updater) {
	g.updaters = append(g.updaters, s)
}

func (g *Game) setRenderers() {
	g.renderers = make([]Renderer, 0)

	g.addRenderer(&systems.MazeRenderer{})
	g.addRenderer(&systems.SpriteRenderer{})
}

func (g *Game) addRenderer(r Renderer) {
	g.renderers = append(g.renderers, r)
}

func (g *Game) Update() error {
	if err := g.update(); err != nil {
		return err
	}

	for _, event := range g.world.DrainEvents() {
		switch event.(type) {
		case events.LevelCompletedEvent:
			g.setWorld(levels.CreateLevelWorld())
		}
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
