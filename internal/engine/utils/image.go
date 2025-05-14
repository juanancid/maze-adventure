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
)

var (
	gameImages   = map[GameImage]*ebiten.Image{}
	imageSources = map[GameImage][]byte{
		ImagePlayer: assets.PlayerImage,
	}
)

func LoadImage(image GameImage) *ebiten.Image {
	if cached, exists := gameImages[image]; exists {
		return cached
	}

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imageSources[image]))
	if err != nil {
		log.Printf("error loading image: %v", err)
		return nil
	}

	gameImages[image] = img
	return img
}
