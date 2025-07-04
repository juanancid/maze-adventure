package mazebuilder

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/juanancid/maze-adventure/internal/core/components"
)

// BuilderConfig holds the configuration for maze generation
type BuilderConfig struct {
	Width         int
	Height        int
	DeadlyCells   int   // Number of deadly cells to place
	FreezingCells int   // Number of freezing cells to place
	Seed          int64 // Optional seed for random generation
}

// NewBuilderConfig creates a new builder configuration with default values
func NewBuilderConfig(width, height int) *BuilderConfig {
	return &BuilderConfig{
		Width:         width,
		Height:        height,
		DeadlyCells:   0, // Default to 0, should be set by caller
		FreezingCells: 0, // Default to 0, should be set by caller
		Seed:          time.Now().UnixNano(),
	}
}

// Validate ensures the builder configuration is valid
func (b *BuilderConfig) Validate() error {
	if b.Width <= 0 || b.Height <= 0 {
		return fmt.Errorf("invalid maze dimensions: width=%d, height=%d", b.Width, b.Height)
	}

	totalCells := b.Width * b.Height
	if b.DeadlyCells < 0 || b.FreezingCells < 0 {
		return fmt.Errorf("special cells count cannot be negative: deadly=%d, freezing=%d", b.DeadlyCells, b.FreezingCells)
	}

	if b.DeadlyCells+b.FreezingCells >= totalCells {
		return fmt.Errorf("too many special cells: deadly=%d, freezing=%d, total cells=%d", b.DeadlyCells, b.FreezingCells, totalCells)
	}

	return nil
}

// Build creates a new maze with the specified configuration
func Build(config *BuilderConfig) (components.Layout, error) {
	if err := config.Validate(); err != nil {
		return components.Layout{}, fmt.Errorf("invalid builder config: %w", err)
	}

	layout := newMazeLayout(config.Width, config.Height)
	r := rand.New(rand.NewSource(config.Seed))
	placeSpecialCells(layout, config, r)

	return layout, nil
}

// placeSpecialCells randomly places special cells in the maze
func placeSpecialCells(layout components.Layout, config *BuilderConfig, r *rand.Rand) {
	// Create a list of all possible positions
	positions := make([]struct{ x, y int }, 0, layout.Cols()*layout.Rows())
	for y := 0; y < layout.Rows(); y++ {
		for x := 0; x < layout.Cols(); x++ {
			positions = append(positions, struct{ x, y int }{x, y})
		}
	}

	// Shuffle positions
	r.Shuffle(len(positions), func(i, j int) {
		positions[i], positions[j] = positions[j], positions[i]
	})

	// Place deadly cells
	for i := 0; i < config.DeadlyCells && i < len(positions); i++ {
		pos := positions[i]
		cell := layout.GetCell(pos.x, pos.y)
		deadlyCell := components.NewDeadlyCell(cell.GetWalls())
		layout.SetCell(pos.x, pos.y, deadlyCell)
	}

	// Place freezing cells
	startIdx := config.DeadlyCells
	for i := 0; i < config.FreezingCells && startIdx+i < len(positions); i++ {
		pos := positions[startIdx+i]
		cell := layout.GetCell(pos.x, pos.y)
		freezingCell := components.NewFreezingCell(cell.GetWalls())
		layout.SetCell(pos.x, pos.y, freezingCell)
	}
}
