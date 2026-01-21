package states

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/juanancid/maze-adventure/internal/gameplay/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"
)

type GameOverState struct {
	manager      *Manager
	levelManager *levels.Manager
	config       config.GameConfig

	blinkTimer int
	blinkOn    bool
}

func NewGameOverState(manager *Manager, levelManager *levels.Manager, config config.GameConfig) *GameOverState {
	return &GameOverState{
		manager:      manager,
		levelManager: levelManager,
		config:       config,
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

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// Restart the game by creating a fresh level manager and transitioning to BootState
		newLevelManager := levels.NewManager()
		bootState := NewBootState(s.manager, newLevelManager, s.config)
		s.manager.ChangeState(bootState)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	return nil
}

func (s *GameOverState) Draw(screen *ebiten.Image) {
	screen.Fill(theme.DefaultTheme.MazeBackground)

	// Header section: Title
	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)

	// Horizontal line below title for visual separation
	s.drawHorizontalSeparator(screen, 40)

	// Main content section: Game over message with improved spacing
	drawCenteredText(screen, "UNIT STATUS: OFFLINE", 58, titleFontSize)
	drawCenteredText(screen, "Traversal interrupted", 88, titleFontSize)

	drawCenteredText(screen, "Sector maintenance incomplete", 128, regularFontSize)

	drawCenteredText(screen, "SYSTEM STANDBY", 208, regularFontSize)

	// Horizontal line above instructions for visual separation
	s.drawHorizontalSeparator(screen, 228)

	// Footer section: Interactive instructions
	if s.blinkOn {
		drawCenteredText(screen, "Press SPACE to retry", 243, regularFontSize)
		drawCenteredText(screen, "Press ESC to disconnect", 258, regularFontSize)
	}
}

// drawHorizontalSeparator draws a subtle horizontal line for visual separation
func (s *GameOverState) drawHorizontalSeparator(screen *ebiten.Image, y float32) {
	// Calculate line width (centered, with margins)
	lineWidth := float32(200) // Moderate width for subtle separation
	centerX := float32(240)   // Screen center (480/2)
	startX := centerX - lineWidth/2
	endX := centerX + lineWidth/2

	// Draw line using theme color with slight transparency for subtlety
	vector.StrokeLine(screen, startX, y, endX, y, 1, theme.DefaultTheme.UIIntroText, false)
}
