package definitions

// Level10 -> Final level
func Level10() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  10,
			Rows:                  10,
			LethalCells:           0,
			FreezingCells:         0,
			Patrollers:            0,
			ExtraConnectionChance: 0.25,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{X: 5, Y: 5},
			Size:     DefaultExitSize,
		},
		Collectibles: Collectibles{
			Number: 15,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 20,
	}
}
