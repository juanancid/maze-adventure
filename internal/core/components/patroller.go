package components

// PatrolPattern defines different movement patterns for patrollers
type PatrolPattern int

const (
	PatrolPatternRandom    PatrolPattern = iota // Random directional movement at intersections
	PatrolPatternLinear                         // Back-and-forth along corridors
	PatrolPatternPerimeter                      // Follow wall boundaries
	PatrolPatternCross                          // Alternate between horizontal and vertical
)

// PatrollerState tracks the current movement state of a patroller
type PatrollerState struct {
	CurrentDirection    int     // Current movement direction (0=up, 1=right, 2=down, 3=left)
	LastDirectionChange float64 // Time of last direction change
	MovementPhase       int     // Current phase for pattern-specific behavior
	SpawnCol            int     // Original spawn column
	SpawnRow            int     // Original spawn row
}

// Patroller represents an NPC that patrols the maze
type Patroller struct {
	ID         int            // Unique identifier for this patroller
	PatrolType PatrolPattern  // Type of patrol pattern
	Speed      float64        // Movement speed
	Damage     int            // Damage dealt to player on contact
	IsActive   bool           // Whether this patroller is currently active
	State      PatrollerState // Current movement state
}

// NewPatroller creates a new patroller with default values
func NewPatroller(id int) *Patroller {
	return &Patroller{
		ID:         id,
		PatrolType: PatrolPatternRandom, // Default to random movement
		Speed:      0.8,                 // Default speed (slower than player)
		Damage:     1,                   // Default damage amount
		IsActive:   true,                // Active by default
		State: PatrollerState{
			CurrentDirection:    0,   // Start moving up
			LastDirectionChange: 0.0, // No previous direction change
			MovementPhase:       0,   // Start at phase 0
			SpawnCol:            0,   // Will be set when placed
			SpawnRow:            0,   // Will be set when placed
		},
	}
}

// NewPatrollerWithPattern creates a patroller with a specific pattern
func NewPatrollerWithPattern(id int, pattern PatrolPattern, spawnCol, spawnRow int) *Patroller {
	patroller := NewPatroller(id)
	patroller.PatrolType = pattern
	patroller.State.SpawnCol = spawnCol
	patroller.State.SpawnRow = spawnRow
	return patroller
}

// GetDamage returns the damage this patroller deals
func (p *Patroller) GetDamage() int {
	return p.Damage
}

// SetActive sets the active state of the patroller
func (p *Patroller) SetActive(active bool) {
	p.IsActive = active
}

// IsPatrollerActive returns whether the patroller is currently active
func (p *Patroller) IsPatrollerActive() bool {
	return p.IsActive
}
