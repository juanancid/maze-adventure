package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/app"
	engineconfig "github.com/juanancid/maze-adventure/internal/engine/config"
	gameplayconfig "github.com/juanancid/maze-adventure/internal/gameplay/config"
)

func main() {
	ebiten.SetWindowSize(engineconfig.ScreenWidth*engineconfig.ScaleFactor, engineconfig.ScreenHeight*engineconfig.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	gameConfig := gameplayconfig.GameConfig{
		StartingHearts: 3,
	}

	g := app.NewGame(gameConfig)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
