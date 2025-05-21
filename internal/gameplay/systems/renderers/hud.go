package renderers

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

type HUD struct {
	faceSource *text.GoTextFaceSource
}

func NewHUD() *HUD {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &HUD{
		faceSource: s,
	}
}

func (r *HUD) Draw(world *entities.World, gameSession *session.GameSession, screen *ebiten.Image) {
	// Draw score
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(8, float64(config.HudHeight/2-4)) // Vertically centered-ish
	textOp.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen,
		fmt.Sprintf("SCORE: %d", gameSession.Score),
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		textOp,
	)

	// Draw level number
	levelText := fmt.Sprintf("SECTOR %d", gameSession.CurrentLevel)
	levelOp := &text.DrawOptions{}
	levelOp.GeoM.Translate(float64(config.ScreenWidth-100), float64(config.HudHeight/2-4))
	levelOp.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen,
		levelText,
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		levelOp,
	)

	// Draw hearts
	hearts := strings.Repeat("♥", gameSession.CurrentHearts) + strings.Repeat("♡", gameSession.MaxHearts-gameSession.CurrentHearts)
	heartsOp := &text.DrawOptions{}
	heartsOp.GeoM.Translate(float64(config.ScreenWidth/2-24), float64(config.HudHeight/2-4))
	heartsOp.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen,
		hearts,
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		heartsOp,
	)
}
