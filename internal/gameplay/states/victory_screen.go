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

type VictoryScreen struct {
	manager    *Manager
	faceSource *text.GoTextFaceSource
}

func NewVictoryScreen(manager *Manager) *VictoryScreen {
	faceSrc, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &VictoryScreen{
		manager:    manager,
		faceSource: faceSrc,
	}
}

func (s *VictoryScreen) OnEnter() {
	// Setup explicitly, e.s., play sound, reset scores, etc.
}

func (s *VictoryScreen) OnExit() {
	// Cleanup explicitly, if needed
}

func (s *VictoryScreen) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		// explicitly restart or go back to menu
		menu := NewBootScreen(s.manager, levels.NewManager())
		s.manager.ChangeState(menu)
	}
	return nil
}

func (s *VictoryScreen) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(160, 100)
	textOp.ColorScale.ScaleWithColor(color.RGBA{R: 255, A: 255})
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	textFace := &text.GoTextFace{
		Source: s.faceSource,
		Size:   8,
	}

	text.Draw(screen, "Game Over - Press ESCAPE", textFace, textOp)
}
