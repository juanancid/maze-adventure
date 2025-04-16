package states

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
)

const (
	bootScreenFile = "internal/engine/assets/images/boot_screen.png"
)

type BootScreen struct {
	manager      *Manager
	levelManager *levels.Manager
	faceSource   *text.GoTextFaceSource
}

func NewBootScreen(manager *Manager, levelManager *levels.Manager) *BootScreen {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &BootScreen{
		manager:      manager,
		levelManager: levelManager,
		faceSource:   s,
	}
}

func (m *BootScreen) OnEnter() {
	// Initialize resources explicitly if needed
}

func (m *BootScreen) OnExit() {
	// Cleanup explicitly if needed
}

func (m *BootScreen) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		playing := NewPlayingState(m.manager, m.levelManager)
		m.manager.ChangeState(playing)
	}
	return nil
}

func (m *BootScreen) Draw(screen *ebiten.Image) {
	bgColor := color.RGBA{R: 0x00, G: 0x14, B: 0x1D, A: 0xFF}
	screen.Fill(bgColor)

	options := &ebiten.DrawImageOptions{}
	bootScreen := utils.MustLoadSprite(bootScreenFile)
	options.GeoM.Translate(140, 90)
	screen.DrawImage(bootScreen, options)

	textHorizontalCenter := float64(config.ScreenWidth / 2)
	textTitleSize := float64(16)
	textNormalSize := float64(8)

	m.drawText(screen, "MAZE ADVENTURE", textHorizontalCenter, 20, textTitleSize)

	// Protocol and codename section
	m.drawText(screen, "Reactivation Protocol: AVA-002", textHorizontalCenter, 50, textNormalSize)
	m.drawText(screen, "Codename: Picatoste", textHorizontalCenter, 65, textNormalSize)

	// System information section
	m.drawText(screen, "MEMORY CORE INTEGRITY: 12%", textHorizontalCenter, 200, textNormalSize)
	m.drawText(screen, "SECTOR MAP: UNAVAILABLE", textHorizontalCenter, 215, textNormalSize)
	m.drawText(screen, "LAST BOOT: UNKNOWN", textHorizontalCenter, 230, textNormalSize)

	// Call to action
	m.drawText(screen, "Press SPACE to wake up…", textHorizontalCenter, 250, 8)
}

func (m *BootScreen) drawText(screen *ebiten.Image, txt string, x, y float64, size float64) {
	textColor := color.RGBA{R: 0x0D, G: 0x82, B: 0x5D, A: 0xFF}

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(textColor)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter

	face := &text.GoTextFace{
		Source: m.faceSource,
		Size:   size,
	}

	text.Draw(screen, txt, face, op)
}
