package utils

import (
	_ "image/png" // To support PNG decoding

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// MustLoadSprite loads an image file and returns an *ebiten.Image.
func MustLoadSprite(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}

	return img
}
