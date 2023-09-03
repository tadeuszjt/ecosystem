package main

import (
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
)

const (
	arrowWidth = 32
)

var (
	arrows = []struct{ start, end geom.Vec2[float32] }{}
)

func drawArrows(c gfx.Canvas, arrowTex gfx.TexID, mat geom.Mat3[float32]) {
	arrowsData := make([]float32, 0, 6*8*len(arrows))
	texCoords := [4]geom.Vec2[float32]{{0, 0}, {1, 0}, {1, 1}, {0, 1}}

	for _, arrow := range arrows {
		dir := arrow.start.Minus(arrow.end).Normal()
		left := geom.Vec2[float32]{dir.Y, -dir.X}

		verts := []geom.Vec2[float32]{
			arrow.start.Plus(left.ScaledBy(arrowWidth)),
			arrow.end.Plus(left.ScaledBy(arrowWidth)),
			arrow.end.Minus(left.ScaledBy(arrowWidth)),
			arrow.start.Minus(left.ScaledBy(arrowWidth)),
		}

		for _, j := range [6]int{0, 1, 2, 0, 2, 3} {
			arrowsData = append(
				arrowsData,
				verts[j].X, verts[j].Y,
				texCoords[j].X, texCoords[j].Y,
				1, 1, 1, 1,
			)
		}
	}

	c.Draw2DVertexData(arrowsData, &arrowTex, &mat)
}
