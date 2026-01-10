package theme

import (
	"image/color"

	"github.com/juanancid/maze-adventure/internal/engine/utils/palette"
)

// Theme defines "what colors mean" in Maze Adventure.
type Theme struct {
	// Global
	Background  color.RGBA
	UIText      color.RGBA
	UISecondary color.RGBA

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
	Background:  palette.ENDESGA16.Ink0,
	UIText:      palette.ENDESGA16.Lime,
	UISecondary: palette.ENDESGA16.Paper1,

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
