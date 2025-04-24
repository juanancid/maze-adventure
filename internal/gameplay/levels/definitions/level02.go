package definitions

// Level02 returns the configuration for level 2
func Level02() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols: 16,
			Rows: 10,
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 15,
				Y: 9,
			},
			Size: 12,
		},
		Collectibles: Collectibles{
			Number: 4,
			Size:   8,
			Value:  1,
		},
	}
}
