package definitions

// Level06 -> Small map, high tension (many patrollers, tight timer)
func Level06() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  10,
			Rows:                  6,
			DeadlyCells:           3,
			FreezingCells:         2,
			Patrollers:            5,
			ExtraConnectionChance: 0.02, // more corridors, fewer shortcuts
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{X: 9, Y: 5},
			Size:     16,
		},
		Collectibles: Collectibles{
			Number: 4,
			Size:   8,
			Value:  1,
		},
		Timer: 45,
	}
}
