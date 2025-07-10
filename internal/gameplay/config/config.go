package config

// GameConfig holds the game's configuration
type GameConfig struct {
	StartingHearts int // Number of hearts the player starts with
	StartingLevel  int // Level to start the game at (1-4, default: 1)
}
