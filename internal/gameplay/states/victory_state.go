package states

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/juanancid/maze-adventure/internal/gameplay/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"
)

type VictoryState struct {
	manager      *Manager
	levelManager *levels.Manager
	config       config.GameConfig

	blinkTimer int
	blinkOn    bool
}

func NewVictoryState(manager *Manager, levelManager *levels.Manager, config config.GameConfig) *VictoryState {
	return &VictoryState{
		manager:      manager,
		levelManager: levelManager,
		config:       config,
	}
}

func (s *VictoryState) OnEnter() {
	s.blinkTimer = 0
	s.blinkOn = false
}

func (s *VictoryState) OnExit() {
	// Cleanup explicitly, if needed
}

func (s *VictoryState) Update() error {
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

func (s *VictoryState) Draw(screen *ebiten.Image) {
	screen.Fill(theme.DefaultTheme.MazeBackground)

	// Header section: Title
	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)

	// Horizontal line below title for visual separation
	s.drawHorizontalSeparator(screen, 40)

	// Main content section: Victory message with improved spacing
	drawCenteredText(screen, "TRAVERSAL COMPLETE", 58, titleFontSize)
	drawCenteredText(screen, "All accessible sectors stabilized", 88, regularFontSize)

	drawCenteredText(screen, "Maintenance task completed", 128, regularFontSize)
	drawCenteredText(screen, "NO FURTHER INSTRUCTIONS", 143, regularFontSize)

	drawCenteredText(screen, "SYSTEM STANDBY", 208, regularFontSize)

	// Horizontal line above instructions for visual separation
	s.drawHorizontalSeparator(screen, 228)

	// Footer section: Interactive instructions
	if s.blinkOn {
		drawCenteredText(screen, "Press SPACE to restart", 243, regularFontSize)
		drawCenteredText(screen, "Press ESC to disconnect", 258, regularFontSize)
	}
}

// drawHorizontalSeparator draws a subtle horizontal line for visual separation
func (s *VictoryState) drawHorizontalSeparator(screen *ebiten.Image, y float32) {
	// Calculate line width (centered, with margins)
	lineWidth := float32(200) // Moderate width for subtle separation
	centerX := float32(240)   // Screen center (480/2)
	startX := centerX - lineWidth/2
	endX := centerX + lineWidth/2

	// Draw line using theme color with slight transparency for subtlety
	vector.StrokeLine(screen, startX, y, endX, y, 1, theme.DefaultTheme.UIIntroText, false)
}
