package states

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
)

const (
	introIllustrationFile = "internal/engine/assets/images/intro_illustration.png"
	titleFontSize         = 16
	regularFontSize       = 8
)

var (
	textColor = color.RGBA{R: 0x36, G: 0x9b, B: 0x48, A: 0xFF}
	bgColor   = color.RGBA{R: 0x00, G: 0x13, B: 0x1F, A: 0xFF}
)

type BootScreen struct {
	manager      *Manager
	levelManager *levels.Manager

	font   *text.GoTextFaceSource
	sprite *ebiten.Image

	blinkTimer int
	blinkOn    bool
}

func NewBootScreen(manager *Manager, levelManager *levels.Manager) *BootScreen {
	fontSource := utils.MustLoadGoTextFaceSource(fonts.PressStart2P_ttf)
	sprite := utils.MustLoadSprite(introIllustrationFile)

	return &BootScreen{
		manager:      manager,
		levelManager: levelManager,
		font:         fontSource,
		sprite:       sprite,
	}
}

func (s *BootScreen) OnEnter() {
	s.blinkTimer = 0
	s.blinkOn = false
}

func (s *BootScreen) OnExit() {}

func (s *BootScreen) Update() error {
	s.blinkTimer++
	if s.blinkTimer >= 60 {
		s.blinkTimer = 0
		s.blinkOn = !s.blinkOn
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		playingScreen := NewPlayingScreen(s.manager, s.levelManager)
		s.manager.ChangeState(playingScreen)
	}
	return nil
}

func (s *BootScreen) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	s.drawIntroIllustration(screen)

	centerX := float64(config.ScreenWidth / 2)
	s.drawText(screen, "MAZE ADVENTURE", centerX, 20, titleFontSize)
	s.drawText(screen, "Reactivation Protocol: AVA-002", centerX, 50, regularFontSize)
	s.drawText(screen, "Codename: Picatoste", centerX, 65, regularFontSize)
	s.drawText(screen, "MEMORY CORE INTEGRITY: 12%", centerX, 200, regularFontSize)
	s.drawText(screen, "SECTOR MAP: UNAVAILABLE", centerX, 215, regularFontSize)
	s.drawText(screen, "LAST BOOT: UNKNOWN", centerX, 230, regularFontSize)

	if s.blinkOn {
		s.drawText(screen, "Press SPACE to wake up…", centerX, 250, regularFontSize)
	}
}

func (s *BootScreen) drawIntroIllustration(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(140, 90)
	screen.DrawImage(s.sprite, options)
}

func (s *BootScreen) drawText(screen *ebiten.Image, txt string, x, y float64, size float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(textColor)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter

	face := &text.GoTextFace{
		Source: s.font,
		Size:   size,
	}

	text.Draw(screen, txt, face, op)
}
