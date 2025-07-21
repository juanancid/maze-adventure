package utils

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/juanancid/maze-adventure/internal/engine/assets"
)

const sampleRate = 44100

var (
	audioContext = audio.NewContext(sampleRate)
	players      = map[SoundEffect]*audio.Player{}
	musicPlayer  *audio.Player
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

// StartBackgroundMusic starts playing the background music in a loop
func StartBackgroundMusic() {
	if musicPlayer != nil {
		// Music is already loaded, just start playing
		if !musicPlayer.IsPlaying() {
			err := musicPlayer.Rewind()
			if err != nil {
				log.Printf("error rewinding background music: %v", err)
				return
			}
			musicPlayer.Play()
		}
		return
	}

	// Load and start the background music
	if err := loadBackgroundMusic(); err != nil {
		log.Printf("Background music failed to load: %v", err)
		log.Printf("Game will continue without background music")
		return
	}

	if musicPlayer != nil {
		musicPlayer.Play()
	}
}

// StopBackgroundMusic stops the background music
func StopBackgroundMusic() {
	if musicPlayer != nil && musicPlayer.IsPlaying() {
		musicPlayer.Pause()
	}
}

// loadBackgroundMusic loads the background music file
func loadBackgroundMusic() error {
	if musicPlayer != nil {
		return nil // Already loaded
	}

	// Check if we have any data
	if len(assets.BackgroundMusic) == 0 {
		return fmt.Errorf("background music file is empty")
	}

	// Decode the OGG Vorbis file
	stream, err := vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(assets.BackgroundMusic))
	if err != nil {
		return fmt.Errorf("vorbis decode error: %w", err)
	}

	// Create an infinite loop stream
	loopStream := audio.NewInfiniteLoop(stream, stream.Length())

	// Create a new player
	player, err := audioContext.NewPlayer(loopStream)
	if err != nil {
		return fmt.Errorf("player creation error: %w", err)
	}

	musicPlayer = player
	return nil
}
