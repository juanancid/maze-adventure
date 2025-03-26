package app

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/events"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/renderers"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/updaters"
)

type Game struct {
	world        *entities.World
	levelManager *levels.Manager

	updaters  []Updater
	renderers []Renderer

	eventBus *events.Bus
}

type Updater interface {
	Update(w *entities.World)
}

type Renderer interface {
	Draw(world *entities.World, screen *ebiten.Image)
}

func NewGame() *Game {
	game := newEmptyGame()

	game.loadNextLevel()
	game.setUpdaters()
	game.setRenderers()
	game.setupEventSubscriptions()

	return game
}

func newEmptyGame() *Game {
	return &Game{
		levelManager: levels.NewManager(),
		eventBus:     events.NewBus(),
	}
}

func (g *Game) loadNextLevel() {
	levelConfig, err := g.levelManager.NextLevel()
	if err != nil {
		panic(err)
	}

	g.world = levels.CreateLevel(levelConfig)
}

func (g *Game) setUpdaters() {
	g.updaters = make([]Updater, 0)

	g.addUpdater(&updaters.InputControl{})
	g.addUpdater(&updaters.Movement{})
	g.addUpdater(&updaters.MazeCollisionSystem{})
	g.addUpdater(updaters.NewScoreSystem(g.eventBus))
}

func (g *Game) addUpdater(s Updater) {
	g.updaters = append(g.updaters, s)
}

func (g *Game) setRenderers() {
	g.renderers = make([]Renderer, 0)

	g.addRenderer(&renderers.MazeRenderer{})
	g.addRenderer(&renderers.SpriteRenderer{})
}

func (g *Game) addRenderer(r Renderer) {
	g.renderers = append(g.renderers, r)
}

func (g *Game) setupEventSubscriptions() {
	g.eventBus.Subscribe(reflect.TypeOf(events.LevelCompletedEvent{}), g.onLevelCompleted)
}

func (g *Game) onLevelCompleted(e events.Event) {
	_, ok := e.(events.LevelCompletedEvent)
	if !ok {
		return
	}

	g.loadNextLevel()
}

func (g *Game) Update() error {
	if err := g.update(); err != nil {
		return err
	}

	g.processEvents()

	return nil
}

func (g *Game) update() error {
	for _, s := range g.updaters {
		s.Update(g.world)
	}

	return nil
}

func (g *Game) processEvents() {
	g.eventBus.Process()
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.renderers {
		r.Draw(g.world, screen)
	}
}

func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
