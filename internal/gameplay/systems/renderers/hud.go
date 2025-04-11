package renderers

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
)

type HUDRenderer struct {
	faceSource *text.GoTextFaceSource
}

func NewHUDRenderer() HUDRenderer {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return HUDRenderer{faceSource: s}
}

func (r HUDRenderer) Draw(w *entities.World, screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(8, float64(config.HudHeight/2-4))
	textOp.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen,
		"MAZE ADVENTURE",
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		textOp,
	)
}
