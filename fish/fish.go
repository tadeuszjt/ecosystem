package main

import (
	"github.com/tadeuszjt/data"
	"github.com/tadeuszjt/geom/generic"
)

var (
	fish struct {
		data.Table
		bodyKeys data.RowT[[3]data.Key]
	}
)

func init() {
	fish.Table = data.Table{&fish.bodyKeys}
}

func addFish(pos geom.Vec2[float64]) {
	ori0 := geom.Ori2[float64]{pos.X, pos.Y, 0}
	ori1 := geom.Ori2[float64]{pos.X + 1, pos.Y, 0}
	ori2 := geom.Ori2[float64]{pos.X + 2, pos.Y, 0}

	key0 := addBody(ori0)
	key1 := addBody(ori1)
	key2 := addBody(ori2)

	addJoint(key0, key1, geom.Vec2[float64]{0.5, 0}, geom.Vec2[float64]{-0.5, 0})
	addJoint(key1, key2, geom.Vec2[float64]{0.5, 0}, geom.Vec2[float64]{-0.5, 0})

	fish.Append([3]data.Key{key0, key1, key2})
}
