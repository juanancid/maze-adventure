package components

// Patroller represents an NPC that patrols the maze
type Patroller struct {
	ID         int     // Unique identifier for this patroller
	PatrolType int     // Type of patrol pattern (for future use)
	Speed      float64 // Movement speed (for future movement implementation)
	Damage     int     // Damage dealt to player on contact
	IsActive   bool    // Whether this patroller is currently active
}

// NewPatroller creates a new patroller with default values
func NewPatroller(id int) *Patroller {
	return &Patroller{
		ID:         id,
		PatrolType: 0,    // Default patrol type
		Speed:      0.5,  // Default speed (slower than player)
		Damage:     1,    // Default damage amount
		IsActive:   true, // Active by default
	}
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
