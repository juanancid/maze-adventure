package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/app"
	"github.com/juanancid/maze-adventure/internal/engine/config"
)

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*config.ScaleFactor, config.ScreenHeight*config.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	g := app.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
