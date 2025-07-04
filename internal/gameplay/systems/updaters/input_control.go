package updaters

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

type InputControl struct{}

func NewInputControl() InputControl {
	return InputControl{}
}

func (is InputControl) Update(world *entities.World, gameSession *session.GameSession) {
	entitiesToControl := world.QueryComponents(&components.InputControlled{}, &components.Velocity{})
	for _, entity := range entitiesToControl {
		handlePlayerInput(world, entity, gameSession)
	}
}

func handlePlayerInput(w *entities.World, entity entities.Entity, gameSession *session.GameSession) {
	controlComp := w.GetComponent(entity, reflect.TypeOf(&components.InputControlled{}))
	velocityComp := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{}))

	if controlComp == nil || velocityComp == nil {
		return // Skip if components are missing
	}

	control := controlComp.(*components.InputControlled)
	velocity := velocityComp.(*components.Velocity)

	updateVelocityFromInputWithGameSession(control, velocity, gameSession)
}

func updateVelocityFromInputWithGameSession(control *components.InputControlled, vel *components.Velocity, gameSession *session.GameSession) {
	// Update freeze state first (check if freeze duration has expired)
	gameSession.UpdateFreezeState()

	// Reset velocity
	vel.DX, vel.DY = 0, 0

	// If player is immobilized (frozen), block all movement input
	if gameSession.IsImmobilized() {
		return // Player cannot move while frozen
	}

	// Normal input processing
	if ebiten.IsKeyPressed(control.MoveLeftKey) {
		vel.DX = -1
	}
	if ebiten.IsKeyPressed(control.MoveRightKey) {
		vel.DX = 1
	}
	if ebiten.IsKeyPressed(control.MoveUpKey) {
		vel.DY = -1
	}
	if ebiten.IsKeyPressed(control.MoveDownKey) {
		vel.DY = 1
	}
}
