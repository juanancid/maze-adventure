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

	for entity, pos := range positions {
		position := pos.(*components.Position)
		spriteComp, ok := sprites[entity].(*components.Sprite)
		if !ok {
			continue
		}

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(spriteComp.Image, options)
	}
}
