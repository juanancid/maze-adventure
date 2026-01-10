package definitions

// Level03 -> Introduce freezing cells
func Level03() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  12,
			Rows:                  8,
			DeadlyCells:           2,
			FreezingCells:         4,
			Patrollers:            4,
			ExtraConnectionChance: 0.07,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 11,
				Y: 7,
			},
			Size: DefaultExitSize,
		},
		Collectibles: Collectibles{
			Number: 4,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 60,
	}
}
