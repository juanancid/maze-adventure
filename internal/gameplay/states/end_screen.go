package states

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type EndScreen struct {
	manager *Manager

	blinkTimer int
	blinkOn    bool
}

func NewEndScreen(manager *Manager) *EndScreen {
	return &EndScreen{
		manager: manager,
	}
}

func (s *EndScreen) OnEnter() {
	s.blinkTimer = 0
	s.blinkOn = false
}

func (s *EndScreen) OnExit() {
	// Cleanup explicitly, if needed
}

func (s *EndScreen) Update() error {
	s.blinkTimer++
	if s.blinkTimer >= 60 {
		s.blinkTimer = 0
		s.blinkOn = !s.blinkOn
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	return nil
}

func (s *EndScreen) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)
	drawCenteredText(screen, "Final Protocol Completed", 50, titleFontSize)
	drawCenteredText(screen, "AVA-002: Codename Picatoste", 80, titleFontSize)

	drawCenteredText(screen, "All sectors explored.", 120, regularFontSize)
	drawCenteredText(screen, "Memory integrity stabilized.", 135, regularFontSize)
	drawCenteredText(screen, "No further instructions received.", 150, regularFontSize)

	drawCenteredText(screen, "SYSTEM SHUTDOWN", 200, regularFontSize)
	drawCenteredText(screen, "Thank you for playing.", 215, regularFontSize)

	if s.blinkOn {
		drawCenteredText(screen, "Press ESC to disconnect…", 250, regularFontSize)
	}
}
