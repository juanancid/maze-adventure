package renderers

import (
	"image/color"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

// getComponentType is a helper function to get the reflect.Type of a component
func getComponentType(component interface{}) reflect.Type {
	return reflect.TypeOf(component)
}

// PatrollerRenderer renders patroller NPCs in the game
type PatrollerRenderer struct{}

// NewPatrollerRenderer creates a new patroller renderer
func NewPatrollerRenderer() PatrollerRenderer {
	return PatrollerRenderer{}
}

// Draw renders all patroller entities
func (pr PatrollerRenderer) Draw(world *entities.World, gameSession *session.GameSession, screen *ebiten.Image) {
	// Query for entities that have Patroller, Position, and Size components
	patrollerEntities := world.QueryComponents(&components.Patroller{}, &components.Position{}, &components.Size{})

	for _, entity := range patrollerEntities {
		patrollerComp := world.GetComponent(entity, getComponentType(&components.Patroller{}))
		positionComp := world.GetComponent(entity, getComponentType(&components.Position{}))
		sizeComp := world.GetComponent(entity, getComponentType(&components.Size{}))

		if patrollerComp == nil || positionComp == nil || sizeComp == nil {
			continue // Skip if any component is missing
		}

		patroller := patrollerComp.(*components.Patroller)
		position := positionComp.(*components.Position)
		size := sizeComp.(*components.Size)

		// Only render active patrollers
		if !patroller.IsPatrollerActive() {
			continue
		}

		// Render the patroller as a distinct colored circle
		renderPatroller(screen, position, size)
	}
}

// renderPatroller draws a patroller at the specified position
func renderPatroller(screen *ebiten.Image, position *components.Position, size *components.Size) {
	// Calculate screen position (add HUD height offset)
	screenX := float32(position.X)
	screenY := float32(position.Y + float64(config.HudHeight))

	// Patroller color - distinctive orange/red color to differentiate from player
	patrollerColor := color.RGBA{R: 255, G: 100, B: 0, A: 255} // Orange

	// Draw the patroller as a filled circle
	radius := float32(size.Width / 2)
	vector.DrawFilledCircle(screen, screenX+radius, screenY+radius, radius, patrollerColor, false)

	// Add a darker border for better visibility
	borderColor := color.RGBA{R: 200, G: 80, B: 0, A: 255} // Darker orange
	vector.StrokeCircle(screen, screenX+radius, screenY+radius, radius, 2, borderColor, false)
}
