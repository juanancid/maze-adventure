package renderers

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/juanancid/maze-adventure/internal/core/components"
	"github.com/juanancid/maze-adventure/internal/core/entities"
	"github.com/juanancid/maze-adventure/internal/engine/config"
)

type HUD struct {
	faceSource *text.GoTextFaceSource
}

func NewHUD() *HUD {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &HUD{faceSource: s}
}

func (r *HUD) Draw(w *entities.World, screen *ebiten.Image) {
	// Draw game title
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(8, float64(config.HudHeight/2-4)) // Vertically centered-ish
	textOp.ColorScale.ScaleWithColor(color.White)

	scores := w.GetComponents(reflect.TypeOf(&components.Score{}))
	var scoreEntity entities.Entity
	for score := range scores {
		scoreEntity = score
		break
	}

	text.Draw(screen,
		fmt.Sprintf("SCORE: %d", w.GetComponent(scoreEntity, reflect.TypeOf(&components.Score{})).(*components.Score).Points),
		&text.GoTextFace{
			Source: r.faceSource,
			Size:   8,
		},
		textOp,
	)

	// Draw level number
	levels := w.GetComponents(reflect.TypeOf(&components.Level{}))
	if len(levels) > 0 {
		// Get the first level component (there should only be one)
		var levelEntity entities.Entity
		for entity := range levels {
			levelEntity = entity
			break
		}

		level := w.GetComponent(levelEntity, reflect.TypeOf(&components.Level{})).(*components.Level)

		levelText := fmt.Sprintf("SECTOR %d", level.Number)
		levelOp := &text.DrawOptions{}
		levelOp.GeoM.Translate(float64(config.ScreenWidth-100), float64(config.HudHeight/2-4))
		levelOp.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen,
			levelText,
			&text.GoTextFace{
				Source: r.faceSource,
				Size:   8,
			},
			levelOp,
		)
	}
}
