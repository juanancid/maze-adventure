package mazebuilder

import (
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
		DeadlyCells:   2,
		FreezingCells: 5,
		Seed:          time.Now().UnixNano(),
	}
}

// Build creates a new maze with the specified configuration
func Build(config *BuilderConfig) components.Layout {
	layout := newMazeLayout(config.Width, config.Height)

	r := rand.New(rand.NewSource(config.Seed))
	placeSpecialCells(layout, config, r)

	return layout
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
		cell.SetType(components.CellTypeDeadly)
		layout.SetCell(pos.x, pos.y, cell)
	}

	// Place freezing cells
	startIdx := config.DeadlyCells
	for i := 0; i < config.FreezingCells && startIdx+i < len(positions); i++ {
		pos := positions[startIdx+i]
		cell := layout.GetCell(pos.x, pos.y)
		cell.SetType(components.CellTypeFreezing)
		layout.SetCell(pos.x, pos.y, cell)
	}
}
