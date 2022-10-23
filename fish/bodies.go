package main

import (
	"github.com/tadeuszjt/data"
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
	"github.com/tadeuszjt/phys2D"
)

var (
	polyHead         = geom.Poly[float32]{{0, -0.2}, {0.6, -0.2}, {0.8, 0}, {0.6, -0.2}, {0, -0.2}}
	polyHeadDrawData = []float32{}

	poly         = geom.Poly[float32]{{0, 0}, {1, -0.2}, {1.2, 0}, {1, 0.2}}
	polyDrawData = []float32{}
	polyColour   = gfx.Green

	world *phys2D.World

	bodies struct {
		data.KeyMap
		physKeys  data.RowT[data.Key]
		colours   data.RowT[gfx.Colour]
		polys     data.RowT[geom.Poly[float32]]
		polyIndex data.RowT[int]
	}
)

func init() {
	{
		centroid := polyHead.Centroid()
		for i := range polyHead {
			polyHead[i] = polyHead[i].Minus(centroid)
		}
		polyHeadDrawData = []float32{
			polyHead[0].X, polyHead[0].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[1].X, polyHead[1].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[3].X, polyHead[3].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[0].X, polyHead[0].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[3].X, polyHead[3].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[4].X, polyHead[4].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[1].X, polyHead[1].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[2].X, polyHead[2].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
			polyHead[3].X, polyHead[3].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
		}
	}

	centroid := poly.Centroid()
	for i := range poly {
		poly[i] = poly[i].Minus(centroid)
	}
	polyDrawData = []float32{
		poly[0].X, poly[0].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
		poly[1].X, poly[1].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
		poly[2].X, poly[2].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
		poly[0].X, poly[0].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
		poly[2].X, poly[2].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
		poly[3].X, poly[3].Y, 0, 0, polyColour.R, polyColour.G, polyColour.B, polyColour.A,
	}

	world = phys2D.NewWorld()
	world.Gravity = geom.Ori2[float64]{}
	world.AirDensity = 0.1

	bodies.KeyMap = data.MakeKeyMap(data.Table{
		&bodies.physKeys,
		&bodies.colours,
		&bodies.polys,
		&bodies.polyIndex,
	})
}

func addJoint(key0, key1 data.Key, offset0, offset1 geom.Vec2[float64]) {
	index0 := bodies.GetIndex(key0)
	index1 := bodies.GetIndex(key1)
	world.AddJoint(bodies.physKeys[index0], bodies.physKeys[index1], offset0, offset1)
}

func addBody(ori geom.Ori2[float64]) data.Key {
	area := poly.Area()
	moi := poly.MomentOfInertia()
	mass := geom.Ori2[float64]{float64(area), float64(area), float64(moi)}

	phyKey := world.AddBody(ori, mass)
	colour := gfx.ColourRand()

	j := len(poly) - 1
	for i := range poly {
		world.AddDragPlate(
			phyKey,
			geom.Vec2Convert[float32, float64](poly[i]),
			geom.Vec2Convert[float32, float64](poly[j]),
		)
		j = i
	}

	return bodies.Append(phyKey, colour, poly, 0)
}

func drawBodies(c gfx.Canvas, mat geom.Mat3[float32]) {
	for i := range bodies.physKeys {
		ori := geom.Ori2Convert[float64, float32](world.GetOrientation(bodies.physKeys[i]))
		modelMat := ori.Mat3Transform()
		modelToFrame := mat.Product(modelMat)
		c.Draw2DVertexData(polyDrawData, nil, &modelToFrame)
	}
}
