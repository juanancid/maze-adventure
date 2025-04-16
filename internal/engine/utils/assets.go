package utils

import (
	"bytes"
	_ "image/png" // To support PNG decoding

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// MustLoadSprite loads an image file and returns an *ebiten.Image.
func MustLoadSprite(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}

	return img
}

// MustLoadGoTextFaceSource parses an OpenType or TrueType font and returns a GoTextFaceSource object.
func MustLoadGoTextFaceSource(font []byte) *text.GoTextFaceSource {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		panic(err)
	}

	return fontSource
}
