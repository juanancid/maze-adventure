package hud

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/utils/palette"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
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
	// Create the font face once to reuse
	face := &text.GoTextFace{
		Source: r.faceSource,
		Size:   8,
	}

	// Measure total width of all hearts to position them at far right
	allHearts := strings.Repeat("♥", gameSession.MaxHearts)
	totalHeartsWidth, _ := text.Measure(allHearts, face, 0)

	// Calculate starting position to align hearts at far right with 8px margin
	baseX := float64(config.ScreenWidth) - totalHeartsWidth - 8
	baseY := float64(config.HudHeight/2 - 4)

	// Draw alive hearts in red
	var aliveHeartsWidth float64
	if gameSession.CurrentHearts > 0 {
		aliveHearts := strings.Repeat("♥", gameSession.CurrentHearts)
		aliveOp := &text.DrawOptions{}
		aliveOp.GeoM.Translate(baseX, baseY)
		aliveOp.ColorScale.ScaleWithColor(palette.ENDESGA16.Red)

		text.Draw(screen, aliveHearts, face, aliveOp)

		// Measure the actual width of alive hearts for positioning dead hearts
		aliveHeartsWidth, _ = text.Measure(aliveHearts, face, 0)
	}

	// Draw dead hearts in gray, positioned after alive hearts
	if gameSession.CurrentHearts < gameSession.MaxHearts {
		deadHearts := strings.Repeat("♥", gameSession.MaxHearts-gameSession.CurrentHearts)

		deadOp := &text.DrawOptions{}
		deadOp.GeoM.Translate(baseX+aliveHeartsWidth, baseY)
		deadOp.ColorScale.ScaleWithColor(palette.ENDESGA16.Ink3)

		text.Draw(screen, deadHearts, face, deadOp)
	}
}
