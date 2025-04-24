package renderers

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
	"github.com/juanancid/maze-adventure/internal/gameplay/session"
)

type Sprite struct{}

func NewSprite() Sprite {
	return Sprite{}
}

func (r Sprite) Draw(world *entities.World, gameSession *session.GameSession, screen *ebiten.Image) {
	positions := world.GetComponents(reflect.TypeOf(&components.Position{}))
	sprites := world.GetComponents(reflect.TypeOf(&components.Sprite{}))
	sizes := world.GetComponents(reflect.TypeOf(&components.Size{}))

	for entity, pos := range positions {
		position := pos.(*components.Position)
		spriteComp, ok := sprites[entity].(*components.Sprite)
		if !ok {
			continue
		}

		sizeComp, ok := sizes[entity].(*components.Size)
		if !ok {
			continue
		}

		options := &ebiten.DrawImageOptions{}
		imageX := spriteComp.Image.Bounds().Dx()
		imageY := spriteComp.Image.Bounds().Dy()

		// Calculate the scale factors
		scaleX := sizeComp.Width / float64(imageX)
		scaleY := sizeComp.Height / float64(imageY)

		// Apply scaling
		options.GeoM.Scale(scaleX, scaleY)

		// Translate the position based on the scaled dimensions
		options.GeoM.Translate(position.X, position.Y+float64(config.HudHeight))

		screen.DrawImage(spriteComp.Image, options)
	}
}
