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

type MissionAccomplished struct {
	manager    *Manager
	faceSource *text.GoTextFaceSource
}

func NewMissionAccomplished(manager *Manager) *MissionAccomplished {
	faceSrc, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &MissionAccomplished{
		manager:    manager,
		faceSource: faceSrc,
	}
}

func (a *MissionAccomplished) OnEnter() {
	// Setup explicitly, e.a., play sound, reset scores, etc.
}

func (a *MissionAccomplished) OnExit() {
	// Cleanup explicitly, if needed
}

func (a *MissionAccomplished) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		// explicitly restart or go back to menu
		menu := NewBootScreen(a.manager, levels.NewManager())
		a.manager.ChangeState(menu)
	}
	return nil
}

func (a *MissionAccomplished) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(160, 100)
	textOp.ColorScale.ScaleWithColor(color.RGBA{R: 255, A: 255})
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	textFace := &text.GoTextFace{
		Source: a.faceSource,
		Size:   8,
	}

	text.Draw(screen, "Game Over - Press ESCAPE", textFace, textOp)
}
