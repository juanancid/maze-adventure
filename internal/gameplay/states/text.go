package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"

	"image/color"
)

const (
	titleFontSize   = 16
	regularFontSize = 8
)

var (
	bgColor = color.RGBA{R: 0x00, G: 0x13, B: 0x1F, A: 0xFF}
	font    *text.GoTextFaceSource
	centerX = float64(config.ScreenWidth / 2)
)

func drawCenteredText(screen *ebiten.Image, txt string, y float64, size float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(centerX, y)
	op.ColorScale.ScaleWithColor(theme.DefaultTheme.UIIntroText)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter

	face := &text.GoTextFace{
		Source: font,
		Size:   size,
	}

	text.Draw(screen, txt, face, op)
}

func init() {
	font = utils.MustLoadGoTextFaceSource(fonts.PressStart2P_ttf)
}
