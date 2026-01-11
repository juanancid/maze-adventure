package definitions

const (
	DefaultPlayerSize       = 12
	DefaultExitSize         = 16
	LargeExitSize           = 24
	DefaultCollectibleSize  = 12
	DefaultCollectibleValue = 1
)

// LevelRegistry holds all level definitions
var LevelRegistry = []func() LevelConfig{
	Level01,
	Level02,
	Level03,
	Level04,
	Level05,
	Level06,
	Level07,
}
