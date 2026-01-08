package palette

import "image/color"

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
	Background:  ENDESGA16.Ink0,
	UIText:      ENDESGA16.Lime,
	UISecondary: ENDESGA16.Paper1,

	// World
	WallNormal:   ENDESGA16.Green,
	WallFreezing: ENDESGA16.Cyan,
	WallLethal:   ENDESGA16.Red,

	// Actors
	PlayerBody: ENDESGA16.Green,
	PlayerCore: ENDESGA16.Cyan,
	EnemyBody:  ENDESGA16.Orange,
	EnemyCore:  ENDESGA16.Red,

	// Items / goal
	CoinPrimary: ENDESGA16.Lime,
	CoinAccent:  ENDESGA16.Yellow,
	GoalPrimary: ENDESGA16.Paper1,
	GoalAccent:  ENDESGA16.Green,
}
