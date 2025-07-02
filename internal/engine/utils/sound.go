package utils

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/juanancid/maze-adventure/internal/engine/assets"
)

const sampleRate = 44100

var (
	audioContext = audio.NewContext(sampleRate)
	players      = map[SoundEffect]*audio.Player{}
)

type SoundEffect int

const (
	SoundCollectibleBip SoundEffect = iota
	SoundLevelCompleted
	SoundDamage
	SoundFreeze
)

var soundSources = map[SoundEffect][]byte{
	SoundCollectibleBip: assets.CollectibleBip,
	SoundLevelCompleted: assets.LevelCompleted,
	SoundDamage:         assets.DamageSound,
	SoundFreeze:         assets.FreezeSound,
}

// PreloadSounds loads all game sounds into the cache
func PreloadSounds() {
	for sound := range soundSources {
		if err := loadSound(sound); err != nil {
			log.Printf("error preloading sound %d: %v", sound, err)
		}
	}
}

func loadSound(sound SoundEffect) error {
	if _, exists := players[sound]; exists {
		return nil
	}

	// Decode the WAV file
	stream, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(soundSources[sound]))
	if err != nil {
		return err
	}

	// Create a new player
	player, err := audioContext.NewPlayer(stream)
	if err != nil {
		return err
	}

	players[sound] = player
	return nil
}

func PlaySound(sound SoundEffect) {
	player, exists := players[sound]
	if !exists {
		log.Printf("sound %d not loaded", sound)
		return
	}

	// Rewind the player to the beginning
	err := player.Rewind()
	if err != nil {
		log.Printf("error rewinding sound: %v", err)
		return
	}
	// Play the sound
	player.Play()
}
