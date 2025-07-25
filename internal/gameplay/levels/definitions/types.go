package definitions

import (
	"fmt"
)

var EmptyLevelConfig = LevelConfig{}

// LevelConfig represents the configuration for a game level
type LevelConfig struct {
	Maze         MazeConfig
	Player       PlayerConfig
	Exit         ExitConfig
	Collectibles Collectibles
	Timer        int // Timer in seconds, 0 means no timer for this level
}

// MazeConfig defines the maze dimensions and special cells
type MazeConfig struct {
	Cols                  int     // Number of columns in the maze
	Rows                  int     // Number of rows in the maze
	DeadlyCells           int     // Number of deadly cells to place
	FreezingCells         int     // Number of freezing cells to place
	Patrollers            int     // Number of patroller NPCs to place
	ExtraConnectionChance float64 // Probability (0.0-1.0) of adding extra connections between cells
}

// Validate ensures the maze configuration is valid
func (m MazeConfig) Validate() error {
	if m.Cols <= 0 || m.Rows <= 0 {
		return fmt.Errorf("invalid maze dimensions: cols=%d, rows=%d", m.Cols, m.Rows)
	}

	totalCells := m.Cols * m.Rows
	if m.DeadlyCells < 0 || m.FreezingCells < 0 || m.Patrollers < 0 {
		return fmt.Errorf("special cells/entities count cannot be negative: deadly=%d, freezing=%d, patrollers=%d", m.DeadlyCells, m.FreezingCells, m.Patrollers)
	}

	// Reserve some cells for player, exit, and collectibles (estimate ~3-5 cells)
	reservedCells := 5
	if m.DeadlyCells+m.FreezingCells+m.Patrollers >= totalCells-reservedCells {
		return fmt.Errorf("too many special cells/entities: deadly=%d, freezing=%d, patrollers=%d, available cells=%d", m.DeadlyCells, m.FreezingCells, m.Patrollers, totalCells-reservedCells)
	}

	if m.ExtraConnectionChance < 0.0 || m.ExtraConnectionChance > 1.0 {
		return fmt.Errorf("extra connection chance must be between 0.0 and 1.0, got: %f", m.ExtraConnectionChance)
	}

	return nil
}

// PlayerConfig defines the player properties
type PlayerConfig struct {
	Size int
}

// ExitConfig defines the exit properties
type ExitConfig struct {
	Position Coordinate
	Size     int
}

// Coordinate represents a position in the maze
type Coordinate struct {
	X int
	Y int
}

type Collectibles struct {
	Number int
	Size   int
	Value  int
}
