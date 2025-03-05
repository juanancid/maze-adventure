package systems

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/ecs"
	"github.com/juanancid/maze-adventure/internal/ecs/components"
)

type SpriteRenderer struct{}

func (r *SpriteRenderer) Draw(w *ecs.World, screen *ebiten.Image) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sprites := w.GetComponents(reflect.TypeOf(&components.Sprite{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))

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
		options.GeoM.Translate(position.X, position.Y)

		screen.DrawImage(spriteComp.Image, options)
	}
}
