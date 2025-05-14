package utils

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/juanancid/maze-adventure/internal/engine/assets"
)

var (
	audioContext = audio.NewContext(44100)
	player       *audio.Player
)

func PlayCollectibleSound() {
	if player == nil {
		// Decode the WAV file
		stream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(assets.CollectibleBip))
		if err != nil {
			log.Printf("error decoding sound: %v", err)
			return
		}

		// Create a new player
		player, err = audioContext.NewPlayer(stream)
		if err != nil {
			log.Printf("error creating audio player: %v", err)
			return
		}
	}

	// Rewind the player to the beginning
	err := player.Rewind()
	if err != nil {
		log.Printf("error rewinding sound: %v", err)
	}
	// Play the sound
	player.Play()
}
