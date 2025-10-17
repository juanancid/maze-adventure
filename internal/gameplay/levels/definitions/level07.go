package definitions

// Level07 -> Calm "Zen" finale (no hazards, exploration vibes)
func Level07() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  16,
			Rows:                  10,
			DeadlyCells:           0,
			FreezingCells:         0,
			Patrollers:            0,
			ExtraConnectionChance: 0.15, // generous loops for roaming
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{X: 15, Y: 9},
			Size:     16,
		},
		Collectibles: Collectibles{
			Number: 5,
			Size:   8,
			Value:  1,
		},
		Timer: 120,
	}
}
