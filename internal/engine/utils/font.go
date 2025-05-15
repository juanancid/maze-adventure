package utils

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// MustLoadGoTextFaceSource parses an OpenType or TrueType font and returns a GoTextFaceSource object.
func MustLoadGoTextFaceSource(font []byte) *text.GoTextFaceSource {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		panic(err)
	}

	return fontSource
}
