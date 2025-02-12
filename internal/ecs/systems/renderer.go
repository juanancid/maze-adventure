package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"maze-adventure/internal/ecs"
	"maze-adventure/internal/ecs/components"
	"reflect"
)

type Renderer struct{}

func (r *Renderer) Draw(w *ecs.World, screen *ebiten.Image) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sprites := w.GetComponents(reflect.TypeOf(&components.Sprite{}))

	for entity, pos := range positions {
		position := pos.(*components.Position)
		sprite := sprites[entity]

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(sprite.(*components.Sprite).Image, options)
		continue
	}
}
