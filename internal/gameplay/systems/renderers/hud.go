package renderers

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
	"github.com/juanancid/maze-adventure/internal/gameplay/systems/renderers/hud"
)

// HUD is a composite renderer that combines all HUD elements
type HUD struct {
	scoreRenderer  *hud.ScoreRenderer
	levelRenderer  *hud.LevelRenderer
	healthRenderer *hud.HealthRenderer
	timerRenderer  *hud.TimerRenderer
}

func NewHUD() *HUD {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &HUD{
		scoreRenderer:  hud.NewScoreRenderer(faceSource),
		levelRenderer:  hud.NewLevelRenderer(faceSource),
		healthRenderer: hud.NewHealthRenderer(faceSource),
		timerRenderer:  hud.NewTimerRenderer(faceSource),
	}
}

func (r *HUD) Draw(world *entities.World, gameSession *session.GameSession, screen *ebiten.Image) {
	r.scoreRenderer.Draw(gameSession, screen)
	r.levelRenderer.Draw(gameSession, screen)
	r.healthRenderer.Draw(gameSession, screen)
	r.timerRenderer.Draw(gameSession, screen)
}
