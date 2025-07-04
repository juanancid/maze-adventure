package utils

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/juanancid/maze-adventure/internal/engine/assets"
)

type GameImage int

const (
	ImagePlayer GameImage = iota
	ImageExit
	ImageCollectible
	ImageIntroIllustration
)

var (
	gameImages   = map[GameImage]*ebiten.Image{}
	imageSources = map[GameImage][]byte{
		ImagePlayer:            assets.PlayerImage,
		ImageExit:              assets.ExitImage,
		ImageCollectible:       assets.CollectibleImage,
		ImageIntroIllustration: assets.IntroIllustration,
	}
)

// PreloadImages loads all game images into the cache
func PreloadImages() {
	for img := range imageSources {
		if err := loadImage(img); err != nil {
			log.Printf("error preloading image %d: %v", img, err)
		}
	}
}

func loadImage(image GameImage) error {
	if _, exists := gameImages[image]; exists {
		return nil
	}

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imageSources[image]))
	if err != nil {
		return err
	}

	gameImages[image] = img
	return nil
}

// GetImage returns the cached image or nil if not loaded
func GetImage(image GameImage) *ebiten.Image {
	img := gameImages[image]
	if img == nil {
		// Try to load the image if it's not cached
		if err := loadImage(image); err != nil {
			log.Printf("failed to load image %d: %v", image, err)
			return nil
		}
		img = gameImages[image]
	}
	return img
}
