package definitions

// Level03 returns the configuration for level 3
func Level03() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols: 32,
			Rows: 20,
		},
		Player: PlayerConfig{
			Size: 6,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 31,
				Y: 19,
			},
			Size: 6,
		},
	}
}
