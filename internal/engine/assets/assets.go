package assets

import (
	_ "embed"
)

//go:embed sounds/collectible-bip.wav
var CollectibleBip []byte

//go:embed sounds/level-completed.wav
var LevelCompleted []byte

//go:embed images/player.png
var PlayerImage []byte
