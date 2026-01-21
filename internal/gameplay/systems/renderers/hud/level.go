package hud

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/gameplay/session"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"
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
	levelText := fmt.Sprintf("LEVEL %d", gameSession.CurrentLevel)
	levelOp := &text.DrawOptions{}
	// Position at bottom-left of HUD, below score
	levelOp.GeoM.Translate(8, 26)
	levelOp.ColorScale.ScaleWithColor(theme.DefaultTheme.UIText)

	text.Draw(screen,
		levelText,
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		levelOp,
	)
}
