package definitions

// Level02 -> Introduce deadly cells
func Level02() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  10,
			Rows:                  6,
			DeadlyCells:           3,
			FreezingCells:         0,
			Patrollers:            1,
			ExtraConnectionChance: 0.04,
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 9,
				Y: 5,
			},
			Size: 16,
		},
		Collectibles: Collectibles{
			Number: 3,
			Size:   8,
			Value:  1,
		},
		Timer: 45,
	}
}
