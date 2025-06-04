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

// HUD is a composite renderer that combines all HUD elements
type HUD struct {
	scoreRenderer  *ScoreRenderer
	levelRenderer  *LevelRenderer
	healthRenderer *HealthRenderer
}

func NewHUD() *HUD {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &HUD{
		scoreRenderer:  NewScoreRenderer(faceSource),
		levelRenderer:  NewLevelRenderer(faceSource),
		healthRenderer: NewHealthRenderer(faceSource),
	}
}

func (r *HUD) Draw(world *entities.World, gameSession *session.GameSession, screen *ebiten.Image) {
	r.scoreRenderer.Draw(gameSession, screen)
	r.levelRenderer.Draw(gameSession, screen)
	r.healthRenderer.Draw(gameSession, screen)
}

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
