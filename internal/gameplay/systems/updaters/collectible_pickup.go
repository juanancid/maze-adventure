package updaters

import (
	"reflect"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/core/queries"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

type CollectiblePickupSystem struct{}

func NewCollectiblePickup() *CollectiblePickupSystem {
	return &CollectiblePickupSystem{}
}

func (s *CollectiblePickupSystem) Update(world *entities.World, gameSession *session.GameSession) {
	playerEntity, found := queries.GetPlayerEntity(world)
	if !found {
		return
	}

	playerPos := world.GetComponent(playerEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	playerSize := world.GetComponent(playerEntity, reflect.TypeOf(&components.Size{})).(*components.Size)
	playerScore := world.GetComponent(playerEntity, reflect.TypeOf(&components.Score{})).(*components.Score)

	collectibles := world.Query(reflect.TypeOf(&components.Collectible{}), reflect.TypeOf(&components.Position{}), reflect.TypeOf(&components.Size{}))
	for _, collectible := range collectibles {
		cPos := world.GetComponent(collectible, reflect.TypeOf(&components.Position{})).(*components.Position)
		cSize := world.GetComponent(collectible, reflect.TypeOf(&components.Size{})).(*components.Size)
		cData := world.GetComponent(collectible, reflect.TypeOf(&components.Collectible{})).(*components.Collectible)

		if intersects(playerPos, playerSize, cPos, cSize) && cData.Kind == components.CollectibleScore {
			playerScore.Points += cData.Value
			world.RemoveEntity(collectible)
		}
	}
}

func intersects(p1 *components.Position, s1 *components.Size, p2 *components.Position, s2 *components.Size) bool {
	return p1.X < p2.X+s2.Width &&
		p1.X+s1.Width > p2.X &&
		p1.Y < p2.Y+s2.Height &&
		p1.Y+s1.Height > p2.Y
}
