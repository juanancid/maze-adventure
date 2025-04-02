package states

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
)

var (
	faceSource *text.GoTextFaceSource
)

type MenuState struct {
	manager      *Manager
	levelManager *levels.Manager
}

func NewMenuState(manager *Manager, levelManager *levels.Manager) *MenuState {
	return &MenuState{
		manager:      manager,
		levelManager: levelManager,
	}
}

func (m *MenuState) OnEnter() {
	// Initialize resources explicitly if needed
}

func (m *MenuState) OnExit() {
	// Cleanup explicitly if needed
}

func (m *MenuState) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		playing := NewPlayingState(m.manager, m.levelManager)
		m.manager.ChangeState(playing)
	}
	return nil
}

func (m *MenuState) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(160, 100)
	textOp.ColorScale.ScaleWithColor(color.White)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	textFace := &text.GoTextFace{
		Source: faceSource,
		Size:   8,
	}

	text.Draw(screen, "Press SPACE to start", textFace, textOp)
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	faceSource = s
}
