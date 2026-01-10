package hud

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"
)

// HealthRenderer handles drawing the health hearts
type HealthRenderer struct {
	faceSource *text.GoTextFaceSource
}

func NewHealthRenderer(faceSource *text.GoTextFaceSource) *HealthRenderer {
	return &HealthRenderer{
		faceSource: faceSource,
	}
}

func (r *HealthRenderer) Draw(gameSession *session.GameSession, screen *ebiten.Image) {
	hearts := strings.Repeat("♥", gameSession.CurrentHearts) + strings.Repeat("·", gameSession.MaxHearts-gameSession.CurrentHearts)
	heartsOp := &text.DrawOptions{}
	heartsOp.GeoM.Translate(float64(config.ScreenWidth/2-54), float64(config.HudHeight/2-4))
	heartsOp.ColorScale.ScaleWithColor(theme.DefaultTheme.UIText)

	text.Draw(screen,
		hearts,
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		heartsOp,
	)
}
