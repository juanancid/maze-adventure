package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/juanancid/maze-adventure/internal/engine/utils"
	"github.com/juanancid/maze-adventure/internal/gameplay/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/levels"
	"github.com/juanancid/maze-adventure/internal/gameplay/theme"
)

type BootState struct {
	stateManager *Manager
	levelManager *levels.Manager
	config       config.GameConfig

	sprite *ebiten.Image

	blinkTimer int
	blinkOn    bool
}

func NewBootState(stateManager *Manager, levelManager *levels.Manager, config config.GameConfig) *BootState {
	// Preload all game assets
	utils.PreloadImages()
	utils.PreloadSounds()

	return &BootState{
		stateManager: stateManager,
		levelManager: levelManager,
		config:       config,
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
		playingState := NewPlayingState(s.stateManager, s.levelManager, s.config)
		s.stateManager.ChangeState(playingState)
	}
	return nil
}

func (s *BootState) Draw(screen *ebiten.Image) {
	screen.Fill(theme.DefaultTheme.IntroBackground)

	// Title
	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)

	// Horizontal line below title for visual separation
	s.drawHorizontalSeparator(screen, 40)

	// Body text with improved spacing
	drawCenteredText(screen, "Maintenance unit πk2t active", 108, regularFontSize)
	drawCenteredText(screen, "Manual traversal required", 123, regularFontSize)

	// Blinking prompt
	if s.blinkOn {
		drawCenteredText(screen, "Press SPACE to begin", 158, regularFontSize)
	}

	// Horizontal line above author credit for visual separation
	s.drawHorizontalSeparator(screen, 240)

	// Author credit
	drawCenteredText(screen, "by Juanan Cid", 255, regularFontSize)
}

// drawHorizontalSeparator draws a subtle horizontal line for visual separation
func (s *BootState) drawHorizontalSeparator(screen *ebiten.Image, y float32) {
	// Calculate line width (centered, with margins)
	lineWidth := float32(200) // Moderate width for subtle separation
	centerX := float32(240)   // Screen center (480/2)
	startX := centerX - lineWidth/2
	endX := centerX + lineWidth/2

	// Draw line using theme color with slight transparency for subtlety
	vector.StrokeLine(screen, startX, y, endX, y, 1, theme.DefaultTheme.UIIntroText, false)
}
