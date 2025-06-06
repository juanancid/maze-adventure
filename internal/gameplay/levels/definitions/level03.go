package definitions

// Level03 returns the configuration for level 3 (no timer)
func Level03() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:          12,
			Rows:          8,
			DeadlyCells:   3,
			FreezingCells: 4,
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 11,
				Y: 7,
			},
			Size: 16,
		},
		Collectibles: Collectibles{
			Number: 3,
			Size:   8,
			Value:  1,
		},
		Timer: 0, // No timer for level 3
	}
}
