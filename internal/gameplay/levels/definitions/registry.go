package definitions

// LevelRegistry holds all level definitions
var LevelRegistry = []func() LevelConfig{
	Level01,
	Level02,
	Level03,
	Level04,
}
