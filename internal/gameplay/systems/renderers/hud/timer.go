package hud

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
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
	timerOp := &text.DrawOptions{}
	// Position the timer in the center-top area of the HUD
	timerOp.GeoM.Translate(float64(config.ScreenWidth/2+20), float64(config.HudHeight/2-4))

	// Change color to red when timer is running low (less than 10 seconds)
	if gameSession.TimerRemaining <= 10 {
		timerOp.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 100, B: 100, A: 255})
	} else {
		timerOp.ColorScale.ScaleWithColor(color.White)
	}

	text.Draw(screen,
		timerText,
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		timerOp,
	)
}
