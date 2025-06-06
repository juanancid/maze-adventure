package definitions

// Level01 returns the configuration for level 1
func Level01() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:          8,
			Rows:          5,
			DeadlyCells:   2,
			FreezingCells: 3,
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 7,
				Y: 4,
			},
			Size: 24,
		},
		Collectibles: Collectibles{
			Number: 2,
			Size:   8,
			Value:  1,
		},
		Timer: 30, // 30 seconds for level 1
	}
}
