package states

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"

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

	drawCenteredText(screen, "MAZE ADVENTURE", 20, titleFontSize)
	drawCenteredText(screen, "Final Protocol Completed", 50, titleFontSize)
	drawCenteredText(screen, "AVA-002: Codename Picatoste", 80, titleFontSize)

	drawCenteredText(screen, "All sectors explored.", 120, regularFontSize)
	drawCenteredText(screen, "Memory integrity stabilized.", 135, regularFontSize)
	drawCenteredText(screen, "No further instructions received.", 150, regularFontSize)

	drawCenteredText(screen, "SYSTEM SHUTDOWN", 200, regularFontSize)
	drawCenteredText(screen, "Thank you for playing.", 215, regularFontSize)

	if s.blinkOn {
		drawCenteredText(screen, "Press SPACE to restart…", 240, regularFontSize)
		drawCenteredText(screen, "Press ESC to disconnect…", 255, regularFontSize)
	}
}
