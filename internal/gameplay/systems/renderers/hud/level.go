package hud

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// LevelRenderer handles drawing the level number
type LevelRenderer struct {
	faceSource *text.GoTextFaceSource
}

func NewLevelRenderer(faceSource *text.GoTextFaceSource) *LevelRenderer {
	return &LevelRenderer{
		faceSource: faceSource,
	}
}

func (r *LevelRenderer) Draw(gameSession *session.GameSession, screen *ebiten.Image) {
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
}
