package hud

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// ScoreRenderer handles drawing the score
type ScoreRenderer struct {
	faceSource *text.GoTextFaceSource
}

func NewScoreRenderer(faceSource *text.GoTextFaceSource) *ScoreRenderer {
	return &ScoreRenderer{
		faceSource: faceSource,
	}
}

func (r *ScoreRenderer) Draw(gameSession *session.GameSession, screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(8, float64(config.HudHeight/2-4))
	textOp.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen,
		fmt.Sprintf("SCORE: %d", gameSession.Score),
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		textOp,
	)
}
