package definitions

// Level02 -> Introduce deadly cells
func Level02() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  10,
			Rows:                  6,
			LethalCells:           3,
			FreezingCells:         0,
			Patrollers:            2,
			ExtraConnectionChance: 0.04,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 9,
				Y: 5,
			},
			Size: DefaultExitSize,
		},
		Collectibles: Collectibles{
			Number: 3,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 45,
	}
}
