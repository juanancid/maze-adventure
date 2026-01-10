package definitions

// Level04 -> Challenge with all mechanics
func Level04() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  14,
			Rows:                  9,
			DeadlyCells:           4,
			FreezingCells:         6,
			Patrollers:            4,
			ExtraConnectionChance: 0.12,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 13,
				Y: 8,
			},
			Size: DefaultExitSize,
		},
		Collectibles: Collectibles{
			Number: 5,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 75,
	}
}
