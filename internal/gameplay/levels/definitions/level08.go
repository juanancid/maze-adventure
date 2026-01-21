package definitions

// Level08 ->
func Level08() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  11,
			Rows:                  7,
			LethalCells:           8,
			FreezingCells:         10,
			Patrollers:            7,
			ExtraConnectionChance: 0.05,
		},
		Player: PlayerConfig{
			Size: DefaultPlayerSize,
		},
		Exit: ExitConfig{
			Position: Coordinate{X: 9, Y: 1},
			Size:     LittleExitSize,
		},
		Collectibles: Collectibles{
			Number: 6,
			Size:   DefaultCollectibleSize,
			Value:  DefaultCollectibleValue,
		},
		Timer: 60,
	}
}
