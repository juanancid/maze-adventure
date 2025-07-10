package definitions

// Level04 -> Challenge with all mechanics
func Level04() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  14,
			Rows:                  9,
			DeadlyCells:           4,
			FreezingCells:         6,
			ExtraConnectionChance: 0.08,
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 13,
				Y: 8,
			},
			Size: 16,
		},
		Collectibles: Collectibles{
			Number: 5,
			Size:   8,
			Value:  1,
		},
		Timer: 75,
	}
}
