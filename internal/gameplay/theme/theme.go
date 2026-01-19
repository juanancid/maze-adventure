package theme

import (
	"image/color"

	"github.com/juanancid/maze-adventure/internal/engine/utils/palette"
)

// Theme defines "what colors mean" in Maze Adventure.
type Theme struct {
	// Global
	HudBackground   color.RGBA
	MazeBackground  color.RGBA
	IntroBackground color.RGBA
	UIText          color.RGBA
	UITimer         color.RGBA
	UIIntroText     color.RGBA

	// World
	WallNormal   color.RGBA
	WallFreezing color.RGBA
	WallLethal   color.RGBA

	// Actors
	PlayerBody color.RGBA
	PlayerCore color.RGBA
	EnemyBody  color.RGBA
	EnemyCore  color.RGBA

	// Items / goal
	CoinPrimary color.RGBA
	CoinAccent  color.RGBA
	GoalPrimary color.RGBA
	GoalAccent  color.RGBA
}

// DefaultTheme is the canonical mapping for Maze Adventure.
// Game code should depend on Theme, not on ENDESGA16 directly.
var DefaultTheme = Theme{
	// Global
	HudBackground:   palette.ENDESGA16.Ink1, // Slightly lighter than maze for visual separation
	MazeBackground:  palette.ENDESGA16.Ink0, // Darkest background for game area
	IntroBackground: palette.ENDESGA16.Ink0,
	UIText:          palette.ENDESGA16.Ink5,
	UITimer:         palette.ENDESGA16.Lime,
	UIIntroText:     palette.ENDESGA16.Green,

	// World
	WallNormal:   palette.ENDESGA16.Green,
	WallFreezing: palette.ENDESGA16.Cyan,
	WallLethal:   palette.ENDESGA16.Red,

	// Actors
	PlayerBody: palette.ENDESGA16.Green,
	PlayerCore: palette.ENDESGA16.Cyan,
	EnemyBody:  palette.ENDESGA16.Orange,
	EnemyCore:  palette.ENDESGA16.Red,

	// Items / goal
	CoinPrimary: palette.ENDESGA16.Lime,
	CoinAccent:  palette.ENDESGA16.Yellow,
	GoalPrimary: palette.ENDESGA16.Paper1,
	GoalAccent:  palette.ENDESGA16.Green,
}
