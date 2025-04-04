package updaters

import "github.com/juanancid/maze-adventure/internal/core/components"

// boundingBox clearly represents a simplified bounding box
type boundingBox struct {
	x, y          float64
	width, height float64
}

func newBoundingBox(pos *components.Position, size *components.Size) boundingBox {
	return boundingBox{
		x:      pos.X,
		y:      pos.Y,
		width:  size.Width,
		height: size.Height,
	}
}

func (bb boundingBox) center() (float64, float64) {
	return bb.x + bb.width/2, bb.y + bb.height/2
}
