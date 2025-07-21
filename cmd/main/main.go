// Maze Adventure Game
//
// Command-line options:
//
//	--start-level N, -l N    Start the game at level N (1-4) for development/testing
//
// Examples:
//
//	go run ./cmd/main                    # Start at level 1 (normal gameplay)
//	go run ./cmd/main --start-level 3    # Start at level 3 (development mode)
//	go run ./cmd/main -l 2               # Start at level 2 (development mode)
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/app"
	engineconfig "github.com/juanancid/maze-adventure/internal/engine/config"
	gameplayconfig "github.com/juanancid/maze-adventure/internal/gameplay/config"
)

func main() {
	// Parse command-line arguments
	startLevel := flag.Int("start-level", 1, "Starting level (1-4)")
	startLevelShort := flag.Int("l", 1, "Starting level (1-4) - short form")
	flag.Parse()

	// Use the short form if provided, otherwise use the long form
	selectedLevel := *startLevel
	if *startLevelShort != 1 && *startLevel == 1 {
		selectedLevel = *startLevelShort
	}

	// Validate level number
	if selectedLevel < 1 || selectedLevel > 4 {
		fmt.Fprintf(os.Stderr, "Error: Invalid level number %d. Must be between 1 and 4.\n", selectedLevel)
		fmt.Fprintf(os.Stderr, "Usage: %s [--start-level N] or [--l N]\n", os.Args[0])
		os.Exit(1)
	}

	if selectedLevel > 1 {
		fmt.Printf("Starting game at level %d (development mode)\n", selectedLevel)
	}

	ebiten.SetWindowSize(engineconfig.ScreenWidth*engineconfig.ScaleFactor, engineconfig.ScreenHeight*engineconfig.ScaleFactor)
	ebiten.SetWindowTitle("Maze Adventure")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)

	gameConfig := gameplayconfig.GameConfig{
		StartingHearts: 3,
		StartingLevel:  selectedLevel,
	}

	g := app.NewGame(gameConfig)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
