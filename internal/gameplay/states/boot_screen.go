package states

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
)

const (
	introIllustrationFile = "internal/engine/assets/images/intro-illustration.png"
)

type BootScreen struct {
	manager      *Manager
	levelManager *levels.Manager

	sprite *ebiten.Image

	blinkTimer int
	blinkOn    bool
}

func NewBootScreen(manager *Manager, levelManager *levels.Manager) *BootScreen {
	sprite := utils.MustLoadSprite(introIllustrationFile)

	return &BootScreen{
		manager:      manager,
		levelManager: levelManager,
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

	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)
	drawCenteredText(screen, "Reactivation Protocol: AVA-002", 50, regularFontSize)
	drawCenteredText(screen, "Codename: Picatoste", 65, regularFontSize)
	drawCenteredText(screen, "MEMORY CORE INTEGRITY: 12%", 200, regularFontSize)
	drawCenteredText(screen, "SECTOR MAP: UNAVAILABLE", 215, regularFontSize)
	drawCenteredText(screen, "LAST BOOT: UNKNOWN", 230, regularFontSize)

	if s.blinkOn {
		drawCenteredText(screen, "Press SPACE to wake up…", 250, regularFontSize)
	}
}

func (s *BootScreen) drawIntroIllustration(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(140, 90)
	screen.DrawImage(s.sprite, options)
}
