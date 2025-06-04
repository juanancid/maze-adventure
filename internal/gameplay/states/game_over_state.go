package states

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverState struct {
	manager *Manager

	blinkTimer int
	blinkOn    bool
}

func NewGameOverState(manager *Manager) *GameOverState {
	return &GameOverState{
		manager: manager,
	}
}

func (s *GameOverState) OnEnter() {
	s.blinkTimer = 0
	s.blinkOn = false
}

func (s *GameOverState) OnExit() {
	// Cleanup explicitly, if needed
}

func (s *GameOverState) Update() error {
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

func (s *GameOverState) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)
	drawCenteredText(screen, "Critical System Failure", 50, titleFontSize)
	drawCenteredText(screen, "AVA-002: Codename Picatoste", 80, titleFontSize)

	drawCenteredText(screen, "Memory core integrity compromised.", 120, regularFontSize)
	drawCenteredText(screen, "Life support systems offline.", 135, regularFontSize)
	drawCenteredText(screen, "Emergency shutdown initiated.", 150, regularFontSize)

	drawCenteredText(screen, "SYSTEM SHUTDOWN", 200, regularFontSize)
	drawCenteredText(screen, "Thank you for playing.", 215, regularFontSize)

	if s.blinkOn {
		drawCenteredText(screen, "Press ESC to disconnect…", 250, regularFontSize)
	}
}
