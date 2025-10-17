package definitions

// Level05 -> Large maze, emphasis on exploration (light hazards)
func Level05() LevelConfig {
	return LevelConfig{
		Maze: MazeConfig{
			Cols:                  18,
			Rows:                  12,
			DeadlyCells:           3,
			FreezingCells:         2,
			Patrollers:            2,
			ExtraConnectionChance: 0.10, // a few more loops for a freer feel
		},
		Player: PlayerConfig{
			Size: 12,
		},
		Exit: ExitConfig{
			Position: Coordinate{X: 17, Y: 11}, // bottom-right
			Size:     16,
		},
		Collectibles: Collectibles{
			Number: 6,
			Size:   8,
			Value:  1,
		},
		Timer: 90,
	}
}
