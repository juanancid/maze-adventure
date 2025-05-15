package states

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
)

type BootState struct {
	stateManager *Manager
	levelManager *levels.Manager

	sprite *ebiten.Image

	blinkTimer int
	blinkOn    bool
}

func NewBootState(stateManager *Manager, levelManager *levels.Manager) *BootState {
	// Preload all game assets
	utils.PreloadImages()
	utils.PreloadSounds()

	return &BootState{
		stateManager: stateManager,
		levelManager: levelManager,
		sprite:       utils.GetImage(utils.ImageIntroIllustration),
	}
}

func (s *BootState) OnEnter() {
	s.blinkTimer = 0
	s.blinkOn = false
}

func (s *BootState) OnExit() {}

func (s *BootState) Update() error {
	s.blinkTimer++
	if s.blinkTimer >= 60 {
		s.blinkTimer = 0
		s.blinkOn = !s.blinkOn
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		playingState := NewPlayingState(s.stateManager, s.levelManager)
		s.stateManager.ChangeState(playingState)
	}
	return nil
}

func (s *BootState) Draw(screen *ebiten.Image) {
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

func (s *BootState) drawIntroIllustration(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(140, 90)
	screen.DrawImage(s.sprite, options)
}
