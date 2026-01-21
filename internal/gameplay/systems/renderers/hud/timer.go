package hud

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/utils/palette"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"
)

// TimerRenderer handles drawing the level timer
type TimerRenderer struct {
	faceSource *text.GoTextFaceSource
}

func NewTimerRenderer(faceSource *text.GoTextFaceSource) *TimerRenderer {
	return &TimerRenderer{
		faceSource: faceSource,
	}
}

func (r *TimerRenderer) Draw(gameSession *session.GameSession, screen *ebiten.Image) {
	if !gameSession.TimerEnabled {
		return
	}

	timerText := gameSession.GetTimerDisplayTime()

	// Create face to measure text width
	face := &text.GoTextFace{
		Source: r.faceSource,
		Size:   8,
	}

	// Measure timer text width to center it horizontally
	timerWidth, _ := text.Measure(timerText, face, 0)

	timerOp := &text.DrawOptions{}
	// Position the timer in horizontal center, vertically centered
	timerOp.GeoM.Translate(float64(config.ScreenWidth)/2-timerWidth/2, float64(config.HudHeight/2-4))

	// Change color to red when timer is running low (less than 10 seconds)
	if gameSession.TimerRemaining <= 10 {
		timerOp.ColorScale.ScaleWithColor(palette.ENDESGA16.Red)
	} else {
		timerOp.ColorScale.ScaleWithColor(theme.DefaultTheme.UITimer)
	}

	text.Draw(screen, timerText, face, timerOp)
}
