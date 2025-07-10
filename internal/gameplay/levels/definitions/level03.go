package definitions

// Level03 -> Introduce freezing cells
func Level03() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  12,
			Rows:                  8,
			DeadlyCells:           2,
			FreezingCells:         4,
			ExtraConnectionChance: 0.05,
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
			Number: 4,
			Size:   8,
			Value:  1,
		},
		Timer: 60,
	}
}
