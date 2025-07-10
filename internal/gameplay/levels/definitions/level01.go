package definitions

// Level01 -> Movement and collecting (no hazards)
func Level01() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  8,
			Rows:                  5,
			DeadlyCells:           0,
			FreezingCells:         0,
			Patrollers:            0,
			ExtraConnectionChance: 0.0,
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
		Timer: 30,
	}
}
