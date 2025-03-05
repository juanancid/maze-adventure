package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juanancid/maze-adventure/internal/game"

	"github.com/juanancid/maze-adventure/internal/config"
)

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*config.ScaleFactor, config.ScreenHeight*config.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	aGame := game.NewGame()
	if err := ebiten.RunGame(aGame); err != nil {
		panic(err)
	}
}
