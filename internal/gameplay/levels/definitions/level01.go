package definitions

// Level01 -> Movement and collecting (no hazards)
func Level01() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  8,
			Rows:                  5,
			LethalCells:           0,
			FreezingCells:         0,
			Patrollers:            0,
			ExtraConnectionChance: 0.0,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{
				X: 7,
				Y: 4,
			},
			Size: DefaultExitSize,
		},
		Collectibles: Collectibles{
			Number: 2,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 30,
	}
}
