package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
)

type CollectiblePickupSystem struct{}

func NewCollectiblePickup() *CollectiblePickupSystem {
	return &CollectiblePickupSystem{}
}

func (s *CollectiblePickupSystem) Update(w *entities.World) {
	players := w.QueryComponents(&components.Position{}, &components.Size{}, &components.InputControlled{})
	if len(players) == 0 {
		return
	}

	var playerEntity entities.Entity
	for _, entity := range players {
		playerEntity = entity
		break
	}

	playerPos := w.GetComponent(playerEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	playerSize := w.GetComponent(playerEntity, reflect.TypeOf(&components.Size{})).(*components.Size)
	playerScore := w.GetComponent(playerEntity, reflect.TypeOf(&components.Score{})).(*components.Score)

	collectibles := w.Query(reflect.TypeOf(&components.Collectible{}), reflect.TypeOf(&components.Position{}), reflect.TypeOf(&components.Size{}))
	for _, collectible := range collectibles {
		cPos := w.GetComponent(collectible, reflect.TypeOf(&components.Position{})).(*components.Position)
		cSize := w.GetComponent(collectible, reflect.TypeOf(&components.Size{})).(*components.Size)
		cData := w.GetComponent(collectible, reflect.TypeOf(&components.Collectible{})).(*components.Collectible)

		if intersects(playerPos, playerSize, cPos, cSize) && cData.Kind == components.CollectibleScore {
			playerScore.Points += cData.Value
			w.RemoveEntity(collectible)
		}
	}
}

func intersects(p1 *components.Position, s1 *components.Size, p2 *components.Position, s2 *components.Size) bool {
	return p1.X < p2.X+s2.Width &&
		p1.X+s1.Width > p2.X &&
		p1.Y < p2.Y+s2.Height &&
		p1.Y+s1.Height > p2.Y
}
