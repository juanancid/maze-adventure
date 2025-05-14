package utils

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/juanancid/maze-adventure/internal/engine/assets"
)

const sampleRate = 44100

type SoundEffect int

const (
	SoundCollectibleBip SoundEffect = iota
	SoundLevelCompleted
)

var (
	audioContext = audio.NewContext(sampleRate)

	soundPlayers = map[SoundEffect]*audio.Player{}
	soundSources = map[SoundEffect][]byte{
		SoundCollectibleBip: assets.CollectibleBip,
		SoundLevelCompleted: assets.LevelCompleted,
	}
)

func PlaySound(effect SoundEffect) {
	player, exists := soundPlayers[effect]
	if !exists {
		data, ok := soundSources[effect]
		if !ok {
			log.Printf("sound '%d' not found", effect)
			return
		}

		stream, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
		if err != nil {
			log.Printf("error decoding sound '%d': %v", effect, err)
			return
		}

		player, err = audioContext.NewPlayer(stream)
		if err != nil {
			log.Printf("error creating player for sound '%d': %v", effect, err)
			return
		}

		soundPlayers[effect] = player
	}

	_ = player.Rewind()
	player.Play()
}
