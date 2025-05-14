package utils

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/juanancid/maze-adventure/internal/engine/assets"
)

var (
	audioContext         = audio.NewContext(44100)
	collectiblePlayer    *audio.Player
	levelCompletedPlayer *audio.Player
)

func PlayCollectibleSound() {
	if collectiblePlayer == nil {
		// Decode the WAV file
		stream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(assets.CollectibleBip))
		if err != nil {
			log.Printf("error decoding sound: %v", err)
			return
		}

		// Create a new player
		collectiblePlayer, err = audioContext.NewPlayer(stream)
		if err != nil {
			log.Printf("error creating audio player: %v", err)
			return
		}
	}

	// Rewind the player to the beginning
	err := collectiblePlayer.Rewind()
	if err != nil {
		log.Printf("error rewinding sound: %v", err)
	}
	// Play the sound
	collectiblePlayer.Play()
}

func PlayLevelCompletedSound() {
	if levelCompletedPlayer == nil {
		// Decode the WAV file
		stream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(assets.LevelCompleted))
		if err != nil {
			log.Printf("error decoding sound: %v", err)
			return
		}

		// Create a new player
		levelCompletedPlayer, err = audioContext.NewPlayer(stream)
		if err != nil {
			log.Printf("error creating audio player: %v", err)
			return
		}
	}

	// Rewind the player to the beginning
	err := levelCompletedPlayer.Rewind()
	if err != nil {
		log.Printf("error rewinding sound: %v", err)
	}
	// Play the sound
	levelCompletedPlayer.Play()
}
