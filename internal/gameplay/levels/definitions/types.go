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
}

// MazeConfig defines the maze dimensions and special cells
type MazeConfig struct {
	Cols          int
	Rows          int
	DeadlyCells   int
	FreezingCells int
}

// Validate ensures the maze configuration is valid
func (m MazeConfig) Validate() error {
	if m.Cols <= 0 || m.Rows <= 0 {
		return fmt.Errorf("invalid maze dimensions: cols=%d, rows=%d", m.Cols, m.Rows)
	}

	totalCells := m.Cols * m.Rows
	if m.DeadlyCells < 0 || m.FreezingCells < 0 {
		return fmt.Errorf("special cells count cannot be negative: deadly=%d, freezing=%d", m.DeadlyCells, m.FreezingCells)
	}

	if m.DeadlyCells+m.FreezingCells >= totalCells {
		return fmt.Errorf("too many special cells: deadly=%d, freezing=%d, total cells=%d", m.DeadlyCells, m.FreezingCells, totalCells)
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
