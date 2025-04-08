package states

import (
	"bytes"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameOverState struct {
	manager    *Manager
	faceSource *text.GoTextFaceSource
}

func NewGameOverState(manager *Manager) *GameOverState {
	faceSrc, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &GameOverState{
		manager:    manager,
		faceSource: faceSrc,
	}
}

func (g *GameOverState) OnEnter() {
	// Setup explicitly, e.g., play sound, reset scores, etc.
}

func (g *GameOverState) OnExit() {
	// Cleanup explicitly, if needed
}

func (g *GameOverState) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		// explicitly restart or go back to menu
		menu := NewTitleScreen(g.manager, levels.NewManager())
		g.manager.ChangeState(menu)
	}
	return nil
}

func (g *GameOverState) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(160, 100)
	textOp.ColorScale.ScaleWithColor(color.RGBA{R: 255, A: 255})
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	textFace := &text.GoTextFace{
		Source: g.faceSource,
		Size:   8,
	}

	text.Draw(screen, "Game Over - Press ESCAPE", textFace, textOp)
}
