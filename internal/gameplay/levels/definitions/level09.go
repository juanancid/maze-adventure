package definitions

// Level09 ->
func Level09() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  14,
			Rows:                  8,
			LethalCells:           40,
			FreezingCells:         0,
			Patrollers:            0,
			ExtraConnectionChance: 0,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{X: 13, Y: 6},
			Size:     LittleExitSize,
		},
		Collectibles: Collectibles{
			Number: 0,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 40,
	}
}
