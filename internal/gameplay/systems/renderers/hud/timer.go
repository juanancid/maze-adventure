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
	timerOp := &text.DrawOptions{}
	// Position the timer in the center-top area of the HUD
	timerOp.GeoM.Translate(float64(config.ScreenWidth/2+20), float64(config.HudHeight/2-4))

	// Change color to red when timer is running low (less than 10 seconds)
	if gameSession.TimerRemaining <= 10 {
		timerOp.ColorScale.ScaleWithColor(palette.ENDESGA16.Red)
	} else {
		timerOp.ColorScale.ScaleWithColor(theme.DefaultTheme.UIText)
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
