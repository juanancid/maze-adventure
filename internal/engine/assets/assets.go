package assets

import (
	_ "embed"
)

//go:embed sounds/collectible-bip.wav
var CollectibleBip []byte

//go:embed sounds/level-completed.wav
var LevelCompleted []byte

//go:embed sounds/damage.wav
var DamageSound []byte

//go:embed sounds/freeze.wav
var FreezeSound []byte

//go:embed images/player.png
var PlayerImage []byte

//go:embed images/exit.png
var ExitImage []byte

//go:embed images/collectible.png
var CollectibleImage []byte

//go:embed images/intro-illustration.png
var IntroIllustration []byte
